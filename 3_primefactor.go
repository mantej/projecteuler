package euler

import (
	"math"
)

func GetPrimeFactors(n int) []int {
	var primes []int
	var i int
	sqrt := int(math.Sqrt(float64(n)))

	for {
		if n%2 == 0 {
			n = n / 2
			primes = append(primes, 2)
		} else {
			break
		}
	}

	for i = 3; i < sqrt; i += 2 {
		for n%i == 0 {
			n = n / i
			primes = append(primes, i)
		}
	}
	return primes
}
