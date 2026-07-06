package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numCajas = 3
const numClientes = 12

func atender(caja, cliente int) {
	dur := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(dur)
	fmt.Printf("Caja %d atendió cliente %d (%v)\n", caja, cliente, dur)
}

// a) cola global: todas las cajas leen del mismo canal
func colaGlobal() time.Duration {
	inicio := time.Now()
	cola := make(chan int, numClientes)
	var wg sync.WaitGroup

	for id := 0; id < numCajas; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for cliente := range cola {
				atender(id, cliente)
			}
		}(id)
	}

	for i := 0; i < numClientes; i++ {
		cola <- i
	}
	close(cola)
	wg.Wait()
	return time.Since(inicio)
}

// b) colas individuales por caja, asignación round-robin
func roundRobin() time.Duration {
	inicio := time.Now()
	cajas := make([]chan int, numCajas)
	for i := range cajas {
		cajas[i] = make(chan int, numClientes)
	}

	for i := 0; i < numClientes; i++ {
		cajas[i%numCajas] <- i
	}
	for i := range cajas {
		close(cajas[i])
	}

	var wg sync.WaitGroup
	for id, ch := range cajas {
		wg.Add(1)
		go func(id int, ch chan int) {
			defer wg.Done()
			for cliente := range ch {
				atender(id, cliente)
			}
		}(id, ch)
	}
	wg.Wait()
	return time.Since(inicio)
}

// c) colas individuales por caja, asignación a la cola más corta
func colaMasCorta() time.Duration {
	inicio := time.Now()
	cajas := make([]chan int, numCajas)
	for i := range cajas {
		cajas[i] = make(chan int, numClientes)
	}

	var wg sync.WaitGroup
	for id, ch := range cajas {
		wg.Add(1)
		go func(id int, ch chan int) {
			defer wg.Done()
			for cliente := range ch {
				atender(id, cliente)
			}
		}(id, ch)
	}

	for i := 0; i < numClientes; i++ {
		minCola := 0
		for j := 1; j < numCajas; j++ {
			if len(cajas[j]) < len(cajas[minCola]) {
				minCola = j
			}
		}
		cajas[minCola] <- i
		time.Sleep(50 * time.Millisecond) // escalona la llegada
	}
	for i := range cajas {
		close(cajas[i])
	}
	wg.Wait()
	return time.Since(inicio)
}

func main() {
	rand.Seed(1) // misma secuencia de tiempos de atención para comparar en igualdad de condiciones

	fmt.Println("=== a) Cola global ===")
	tGlobal := colaGlobal()

	rand.Seed(1)
	fmt.Println("\n=== b) Round-robin ===")
	tRR := roundRobin()

	rand.Seed(1)
	fmt.Println("\n=== c) Cola más corta ===")
	tCorta := colaMasCorta()

	fmt.Println("\n=== d) Comparación de tiempos ===")
	fmt.Printf("Cola global:   %v\n", tGlobal)
	fmt.Printf("Round-robin:   %v\n", tRR)
	fmt.Printf("Cola más corta: %v\n", tCorta)
}
