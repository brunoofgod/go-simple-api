package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Shoes", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, p.Name, "Shoes")
	assert.Equal(t, p.Price, 10.0)
}
func TestNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Shoes", 0)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrPriceIsRequired)
}
