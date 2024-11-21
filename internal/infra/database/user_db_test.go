package database

import (
	"testing"

	"github.com/brunoofgod/go-simple-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("Johrn Cena", "jc@hj.com", "123456")
	if err != nil {
		t.Error(err)
	}
	userDB := NewUserDB(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	db.First(&userFound, "email = ?", user.Email)

	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.Email, user.Email)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.ID.String(), user.ID.String())
	assert.NotEmpty(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("Johrn Cena", "jc@hj.com", "123456")
	if err != nil {
		t.Error(err)
	}
	userDB := NewUserDB(db)

	err = userDB.Create(user)

	if err != nil {
		t.Error(err)
	}

	userFound, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.Email, user.Email)
}
