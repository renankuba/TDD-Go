package money

import (
	"testing"

	"github.com/renankuba/TDD-Go/money"
	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := new(money.Dollar)
	five.New(5)
	ten := new(money.Dollar)
	ten.New(10)
	assert.Equal(t, five.Times(2), ten)
	fifteen := new(money.Dollar)
	fifteen.New(15)
	assert.Equal(t, five.Times(3), fifteen)
}

func TestFrancMultiplication(t *testing.T) {
	five := new(money.Franc)
	five.New(5)
	ten := new(money.Franc)
	ten.New(10)
	assert.Equal(t, five.Times(2), ten)
	fifteen := new(money.Franc)
	fifteen.New(15)
	assert.Equal(t, five.Times(3), fifteen)
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
