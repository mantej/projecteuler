package main

import "fmt"

func multiples(n int) int {
	var total int

	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}

func main() {
	fmt.Println(multiples(1000))
}
