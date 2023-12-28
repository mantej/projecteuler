package euler

func sumOfSquares(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSums(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
