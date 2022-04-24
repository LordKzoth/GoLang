package leetcode

func longestPalindrome(s string) string {
	sLen := len(s)

	dp := make([][]bool, sLen)
	for i := range dp {
		dp[i] = make([]bool, sLen)
	}

	// Len = 1
	var maxLen, currentLen, start int

	for j := 0; j < sLen; j++ {
		for i := j; i < sLen; i++ {
			if s[i] == s[j] {
				dp[i][j] = true
			}

			if j != i {
				continue
			}

			// Set start Len
			if currentLen == 0 && j >= currentLen+1 && dp[i][j-1] { // EVEN
				currentLen = 2

			} else if currentLen == 0 && j >= currentLen+2 && dp[i][j-2] { // ODD
				currentLen = 3

			} else { // Find next point
				if currentLen != 0 && j >= currentLen+1 && dp[i][j-(currentLen+1)] {
					currentLen += 2
				}
			}

			if currentLen > maxLen {
				start = i - currentLen + 1
				maxLen = currentLen
			} else {
				if j >= 1 && dp[i][j-1] { // EVEN
					currentLen = 2

				} else if j >= 2 && dp[i][j-2] { // ODD
					currentLen = 3

				} else {
					currentLen = 0
				}
			}
		}
	}

	if maxLen == 0 {
		return string(s[0])
	} else {
		return s[start : start+maxLen]
	}
}