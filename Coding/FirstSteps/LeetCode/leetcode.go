package leetcode

import (
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	if length := len(nums1); length % 2 != 0 {
		return float64(nums1[length / 2])
	} else {
		if length == 0{
			return 0
		}
		
		return float64(nums1[length / 2] + nums1[(length / 2)-1]) / 2
	}
}