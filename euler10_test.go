package euler

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	if multiples(1000) != 233168 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 1 Passed")
	}
}

func Test2(t *testing.T) {
	if SumEvenTerms() != 4613732 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 2 Passed")
	}
}

func Test3(t *testing.T) {
	primes := GetPrimeFactors(600851475143)
	if primes[len(primes)-1] != 6857 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 3 Passed")
	}
}

func Test4(t *testing.T) {
	if LargestPalindrome() != 906609 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 4 Passed")
	}
}

func Test5(t *testing.T) {
	if smallest_multiple() != 232792560 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 5 Passed")
	}
}

func Test6(t *testing.T) {
	if squareOfSums(100)-sumOfSquares(100) != 25164150 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 6 Passed")
	}
}

func Test7(t *testing.T) {
	if Prime10001() != 104743 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 7 Passed")
	}
}

func Test8(t *testing.T) {
	if LargestProduct() != 23514624000 {
		t.Fail()
	} else {
		fmt.Println("[*} Problem 8 Passed")
	}
}

func Test9(t *testing.T) {
	if PythagoreanTriplet() != 31875000 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 9 Passed")
	}
}

func Test10(t *testing.T) {
	if PrimeSummation() != 142913828922 {
		t.Fail()
	} else {
		fmt.Println("[*] Problem 10 Passed")
	}
}
