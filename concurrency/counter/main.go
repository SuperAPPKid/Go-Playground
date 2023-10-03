package main

import (
	"fmt"
	"time"
)

func calculate(target int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 1; i <= target; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func main() {
	start := time.Now()
	fmt.Println("Start")

	c := calculate(100000)

	for n := range c {
		fmt.Println(n)
	}

	fmt.Printf("Done... %.5f seconds\n", time.Since(start).Seconds())
}
