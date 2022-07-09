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
	assert.Equal(t, "USD", money.NewDollar(1).GetCurrency())
	assert.Equal(t, "CHF", money.NewFranc(1).GetCurrency())
}

func TestDifferentTypeEquality(t *testing.T) {
	assert.Equal(t, money.NewMoney(10, "CHF"), money.NewFranc(10))
}

func TestSimpleAddition(t *testing.T) {
	five := money.NewDollar(5)
	sum := five.Plus(five)
	bank := money.NewBank()
	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, money.NewDollar(10), reduced)
}

func TestPlusReturnsSum(t *testing.T) {
	five := money.NewDollar(5)
	result := five.Plus(five)
	sum := result.(*money.Sum)
	assert.Equal(t, five, sum.Augend)
	assert.Equal(t, five, sum.Addend)
}

func TestReduceSum(t *testing.T) {
	var sum money.Expression = money.NewSum(money.NewDollar(3), money.NewDollar(4))
	bank := money.NewBank()
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, money.NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := money.NewBank()
	result := bank.Reduce(money.NewDollar(1), "USD")
	assert.Equal(t, money.NewDollar(1), result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(money.NewFranc(2), "USD")
	assert.Equal(t, money.NewDollar(1), result)
}

func TestIdentiyRate(t *testing.T) {
	assert.Equal(t, float64(1), money.NewBank().Rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := money.NewDollar(5)
	tenFrancs := money.NewFranc(10)
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, money.NewDollar(10), result)
}
