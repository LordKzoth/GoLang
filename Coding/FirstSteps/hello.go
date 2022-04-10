package main

import (
	"fmt"
	"math"
)

func main() {
	object := Circle{0, 0, 5}
    fmt.Println(object.GetArea())
}

type Circle struct {
    x, y float64
    r float64
}

func (circle Circle) GetArea() float64 {
    return math.Pi * circle.r * circle.r
}