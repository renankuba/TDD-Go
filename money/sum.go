package money

type Sum struct {
	Augend *Money
	Addend *Money
}

func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.Augend.amount + s.Addend.amount
	return newMoney(amount, to)
}

func NewSum(augend *Money, addend *Money) (s *Sum) {
	s = new(Sum)
	s.Augend = augend
	s.Addend = addend
	return s
}
