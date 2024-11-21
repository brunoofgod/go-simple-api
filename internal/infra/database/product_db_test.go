package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/brunoofgod/go-simple-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {

	db := createConnection(t)

	product, err := entity.NewProduct("Shoes", 10.0)
	assert.NoError(t, err)

	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestUpdateProduct(t *testing.T) {

	db := createConnection(t)

	product, err := entity.NewProduct("Shoes", 10.0)
	assert.NoError(t, err)

	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

	product.Name = "Shoes 2"
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)
}

func TestFindById(t *testing.T) {

	db := createConnection(t)

	product, err := entity.NewProduct("Shoes", 10.0)
	assert.NoError(t, err)

	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)
}

func TestFindAll(t *testing.T) {
	db := createConnection(t)

	for i := 0; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Shoes %d", i), rand.Float64()*100.0)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProductDB(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Shoes 0", products[0].Name)
	assert.Equal(t, "Shoes 9", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Shoes 10", products[0].Name)
	assert.Equal(t, "Shoes 19", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Shoes 20", products[0].Name)
	assert.Equal(t, "Shoes 22", products[2].Name)
}

func TestDeleteProduct(t *testing.T) {
	db := createConnection(t)

	product, err := entity.NewProduct("Shoes", 10.0)
	assert.NoError(t, err)

	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}

func createConnection(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	return db
}
