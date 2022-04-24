package leetcode

func longestPalindrome(s string) string {
	sLen := len(s)

	dp := make([][]bool, sLen)
	for i := range dp {
		dp[i] = make([]bool, sLen)
	}

	// Len = 1
	maxLen, start := 1, 0

	for i := 0; i < sLen; i++ {
		dp[i][i] = true
	}

	// Len = 2
	for i := 0; i < sLen-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			maxLen = 2
			start = i
		}
	}

	// Len > 2
	for k := 3; k <= sLen; k++ {
		for i := 0; i < sLen-k+1; i++ {
			j := i + k - 1

			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true

				if k > maxLen {
					maxLen = k
					start = i
				}
			}
		}
	}

	return s[start : start+maxLen]
}