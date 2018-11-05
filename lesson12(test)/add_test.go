package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{0, 2, 2},
		{0, 0, 0},
		{-1, 1, 1},
		{math.MaxInt32, 1, math.MinInt32},
	}

	for _, test := range tests {
		if actual := add(test.a, test.b); actual != test.c {
			t.Errorf("add(%d,%d); got %d;but expected %d", test.a, test.b, test.c, actual)
		}
	}
}

func BenchmarkAdd(bench *testing.B) {
	a := math.MaxInt64
	b := 1
	c := math.MinInt64
	for i := 0; i < bench.N; i++ {
		if actual := add(a, b); actual != c {
			bench.Errorf("add(%d,%d); got %d;but expected %d", a, b, c, actual)
		}
	}
}
