package money

type Expression interface {
	Reduce(bank *Bank, to string) *Money
	Plus(addend Expression) Expression
	Times(multiplier float64) Expression
}
