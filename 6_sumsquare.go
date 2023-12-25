package main

import "fmt"

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

func main() {
	fmt.Println(squareOfSums(100) - sumOfSquares(100))
}
