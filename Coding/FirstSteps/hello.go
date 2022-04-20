package main

import (
	"fmt"
)

func main() {
    c   := make(chan int)
    end := make(chan int)

    
    go func(){
        for {
            select {
                case value := <- c: {
                    fmt.Println(value)
                }
                case <- end: {
                    fmt.Println("Process complited!")
                    break
                }
            }
        }
    } ()

    squares(c, end)
}

func squares(c chan int, end chan int) {
    defer close(c)
    defer close(end)

    for i := 1; i < 10; i++ {
        c <- i * i
    }

    end <- 0
}