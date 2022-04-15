package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//command := "curl"
	//args 	:= "http://kzoth_domain.com/index.php?filename="

	// Read payload with list of known filenames
	file, err := os.Open("all.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	filenames, _ := ioutil.ReadAll(file)
	
	fmt.Println(filenames)
}