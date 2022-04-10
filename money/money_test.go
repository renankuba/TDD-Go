package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := new(Dollar)
	five.New(5)
	product := five.times(2)
	assert.Equal(t, product.amount, 10.0)
	product = five.times(3)
	assert.Equal(t, product.amount, 15.0)
}

func TestEquality(t *testing.T) {
	one := new(Dollar)
	one.New(5)
	two := new(Dollar)
	two.New(5)
	three := new(Dollar)
	three.New(6)
	assert.Equal(t, one, two)
	assert.NotEqual(t, one, three)
}
