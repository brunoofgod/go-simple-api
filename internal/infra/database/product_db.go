package database

import (
	"github.com/brunoofgod/go-simple-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) Update(product *entity.Product) error {
	return p.DB.Save(product).Error
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *Product) FindAll(page, limit int, sortOrder string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	db := p.DB.Order("created_at " + sortOrder)

	if page != 0 && limit != 0 {
		db = db.Limit(limit).Offset((page - 1) * limit)
	}

	err = db.Find(&products).Error

	return products, err
}
