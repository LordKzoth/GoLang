package leetcode

import (
	"testing"
)

type addTest struct {
	arg1 []int
	arg2 []int
	expected float64
}

var addTests = []addTest {
	addTest{[]int{1, 3}, []int{2}, 2},
	addTest{[]int{1, 4}, []int{3, 2}, 2.5},
	addTest{[]int{1, 3}, []int{}, 2},
	addTest{[]int{1}, []int{2}, 1.5},
	addTest{[]int{}, []int{2}, 2},
	addTest{[]int{}, []int{}, 0},
}

// Table-Driven Test
func TestFindMedianSortedArrays(t *testing.T) {
	for _, test := range addTests {
		if output := FindMedianSortedArrays(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %.2f not equal to expected %.2f", output, test.expected)
		}
	}
}

// Benchmark Test
func BenchmarkFindMedianSortedArrays(b *testing.B) {
	for _, test := range addTests {
		for i := 0; i < b.N; i++ {
			FindMedianSortedArrays(test.arg1, test.arg2)
		}
	}
}