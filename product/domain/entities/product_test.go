package entities_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"product/domain/entities"
	"testing"
)

func TestProductValid(t *testing.T) {
	product, err := entities.NewProduct(faker.NAME, 1000)
	assert.Nil(t, err)
	assert.Equal(t, product.Price, 1000.0)
}
