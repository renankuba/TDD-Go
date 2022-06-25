package money

type Dollar struct {
	Currency
}

func NewDollar(amount float64) *Dollar {
	value := *newCurrency(amount, "USD")
	return &Dollar{value}
}

func (d *Dollar) Times(multiplier float64) Money {
	return NewMoney(d.amount*multiplier, "USD")
}
