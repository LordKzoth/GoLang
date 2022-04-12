package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
    rd := bufio.NewReader(os.Stdin)
    wr := bufio.NewWriter(os.Stdout)
    
    var sum int
    for {
        num, err := rd.ReadString('\n')

        if err != io.EOF {
            conv, _ := strconv.Atoi(string(num[:len(num) - 2]))
            sum += conv
        } else {
            break
        }
    }

    wr.WriteString(strconv.Itoa(sum))
    wr.Flush()
}