package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	in1  := make(chan int)
	in2  := make(chan int)
	done := make(chan struct{})

	go GoChannels(in1, in2, done)

	for i := 0; i < 10; i++ {
		in1 <- i
		in2 <- i*2
	}

	<- done
}

func GoChannels(in1 chan int, in2 chan int, done chan struct{}) {
	valuesFrom1 := make([]int, 0, 10)
	valuesFrom2 := make([]int, 0, 10)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func(){
		defer wg.Done()
		
		for i := 0; i < 10; i++ {
			valuesFrom1 = append(valuesFrom1, <-in1)
			valuesFrom2 = append(valuesFrom2, <-in2)
		}
	} ()

	wg.Wait()
	fmt.Println(valuesFrom1)
	fmt.Println(valuesFrom2)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(_i int) {
			defer wg.Done()

			valuesFrom1[_i] = LongProcess(valuesFrom1[_i])
		} (i)

		wg.Add(1)
		go func(_i int) {
			defer wg.Done()

			valuesFrom2[_i] = LongProcess(valuesFrom2[_i])
		} (i)
	}

	wg.Wait()

	fmt.Println(valuesFrom1)
	fmt.Println(valuesFrom2)

	done <- struct{}{}
}

func LongProcess(x int) int {
	time.Sleep(time.Duration(rand.Int31n(10)) * time.Second)
	return x * (-1)
 }