package money

type Franc struct {
	Currency
}

func NewFranc(amount float64) *Franc {
	value := *newCurrency(amount, "CHF")
	return &Franc{value}
}

func (d *Franc) Times(multiplier float64) Money {
	return NewMoney(d.amount*multiplier, "CHF")
}
