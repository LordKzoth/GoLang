package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    inScanner := bufio.NewScanner(os.Stdin)
    inScanner.Scan()
    
    text := inScanner.Text()
    
    nums := make([]float64, 0)
    for _, value := range strings.Split(text, ";") {
        tempValue, _ := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(value, " ", ""), ",", "."), 64)
        nums = append(nums, tempValue)
    }

    fmt.Printf("%.4f", nums[0] / nums[1])
}