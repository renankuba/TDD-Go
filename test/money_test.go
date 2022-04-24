package money

import (
	"testing"

	"github.com/renankuba/TDD-Go/money"
	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := money.NewDollar(5)
	assert.Equal(t, money.NewDollar(10), five.Times(2))
	assert.Equal(t, money.NewDollar(15), five.Times(3))
}

func TestFrancMultiplication(t *testing.T) {
	five := money.NewFranc(5)
	assert.Equal(t, money.NewFranc(10), five.Times(2))
	assert.Equal(t, money.NewFranc(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, money.NewDollar(5), money.NewDollar(5))
	assert.NotEqual(t, money.NewDollar(5), money.NewDollar(6))
	assert.Equal(t, money.NewFranc(5), money.NewFranc(5))
	assert.NotEqual(t, money.NewFranc(5), money.NewFranc(6))
	assert.NotEqual(t, money.NewDollar(5), money.NewFranc(5))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", money.NewDollar(1).Currency())
	assert.Equal(t, "CHF", money.NewFranc(1).Currency())
}
