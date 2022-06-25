package money

type Currency struct {
	amount   float64
	currency string
}

func (d *Currency) GetCurrency() string {
	return d.currency
}

func (d *Currency) Times(multiplier float64) Money {
	return nil
}

func (d *Currency) GetAmount() float64 {
	return d.amount
}

func (d *Currency) Equals(other Money) bool {
	if other == nil {
		return false
	}
	return other.GetCurrency() == d.currency &&
		other.GetAmount() == d.amount
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
	GetAmount() float64
	Equals(other Money) bool
}

func NewMoney(amount float64, currency string) Money {
	return newCurrency(amount, currency)
}
