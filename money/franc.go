package money

type Franc struct {
	amount float64
}

func (d *Franc) New(amount float64) {
	d.amount = amount
}

func (d *Franc) Times(multiplier float64) (product *Franc) {
	product = new(Franc)
	product.New(d.amount * multiplier)
	return product
}
