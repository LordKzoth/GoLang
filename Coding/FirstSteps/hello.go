package main

import (
	"fmt"
)

func main() {
	var workArray [10]int
	var changes [6]int

	for index := 0; index < 16; index++ {
		if index < 10 {
			fmt.Scan(&workArray[index])
		} else {
			fmt.Scan(&changes[index-10])
		}
	}

	for index := 0; index < 6; index += 2 {
		workArray[changes[index]], workArray[changes[index+1]] = workArray[changes[index+1]], workArray[changes[index]]
	}

	fmt.Println(workArray)
	fmt.Println(changes)
}
