package money

import (
	"testing"

	"github.com/renankuba/TDD-Go/money"
	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := new(money.Dollar)
	five.New(5)
	product := five.Times(2)
	assert.Equal(t, product.Amount, 10.0)
	product = five.Times(3)
	assert.Equal(t, product.Amount, 15.0)
}

func TestEquality(t *testing.T) {
	one := new(money.Dollar)
	one.New(5)
	two := new(money.Dollar)
	two.New(5)
	three := new(money.Dollar)
	three.New(6)
	assert.Equal(t, one, two)
	assert.NotEqual(t, one, three)
}
