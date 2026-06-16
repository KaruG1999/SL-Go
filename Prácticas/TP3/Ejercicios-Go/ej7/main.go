package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ===== Solución 1: RWMutex (Compaerten memoria) =====

var (
	rw sync.RWMutex      // Candado lect/escr
	dato int			// dato que comparten
	wg sync.WaitGroup  // contador de gorutinas
)
	
func lector(id int) {
	defer wg.Done() // Resta 1 a waitgroup
	rw.RLock()  	// Lock lectura -> pueden leer pero no escribir
	fmt.Printf("Lector %d lee %d\n", id, dato)
	time.Sleep(time.Second) // Simula tiempo
	rw.RUnlock()  // Unlock lectura -> libera reserva
}

func escritor (id int) {
	defer wg.Done() 	
	rw.Lock() 	// Lock exclusivo
	dato++ 		// modifica dato
	fmt.Printf("Escritor %d lee %d\n", id, dato)
	time.Sleep(time.Second) 
	rw.Unlock() // Unlock exclusivo
}

func main(){
	cantLectores := rand.Intn(5) + 1
	cantEscritores := rand.Intn(5) + 1

	wg.Add(cantLectores + cantEscritores) // Config waitgroup

	for i:=1; i<=cantLectores; i++{
		go lector(i) 	// Lanza cada lector
	}
	for i:=0; i<cantEscritores; i++{
		go escritor(i) // Lanza escritores
	}

	wg.Wait() // Sigue hasta que contador de gorutinas llegue a 0
}