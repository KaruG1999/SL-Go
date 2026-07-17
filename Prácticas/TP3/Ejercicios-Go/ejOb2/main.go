package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numCajas = 3
const numClientes = 12

// Ahora cada cliente tiene su numero de productos asociados
type cliente struct (
	id: int,
	cantProd: int,
)

// se generan todas las duraciones acá antes de arrancar ninguna goroutine,
// para que el cliente i tarde siempre lo mismo en las tres versiones    // Modificar? 
func generarDuraciones() []time.Duration {
	duraciones := make([]time.Duration, numClientes)
	for i := range duraciones {
		duraciones[i] = time.Duration(rand.Intn(1000)) * time.Millisecond
	}
	return duraciones
}

func atender(caja, cliente int, dur time.Duration) {
	time.Sleep(dur)
	fmt.Printf("Caja %d atendió cliente %d (%v)\n", caja, cliente, dur)
}

// a) cola global: todas las cajas leen del mismo canal
func colaGlobal(duraciones []time.Duration) time.Duration {
	inicio := time.Now()
	cola := make(chan int, numClientes)
	var wg sync.WaitGroup

	for id := 0; id < numCajas; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for cliente := range cola {
				atender(id, cliente, duraciones[cliente])
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
func roundRobin(duraciones []time.Duration) time.Duration {
	inicio := time.Now()
	cajas := make([]chan int, numCajas)
	for i := range cajas {
		cajas[i] = make(chan int, numClientes)
	}

	for i := 0; i < numClientes; i++ {
		cajas[i%numCajas] <- i // reparte en orden fijo: cliente 0->caja0, 1->caja1, 2->caja2, 3->caja0...
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
				atender(id, cliente, duraciones[cliente])
			}
		}(id, ch)
	}
	wg.Wait()
	return time.Since(inicio)
}

// c) colas individuales por caja, asignación a la cola más corta

// Implementar caja rapida  
// Cliente genera num aleatorio (1 al 100) -> Pero ahora cliente tendria un numero y una cant de prod asignadas
// Deberia tener un struct Cliente 
// si es menor igual a 10 -> se envia a la caja rapida

func colaMasCorta(duraciones []time.Duration) time.Duration {
	inicio := time.Now()
	cajas := make([]chan int, numCajas) // uso otra caja aparte 
	cajaRapida := make([]chan int, 1)

	// creo cliente le asigno su id y Genero aleatorios y los asigno a clientes 
	for i:=1; i<numClientes; i++ {

		cantProd := rand.Intn(100) // Chequea el random
	}
    

	for i := range cajas {
		cajas[i] = make(chan int, numClientes)
	}

	var wg sync.WaitGroup
	for id, ch := range cajas {
		wg.Add(1)
		go func(id int, ch chan int) {
			defer wg.Done()
			for cliente := range ch {
				// 
				if cliente.cantProd <= 10 {
					atender(id, cliente, duraciones[cliente])
				}
				atender(id, cliente, duraciones[cliente])
			}
		}(id, ch)
	}

	for i := 0; i < numClientes; i++ {
		minCola := 0
		for j := 1; j < numCajas; j++ {
			// len(canal) -> clientes esperando a ser antendidos en canal
			if len(cajas[j]) < len(cajas[minCola]) {
				minCola = j
			}
		}
		cajas[minCola] <- i
		time.Sleep(50 * time.Millisecond) // escalona la llegada
	}
	for i := range cajas {
		close(cajas[i]) //  avisa a ese range que ya está, no va a venir nada más, sale bucle
	}
	wg.Wait()
	return time.Since(inicio)
}

func main() {
	duraciones := generarDuraciones() // mismos tiempos de atención para las tres versiones

	fmt.Println("=== a) Cola global ===")
	tGlobal := colaGlobal(duraciones)

	fmt.Println("\n=== b) Round-robin ===")
	tRR := roundRobin(duraciones)

 	// creo slice de clientes? aca


	fmt.Println("\n=== c) Cola más corta ===")
	tCorta := colaMasCorta(duraciones)

	fmt.Println("\n=== d) Comparación de tiempos ===")
	fmt.Printf("Cola global:   %v\n", tGlobal)
	fmt.Printf("Round-robin:   %v\n", tRR)
	fmt.Printf("Cola más corta: %v\n", tCorta)
}
