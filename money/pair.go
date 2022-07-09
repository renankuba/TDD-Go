package money

type Pair struct {
	From string
	To   string
}

func newPair(from string, to string) (p *Pair) {
	p = new(Pair)
	p.From = from
	p.To = to
	return p
}
