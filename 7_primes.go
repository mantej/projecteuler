package euler

type primes struct {
	value, position int
}

func NewPrimes() *primes {
	return &primes{
		value:    3,
		position: 2,
	}
}

func (p *primes) NextPrime() {
	var prime bool
	for i := p.value + 2; ; i += 2 {
		prime = true
		for j := 3; j < i; j += 2 {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			p.value = i
			p.position = p.position + 1
			return
		}
	}
}

func Prime10001() int {
	p := NewPrimes()

	for p.position < 10001 {
		p.NextPrime()
	}

	return p.value
}
