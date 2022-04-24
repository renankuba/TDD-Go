package money

type Dollar struct {
	MoneyValue
}

func NewDollar(amount float64) *Dollar {
	value := *newMoneyValue(amount, "USD")
	return &Dollar{value}
}

func (d *Dollar) Times(multiplier float64) Money {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) Currency() string {
	return d.currency
}
