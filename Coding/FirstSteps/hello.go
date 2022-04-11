package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    var text string

    ioScanner := bufio.NewScanner(os.Stdin)
    
    ioScanner.Scan()
    text = ioScanner.Text()

    fmt.Println(text)
}