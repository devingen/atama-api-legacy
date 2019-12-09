package util

import (
	"fmt"
	"testing"
)

func TestMaxIterationLimit(t *testing.T) {

	n := 0
	for n < 100 {
		fmt.Println(n, MaxIterationLimit(n))
		n++
	}
}
