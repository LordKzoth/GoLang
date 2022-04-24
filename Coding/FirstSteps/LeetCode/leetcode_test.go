package leetcode

import (
	"testing"
)

type addTest struct {
	arg1     string
	expected string
}

var addTests = []addTest {
	addTest{"dodosso", "osso"},
	addTest{"ababababa", "ababababa"},
	addTest{"babad", "bab"},
	addTest{"babads", "bab"},
	addTest{"cbbd", "bb"},
	addTest{"cbbds", "bb"},
	addTest{"a", "a"},
	addTest{"aa", "aa"},
	addTest{"aba", "aba"},
	addTest{"dbbd", "dbbd"},
	addTest{"abcd", "a"},
	addTest{"dasa", "asa"},
	addTest{"dassa", "assa"},
	addTest{"dasaa", "asa"},
	addTest{"daasaa", "aasaa"},
	addTest{"dasaaa", "asa"},
	addTest{"ccc", "ccc"},
	addTest{"accca", "accca"},
	addTest{"acccca", "acccca"},
	addTest{"cccc", "cccc"},
	addTest{"ccccc", "ccccc"},
	addTest{"cccccc", "cccccc"},
	addTest{"ccccccc", "ccccccc"},
	addTest{"cccccccc", "cccccccc"},
}

// Table-Driven Test
func TestLongestPalindrome(t *testing.T) {
	for _, test := range addTests {
		if output := longestPalindrome(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

// Benchmark Test
func BenchmarkLongestPalindrome(b *testing.B) {
	for _, test := range addTests {
		for i := 0; i < b.N; i++ {
			longestPalindrome(test.arg1)
		}
	}
}