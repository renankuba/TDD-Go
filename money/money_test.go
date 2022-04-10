package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := new(Dollar)
	five.Init(5)
	product := five.times(2)
	assert.Equal(t, product.amount, 10.0)
	product = five.times(3)
	assert.Equal(t, product.amount, 15.0)
}
