package api

import (
	"fmt"
	"testing"
)

func TestMaxIterationLimit(t *testing.T) {
	m := 2
	for m < 100 {
		fmt.Println(m, MaxIterationLimit(m))
		m++
	}
}

func TestMaxIterationLevel(t *testing.T) {
	m := 2
	for m < 130 {
		fmt.Println(m, MaxIterationLevel(m))
		m++
	}
}
