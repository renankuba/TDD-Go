package money

type Dollar struct {
	MoneyValue
}

func NewDollar(amount float64) (d *Dollar) {
	d = new(Dollar)
	d.amount = amount
	return d
}

func (d *Dollar) Times(multiplier float64) (product Money) {
	product = NewDollar(d.amount * multiplier)
	return product
}
