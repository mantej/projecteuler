package euler

import (
	"testing"
)

func Test12(t *testing.T) {
	triangle := NewTriangle()

	for {
		triangle.Next()
		count := numDivisors(triangle.num)

		if count > 500 {
			break
		}
	}

	if triangle.num != 76576500 {
		t.Fail()
	}
}
