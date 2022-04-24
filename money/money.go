package money

type MoneyValue struct {
	amount float64
}

type Money interface {
	Times(multiplier float64) Money
}
