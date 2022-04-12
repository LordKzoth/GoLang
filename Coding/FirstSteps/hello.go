package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("task.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()


    buf := bytes.NewBuffer(nil)
    io.Copy(buf, file)

    index := 1
    for _, num := range strings.Split(buf.String(), ";") {
        if num == "0" {
            fmt.Println(index)
            break
        }

        index++
    }
}