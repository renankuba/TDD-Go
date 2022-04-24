package money

type Franc struct {
	Currency
}

func NewFranc(amount float64) *Franc {
	value := *newMoneyValue(amount, "CHF")
	return &Franc{value}
}

func (d *Franc) Times(multiplier float64) Money {
	return NewFranc(d.amount * multiplier)
}
