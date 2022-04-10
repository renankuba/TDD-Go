package money

type Dollar struct {
	amount float64
}

func (d *Dollar) Init(amount float64) {
	d.amount = amount
}

func (d *Dollar) times(multiplier float64) (product *Dollar) {
	product = new(Dollar)
	product.Init(d.amount * multiplier)
	return product
}
