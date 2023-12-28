package euler

func PrimeSummation() int {
	p := NewPrimes()
	var sum int
	sum = 2

	for p.value < 2000000 {
		sum += p.value
		p.NextPrime()
	}
	return sum
}
