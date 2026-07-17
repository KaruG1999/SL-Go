package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
)

const numDatos = 20
const numTrabajadores = 4

type Dato struct {
	numero    int
	prioridad int
}

var fileMu sync.Mutex

func sumarDigitos(n int) int {
	suma := 0
	for n > 0 {
		suma += n % 10
		n /= 10
	}
	return suma
}

func invertirDigitos(n int) int {
	inv := 0
	for n > 0 {
		inv = inv*10 + n%10
		n /= 10
	}
	return inv
}

// el mutex evita que dos trabajadores escriban al mismo archivo a la vez
func escribirArchivo(f *os.File, texto string) {
	fileMu.Lock()
	defer fileMu.Unlock()
	f.WriteString(texto)
}

// procesa datos según su prioridad hasta que el canal se cierre.
// archP0 y archP1 ya vienen abiertos desde main (una sola vez para todo
// el programa, antes se abrían y cerraban en cada tarea -> corrección profes)
//
// ch es "<-chan Dato": el "<-" antes de chan significa "solo lectura"
func trabajador(id int, ch <-chan Dato, acum *int, mu *sync.Mutex, wg *sync.WaitGroup, archP0, archP1 *os.File) {
	defer wg.Done()
	for d := range ch {
		switch d.prioridad {
		case 0:
			res := sumarDigitos(d.numero)
			escribirArchivo(archP0, fmt.Sprintf("(0, %d)\n", res))
			fmt.Printf("trabajador %d: P0 %d -> suma dígitos %d\n", id, d.numero, res)
		case 1:
			res := invertirDigitos(d.numero)
			escribirArchivo(archP1, fmt.Sprintf("(1, %d)\n", res))
			fmt.Printf("trabajador %d: P1 %d -> invertido %d\n", id, d.numero, res)
		case 2:
			fmt.Printf("trabajador %d: P2 %d x 10 = %d\n", id, d.numero, d.numero*10)
		case 3:
			mu.Lock()
			*acum += d.numero
			fmt.Printf("trabajador %d: P3 %d -> acumulado = %d\n", id, d.numero, *acum)
			mu.Unlock()
		}
	}
}

// revisa siempre primero la cola de mayor prioridad: empieza por la 0,
// y solo si está vacía pasa a la 1, después la 2, la 3.
//
// colas es un array de tamaño fijo [4]chan Dato (no un slice), uno por cada
// prioridad. trabajadores es "chan<- Dato": el "<-" después de chan
// significa "solo escritura" (al revés que en trabajador()).
func planificador(colas [4]chan Dato, trabajadores chan<- Dato, total int) {
	despachados := 0
	for despachados < total {
		enviado := false
		for p := 0; p < 4; p++ {
			select {
			case d := <-colas[p]:
				trabajadores <- d
				despachados++
				enviado = true
			default: // sin esto, el select se quedaría esperando algo en colas[p]
			}
			if enviado {
				break // vuelve a arrancar desde p=0
			}
		}
		if !enviado {
			runtime.Gosched() // nada listo en ninguna cola: cede el turno en vez de reintentar sin parar
		}
	}
	close(trabajadores) // avisa a los trabajadores que no va a llegar nada más
}

func main() {
	// se abren una sola vez para todo el programa, no por cada dato procesado
	archP0, err := os.Create("prioridad0.txt")
	if err != nil {
		fmt.Println("error abriendo prioridad0.txt:", err)
		os.Exit(1)
	}
	defer archP0.Close()

	archP1, err := os.Create("prioridad1.txt")
	if err != nil {
		fmt.Println("error abriendo prioridad1.txt:", err)
		os.Exit(1)
	}
	defer archP1.Close()

	var colas [4]chan Dato
	for i := range colas {
		colas[i] = make(chan Dato, numDatos)
	}

	for i := 0; i < numDatos; i++ {
		d := Dato{numero: rand.Intn(10000), prioridad: rand.Intn(4)}
		colas[d.prioridad] <- d
	}

	trabajadores := make(chan Dato, numTrabajadores)
	var acumulado int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for id := 0; id < numTrabajadores; id++ {
		wg.Add(1)
		go trabajador(id, trabajadores, &acumulado, &mu, &wg, archP0, archP1)
	}

	planificador(colas, trabajadores, numDatos)

	wg.Wait()
	fmt.Println("\nProcesamiento terminado. Acumulado final (prioridad 3):", acumulado)
}
