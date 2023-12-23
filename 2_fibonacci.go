package main

import "fmt"

var history map[int]int

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	_, exists := history[n]
	if !exists {
		history[n] = fibonacci(n-1) + fibonacci(n-2)
	}
	return history[n]
}

func main() {
	history = make(map[int]int)
	i := 1
	var fib, sum int
	for {
		fib = fibonacci(i)
		if fib > 4000000 {
			break
		}
		if fib%2 == 0 {
			sum += fib
		}
		i = i + 1
	}
	fmt.Println(sum)
}
