package main

import (
	"fmt"
	"time"

	channels "github.com/slonegd-otus-go/optional_01_channels"
)

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	out := channels.Merge(ch1, ch2, ch3)

	go func() {
		for value := range out {
			fmt.Println(value)
		}
	}()

	ch1 <- 1
	ch2 <- 2
	ch2 <- 3

	close(ch2)
	time.Sleep(100 * time.Millisecond)

	ch1 <- 10
	ch3 <- 30

	close(ch1)
	time.Sleep(100 * time.Millisecond)

	ch3 <- 300

	close(ch3)
	time.Sleep(100 * time.Millisecond)
}
