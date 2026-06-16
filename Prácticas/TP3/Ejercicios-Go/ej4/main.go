package main

import (
	"fmt"
	"time"
)

// c-i: usando time.Sleep para esperar que hello() termine
func helloSleep() {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
	}
	fmt.Println("Termina Goroutine de hello")
}

// c-ii: usando channel synchronization
func helloChannel(done chan bool) {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
	}
	fmt.Println("Termina Goroutine de hello")
	done <- true
}

func main() {
	// --- c-i: time.Sleep ---
	fmt.Println("=== c-i: time.Sleep ===")
	fmt.Println("Inicia Goroutine del main")
	go helloSleep()
	time.Sleep(time.Second)
	fmt.Println("Termina Goroutine del main")

	// --- c-ii: channel synchronization ---
	fmt.Println("\n=== c-ii: Channel Synchronization ===")
	fmt.Println("Inicia Goroutine del main")
	done := make(chan bool)
	go helloChannel(done)
	<-done
	fmt.Println("Termina Goroutine del main")
}
