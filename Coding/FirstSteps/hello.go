package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}

func twoSum(nums []int, target int) []int {
    lookup := make(map[int]int)
    for i, v := range nums {
        j, ok := lookup[-v]
        lookup[v - target] = i
        if ok {
            return []int{j, i}
        }
    }
    return []int{}
}