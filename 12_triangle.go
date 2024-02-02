package euler

type Triangle struct {
	num, index int
}

func NewTriangle() *Triangle {
	return &Triangle{
		num:   1,
		index: 1,
	}
}

// Next updates t to the next Triangle number
func (t *Triangle) Next() {
	t.index = t.index + 1
	t.num += t.index
}

func numDivisors(n int) int {
	var count int
	count = 2 // 1 and the number itself, both are divisors

	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			count = count + 1
		}
	}
	return count
}
