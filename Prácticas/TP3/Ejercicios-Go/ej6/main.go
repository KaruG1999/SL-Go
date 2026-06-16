package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ===== Forma 1: Memoria compartida con Mutex =====

var mu1 sync.Mutex
var buffer1 []int

func productor1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		val := rand.Intn(100)
		mu1.Lock()
		buffer1 = append(buffer1, val)
		fmt.Printf("[Mutex] Productor %d produjo %d\n", id, val)
		mu1.Unlock()
	}
}

func consumidor1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	consumed := 0
	for consumed < 3 {
		mu1.Lock()
		if len(buffer1) > 0 {
			val := buffer1[0]
			buffer1 = buffer1[1:]
			mu1.Unlock()
			fmt.Printf("[Mutex] Consumidor %d procesó %d\n", id, val)
			consumed++
		} else {
			mu1.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func forma1() {
	fmt.Println("=== Forma 1: Mutex ===")
	var wg sync.WaitGroup
	wg.Add(4)
	go productor1(1, &wg)
	go productor1(2, &wg)
	go consumidor1(1, &wg)
	go consumidor1(2, &wg)
	wg.Wait()
}

// ===== Forma 2: Canal unbuffered =====

func productor2(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		val := rand.Intn(100)
		fmt.Printf("[Unbuffered] Productor %d produjo %d\n", id, val)
		ch <- val
	}
}

func consumidor2(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("[Unbuffered] Consumidor %d procesó %d\n", id, val)
	}
}

func forma2() {
	fmt.Println("\n=== Forma 2: Canal unbuffered ===")
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(4)
	go productor2(1, ch, &wg)
	go productor2(2, ch, &wg)
	go consumidor2(1, ch, &wg)
	go consumidor2(2, ch, &wg)
	wg.Wait()
}

// ===== Forma 3: Canal buffered + terminación dinámica con WaitGroup =====

func productor3(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		val := rand.Intn(100)
		fmt.Printf("[Buffered] Productor %d produjo %d\n", id, val)
		ch <- val
	}
}

func forma3() {
	fmt.Println("\n=== Forma 3: Canal buffered + terminación dinámica ===")
	const numProductores = 2
	const numConsumidores = 3 // puede ser distinto al número de productores
	ch := make(chan int, numProductores*3)

	var wgProd sync.WaitGroup
	for i := 1; i <= numProductores; i++ {
		wgProd.Add(1)
		go productor3(i, ch, &wgProd)
	}

	// cierra el canal cuando todos los productores terminan
	go func() {
		wgProd.Wait()
		close(ch)
	}()

	// los consumidores iteran hasta que el canal se cierre y vacíe
	var wgCons sync.WaitGroup
	for i := 1; i <= numConsumidores; i++ {
		wgCons.Add(1)
		go func(id int) {
			defer wgCons.Done()
			for val := range ch {
				fmt.Printf("[Buffered] Consumidor %d procesó %d\n", id, val)
			}
		}(i)
	}
	wgCons.Wait()
}

func main() {
	forma1()
	forma2()
	forma3()
}
