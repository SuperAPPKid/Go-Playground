package main

import (
	"fmt"
	"time"
)

func newCounter(start, max int) <-chan int {
	out := make(chan int)

	go func() {
		for i := start; i <= max; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func newFilter(in <-chan int, divided int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			if num%divided != 0 {
				out <- num
			}
		}
		close(out)
	}()
	return out
}

func main() {
	start := time.Now()
	max := 1000000
	p := 2
	var ch <-chan int

	ch = newCounter(p, max)
	for {
		fmt.Printf("%d\t", p)
		ch = newFilter(ch, p)

		np, ok := <-ch
		if !ok {
			fmt.Printf("\n%.2f seconds\n", time.Since(start).Seconds())
			break
		}
		p = np
	}
}
