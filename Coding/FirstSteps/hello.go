package main

import (
    "fmt"
    "strconv"
)

func main() {
    var userInput string
    fmt.Scan(&userInput)
    fmt.Println(FindRoot(userInput))
}

func FindRoot(number string) string{
    if len(number) > 1 {
        return FindRoot(DigitsSum(number))
    } else {
        return number
    }
}

func DigitsSum(number string) string {
    var result int
    for _, element := range number {
        result += int(element)
    }
    
    result -= len(number) * '0'
    return strconv.Itoa(result)
}