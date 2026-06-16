package main

import(
	"fmt"
	"sync"
	"time"
)

func lector(id int, pedirLeer chan int, finLeer chan int, permisoLeer chan bool, wg *sync.WaitGroup){
	defer wg.Done()

	pedirLeer <- id // Encia id al coordinador
	<- permisoLeer // Se bloquea esperando ok del coordinador

	fmt.Printf("Lector %d leyendo...\n", id) //  Lee los datos de forma segura.
	time.Sleep(2 * time.Second)              
	fmt.Printf("Lector %d terminó\n", id)

	finLeer <- id // Envía su ID al coordinador para avisar que ya liberó el recurso.
}

func escritor(id int, pedirEscribir chan int, finEscribir chan bool, wg *sync.WaitGroup){
	defer wg.Done()

	pedirEscribir <- id 
	<- permisoEscribir

	fmt.Printf("Escritor %d escribiendo...\n", id) 
	time.Sleep(2 * time.Second)                    
	fmt.Printf("Escritor %d terminó\n", id)

	finEscribir <- id // Envía su ID al coordinador para avisar que ya terminó su labor.
}

// Servidor centra usa los canales como guardas
func coordinador(pedirLeer chan int, finLeer chan int, pedirEscribir chan int, finEscribir chan int,
	permisoLeer chan bool, permisoEscribir chan bool){

		




	}




func main() {
	var wg sync.WaitGroup

	pedirLeer := make(chan int)
	pedirEscribir := make (chan int)

	finLeer :=make(chan int)
	finEscribir := make(chan int)

	// dan el ok para el arranque de gorutinas
	permisoLeer := make (chan bool)
	permisoEscribir := make (chan bool)

	// Lanzamos coordinador -> sin wg (corre infinitamente)
	go coordinador(pedirLeer, pedirEscribir, finLeer, finEscribir, permisoLeer, permisoEscribir)

	// 
}	go lector(1, pedirLeer, finLeer, permisoLeer, &wg)
	go lector(2, pedirLeer, finLeer, permisoLeer, &wg)
	go lector(3, pedirLeer, finLeer, permisoLeer, &wg)
	go escritor(1, pedirEscribir, finEscribir, permisoEscribir, &wg)
	go escritor(2, pedirEscribir, finEscribir, permisoEscribir, &wg)
	
	wg.Wait()
