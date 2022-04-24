package money

import (
	"testing"

	"github.com/renankuba/TDD-Go/money"
	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := money.NewDollar(5)
	assert.Equal(t, five.Times(2), money.NewDollar(10))
	assert.Equal(t, five.Times(3), money.NewDollar(15))
}

func TestFrancMultiplication(t *testing.T) {
	assert.Equal(t, money.NewDollar(5), money.NewDollar(5))
	assert.NotEqual(t, money.NewDollar(5), money.NewDollar(6))
	assert.Equal(t, money.NewFranc(5), money.NewFranc(5))
	assert.NotEqual(t, money.NewFranc(5), money.NewFranc(6))
}
