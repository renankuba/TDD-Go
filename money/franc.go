package money

type Franc struct {
	MoneyValue
}

func NewFranc(amount float64) (d *Franc) {
	d = new(Franc)
	d.amount = amount
	return d
}

func (d *Franc) Times(multiplier float64) (product Money) {
	product = NewFranc(d.amount * multiplier)
	return product
}
