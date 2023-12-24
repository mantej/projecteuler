package main

import "fmt"

func smallest_multiple() int {
	for i := 20; ; i += 20 {
		for j := 20; j >= 1; j-- {
			if j == 1 {
				return i
			}
			if i%j != 0 {
				break
			}
		}
	}
}

func main() {
	fmt.Println(smallest_multiple())
}
