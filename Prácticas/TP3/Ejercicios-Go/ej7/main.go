package main 

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	rw sync.RWMutex 	
	dato int
	wg sync.WaitGroup 	// Contador de gorutinas 
)

func lector(id int) {
	defer wg.Done()
	rw.RLock()
	fmt.Printf("Lector %d lee %d\n", id, dato)
	time.Sleep(time.Second)
	rw.RUnlock()
}

func escritor(id int){
	defer wg.Done()
	rw.Lock()
	dato++
	fmt.Printf("Escritor %d escribe %d\n", id, dato)
	time.Sleep(time.Second)
	rw.Unlock()
}

func main (){
	cantLectores := rand.Intn(5) + 1
	cantEscritores := rand.Intn(5) + 1

	wg.Add(cantEscritores + cantLectores)

	for i := 1; i <= cantLectores; i++ {
		go lector(i)
	}

	for i := 1; i <= cantEscritores; i++ {
		go escritor(i)
	}
	wg.Wait()
}