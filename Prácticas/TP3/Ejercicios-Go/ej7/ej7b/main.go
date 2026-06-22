package main

import (
	"fmt"
	"sync"
	"time"
)

import(
	"fmt"
	"sync"
	"time"
)

func lector(int id, pedirLeer chan int, finLeer chan int, permisoLeer chan bool, wg *sunc.WaitGroup){
	defer wg.Done()
	// Pide leer al coord enviandole id y espera respuesta
	pedirLeer <- id
	<- permisoLeer

	fmt.Printf("Lector %d leyendo...\n", id) 
	time.Sleep(2 * time.Second)              
	fmt.Printf("Lector %d terminó\n", id)

	// avisa que termino de leer 
	finLeer <- id
}

func escritor(int id, pedirEscribir chan int, finEscribir chan int, permisoEscribir chan bool, wg *sync.WaitGroup){
	defer wg.Done()

	pedirEscribir <- id
	<- permisoEscribir

	fmt.Printf("Lector %d leyendo...\n",id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Lector %d terminó\n", id)

	finEscribir <- id
}


// Servidor centra usa los canales como guardas
func coordinador(pedirLeer chan int, finLeer chan int, pedirEscribir chan int, finEscribir chan int,
	permisoLeer chan bool, permisoEscribir chan bool){

		lectoresActivos :=0
		escribiendo := false

		// Guardas dinamicas para el select
		var lecturaGuard chan int
		var escrituraGuard chan int

		for {
			// Reset lógico (no se puede leer ni escribir canales con nil -> las apaga)
			lecturaGuard = nil
			escrituraGuard = nil

			// solo lee sino hay nadie escribiendo
			if !escribiendo {
				lecturaGuard = pedirLeer 
			}

			// solo escribe si no hay nadie leyendo ni escribiendo 
			if lectoresActivos == 0 && !escribiendo {
				escrituraGuard = pedirEscribir
			}

			select{
				// Recibe nro Id 
			case id := <-lecturaGuard:
				lectoresActivos++
				fmt.Printf("Coordinador: entra lector %d (lectores = %d)", id, lectoresActivos)
				permisoLeer <- true // Le da el ok para leer
			case id := <- finLeer:
				lectoresActivos--
				fmt.Printf("Coordinador: sale lector %d (lectores = %d)/n",id, lectoresActivos)
			case id := <- escrituraGuard:
				escribiendo = true
				fmt.Printf("Coordinador: entra escritor %d", id)
				permisoEscribir <- true
			case id := <- finEscribir:
				escribiendo=false
				fmt.Printf("Coordinador: sale escritor %d", id)
			}

		}
		
}



func main(){
	var wg sync.WaitGroup

	pedirLeer := make(chan int)
	pedirEscribir := make(chan int)

	finLeer := make(chan int)
	finEscribir := make(chan int)

	permisoLeer := make(chan bool)
	permisoEscribir := make(chan bool)

	go coordinador(pedirLeer, finLeer,pedirEscribir,finEscribir,permisoLeer,permisoEscribir)

	wg.Add(5)
	go lector(1, pedirLeer, finLeer, permisoLeer, &wg)
	go lector(2, pedirLeer, finLeer, permisoLeer, &wg)
	go lector(3, pedirLeer, finLeer, permisoLeer, &wg)
	go escritor(1, pedirEscribir, finEscribir, permisoEscribir, &wg)
	go escritor(2, pedirEscribir, finEscribir, permisoEscribir, &wg)

	wg.Wait()
}