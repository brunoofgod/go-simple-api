package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Johrn Cena", "jc@hj.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Email)

	assert.Equal(t, "Johrn Cena", user.Name)
	assert.Equal(t, "jc@hj.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Johrn Cena", "jc@hj.com", "123456")
	assert.Nil(t, err)

	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))

	assert.NotEqual(t, "123456", user.Password)
}
