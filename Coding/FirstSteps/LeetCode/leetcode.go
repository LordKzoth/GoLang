package leetcode

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	minimum := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	maximum := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	l1, l2 := len(nums1), len(nums2)

	if l1 == 0 && l2 == 0 {
		return 0
	} else if l1 == 0 {
		if l2%2 == 0 {
			return float64(nums2[l2/2]+nums2[(l2/2)-1]) / 2
		} else {
			return float64(nums2[l2/2])
		}
	} else if l2 == 0 {
		if l1%2 == 0 {
			return float64(nums1[l1/2]+nums1[(l1/2)-1]) / 2
		} else {
			return float64(nums1[l1/2])
		}
	}

	if l1 > l2 {
		return FindMedianSortedArrays(nums2, nums1)
	}

	max, min := l1, 0
	var median, i, j int

	for min <= max {
		i = (min + max) / 2
		j = (l1+l2+1)/2 - i

		if j < 0 {
			max = i - 1
			continue
		}

		if i < l1 && j > 0 && nums2[j-1] > nums1[i] {
			min = i + 1
		} else if i > 0 && j < l2 && nums2[j] < nums1[i-1] {
			max = i - 1
		} else {
			if i == 0 {
				median = nums2[j-1]
			} else if j == 0 {
				median = nums1[i-1]
			} else {
				median = maximum(nums1[i-1], nums2[j-1])
			}

			break
		}
	}

	if (l1+l2)%2 == 1 {
		return float64(median)
	}
	if i == l1 {
		return float64(median+nums2[j]) / 2
	}
	if j == l2 {
		return float64(median+nums1[i]) / 2
	}

	return float64(median+minimum(nums1[i], nums2[j])) / 2
}