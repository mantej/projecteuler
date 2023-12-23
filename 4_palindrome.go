package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(s string) bool {
	length := len(s)

	if length == 1 {
		return true
	}

	i, j := 0, length-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i = i + 1
		j = j - 1
	}
	return true
}

func main() {
	var isPal bool
	largestPalindrome := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			num := i * j
			strnum := strconv.Itoa(num)
			isPal = isPalindrome(strnum)
			if isPal {
				if num > largestPalindrome {
					largestPalindrome = num
				}
			}
		}
	}

	fmt.Println(largestPalindrome)
}
