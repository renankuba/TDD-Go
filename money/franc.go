package money

type Franc struct {
	MoneyValue
}

func NewFranc(amount float64) *Franc {
	value := *newMoneyValue(amount, "CHF")
	return &Franc{value}
}

func (d *Franc) Times(multiplier float64) Money {
	return NewFranc(d.amount * multiplier)
}

func (d *Franc) Currency() string {
	return d.currency
}
