package money

type Franc struct {
	Money
}

func NewFranc(amount float64) (d *Franc) {
	d = new(Franc)
	d.amount = amount
	return d
}

func (d *Franc) Times(multiplier float64) (product *Franc) {
	product = NewFranc(d.amount * multiplier)
	return product
}
