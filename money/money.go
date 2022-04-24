package money

type MoneyValue struct {
	amount   float64
	currency string
}

func newMoneyValue(amount float64, currency string) (m *MoneyValue) {
	m = new(MoneyValue)
	m.amount = amount
	m.currency = currency
	return m
}

type Money interface {
	Times(multiplier float64) Money
	Currency() string
}
