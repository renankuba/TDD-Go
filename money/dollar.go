package money

type Dollar struct {
	amount float64
}

func (d *Dollar) New(amount float64) {
	d.amount = amount
}

func (d *Dollar) times(multiplier float64) (product *Dollar) {
	product = new(Dollar)
	product.New(d.amount * multiplier)
	return product
}
