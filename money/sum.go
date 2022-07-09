package money

type Sum struct {
	Augend Expression
	Addend Expression
}

func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.Augend.Reduce(bank, to).amount + s.Addend.Reduce(bank, to).amount
	return newMoney(amount, to)
}

func (s *Sum) Plus(addend Expression) Expression {
	return nil
}

func NewSum(augend Expression, addend Expression) (s *Sum) {
	s = new(Sum)
	s.Augend = augend
	s.Addend = addend
	return s
}
