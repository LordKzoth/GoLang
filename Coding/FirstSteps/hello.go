package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    ioScanner := bufio.NewScanner(os.Stdin)
    ioScanner.Scan()

    text := []rune(ioScanner.Text())

    for i := 0; i < len(text) / 2; i++ {
        if !(text[i] == text[len(text) - 1 - i]) {
            fmt.Println("Нет")
            return
        }
    }

    fmt.Println("Палиндром")
}