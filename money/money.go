package money

type Money struct {
	amount   float64
	currency string
}

func (d *Money) Times(multiplier float64) *Money {
	return newMoney(d.amount*multiplier, d.currency)
}

func (d *Money) GetCurrency() string {
	return d.currency
}

func (d *Money) Plus(addend *Money) Expression {
	return newMoney(d.amount+addend.amount, d.currency)
}

func newMoney(amount float64, currency string) (m *Money) {
	m = new(Money)
	m.amount = amount
	m.currency = currency
	return m
}

func NewMoney(amount float64, currency string) *Money {
	return newMoney(amount, currency)
}

func NewFranc(amount float64) *Money {
	return newMoney(amount, "CHF")
}

func NewDollar(amount float64) *Money {
	return newMoney(amount, "USD")
}
