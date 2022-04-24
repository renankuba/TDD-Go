package money

type Currency struct {
	amount   float64
	currency string
}

func (d *Currency) GetCurrency() string {
	return d.currency
}

func newCurrency(amount float64, currency string) (m *Currency) {
	m = new(Currency)
	m.amount = amount
	m.currency = currency
	return m
}

type Money interface {
	Times(multiplier float64) Money
	GetCurrency() string
}
