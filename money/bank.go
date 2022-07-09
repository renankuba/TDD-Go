package money

type Bank struct {
	rates map[Pair]float64
}

func NewBank() (b *Bank) {
	b = new(Bank)
	b.rates = make(map[Pair]float64)
	return b
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(b, to)
}

func (b *Bank) Rate(from string, to string) float64 {
	if from == to {
		return 1
	}
	return b.rates[*newPair(from, to)]
}

func (b *Bank) AddRate(from string, to string, rate float64) {
	b.rates[*newPair(from, to)] = rate
}
