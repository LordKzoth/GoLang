package main

import (
	"fmt"
)

func main() {
	var start, percent, goal int
	fmt.Scan(&start, &percent, &goal)

	var years int = 0
	for currentSum := start; currentSum < goal; currentSum += int(float64(currentSum) * (float64(percent) / 100)) {
		years++
	}

	fmt.Println(years)
}
