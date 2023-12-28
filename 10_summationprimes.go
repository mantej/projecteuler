package main

import "fmt"

type primes2 struct {
	value int
}

func NewPrimes2() *primes2 {
	return &primes2{
		value: 3,
	}
}

func (p *primes2) Next() {
	var isPrime bool
	for i := p.value + 2; ; i += 2 {
		isPrime = true
		for j := 3; j < i; j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			p.value = i
			return
		}
	}
}

func main() {
	p := NewPrimes2()
	var sum int
	sum = 2

	for p.value < 2000000 {
		sum += p.value
		p.Next()
	}

	fmt.Println(sum)
}
