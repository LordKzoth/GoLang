package main

import (
	"fmt"
)

func main() {
	var a int = 1
    var b *int = &a

    *b++

    fmt.Println(*b)
}