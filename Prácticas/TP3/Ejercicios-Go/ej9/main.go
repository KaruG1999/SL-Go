package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Second)
			ch1 <- i
		}
	}()

	go func() {
		for i := 0; ; i++ {
			time.Sleep(2 * time.Second)
			ch2 <- i
		}
	}()

	timeout1 := time.After(5 * time.Second)
	timeout2 := time.After(10 * time.Second)

	for {
		select {
		case v := <-ch1:
			fmt.Println("ch1:", v)
		case v := <-ch2:
			fmt.Println("ch2:", v)
		case <-timeout1:
			fmt.Println("timeout ch1 (5s)")
			ch1 = nil
		case <-timeout2:
			fmt.Println("timeout ch2 (10s)")
			return
		}
	}
}
