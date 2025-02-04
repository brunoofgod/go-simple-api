package database

import (
	"github.com/brunoofgod/go-simple-api/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}
