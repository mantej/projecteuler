package euler

func multiples(n int) int {
	var total int

	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}
