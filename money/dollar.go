package money

type Dollar struct {
	Amount float64
}

func (d *Dollar) New(amount float64) {
	d.Amount = amount
}

func (d *Dollar) Times(multiplier float64) (product *Dollar) {
	product = new(Dollar)
	product.New(d.Amount * multiplier)
	return product
}
