package main

import (
	"fmt"
)

func main() {
	var slice []int = []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4, 5)
	fmt.Println(slice)

	fmt.Printf("%p", slice)
}
