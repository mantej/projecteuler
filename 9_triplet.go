package main

import (
	"fmt"
	"math"
)

func IsPerfectSquare(n int) bool {
	var int_root int
	int_root = int(math.Sqrt(float64(n)))
	return int_root*int_root == n
}

func main() {
done:
	for a := 1; a < 500; a++ {
		for b := 1; b < 500; b++ {
			c2 := a*a + b*b
			if IsPerfectSquare(c2) {
				c := int(math.Sqrt(float64(c2)))
				if a+b+c == 1000 {
					fmt.Println(a * b * c)
					break done
				}
			}
		}
	}
}
