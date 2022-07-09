package money

type Money struct {
	amount   float64
	currency string
}

func (m *Money) Times(multiplier float64) Expression {
	return newMoney(m.amount*multiplier, m.currency)
}

func (m *Money) GetCurrency() string {
	return m.currency
}

func (m *Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

func (m *Money) Reduce(bank *Bank, to string) *Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
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
