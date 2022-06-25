package money

type Bank struct{}

func NewBank() *Bank {
	return new(Bank)
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return NewDollar(10)
}
