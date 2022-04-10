package main

import (
	"fmt"
)

func main() {
	var a int = 1
    var b *int = &a

    *b++

    fmt.Println(a, *b)

    var c **int = &b;

    fmt.Println(a, *b, **c)

    **c++

    fmt.Println(a, *b, **c)
    fmt.Println(b, *c, c)

}