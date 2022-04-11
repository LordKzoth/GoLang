package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
    var text []rune

    ioScanner := bufio.NewScanner(os.Stdin)
    
    ioScanner.Scan()
    text = []rune(ioScanner.Text())

    if (unicode.IsUpper(text[0]) && text[len(text) - 1] == '.') {
        fmt.Println("Right")
    } else {
        fmt.Println("Wrong")
    }
}