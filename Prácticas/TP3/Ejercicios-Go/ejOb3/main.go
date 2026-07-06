package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

const numDatos = 20
const numWorkers = 4

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

func escribirArchivo(nombre, texto string) {
	fileMu.Lock()
	defer fileMu.Unlock()
	f, err := os.OpenFile(nombre, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error escribiendo archivo:", err)
		return
	}
	defer f.Close()
	f.WriteString(texto)
}

// worker procesa datos según su prioridad hasta que el canal se cierre
func worker(id int, ch <-chan Dato, acum *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range ch {
		switch d.prioridad {
		case 0:
			res := sumarDigitos(d.numero)
			escribirArchivo("prioridad0.txt", fmt.Sprintf("(0, %d)\n", res))
			fmt.Printf("worker %d: P0 %d -> suma dígitos %d\n", id, d.numero, res)
		case 1:
			res := invertirDigitos(d.numero)
			escribirArchivo("prioridad1.txt", fmt.Sprintf("(1, %d)\n", res))
			fmt.Printf("worker %d: P1 %d -> invertido %d\n", id, d.numero, res)
		case 2:
			fmt.Printf("worker %d: P2 %d x 10 = %d\n", id, d.numero, d.numero*10)
		case 3:
			mu.Lock()
			*acum += d.numero
			fmt.Printf("worker %d: P3 %d -> acumulado = %d\n", id, d.numero, *acum)
			mu.Unlock()
		}
	}
}

// revisa siempre primero la cola de mayor prioridad disponible
func scheduler(colas [4]chan Dato, workers chan<- Dato, total int) {
	despachados := 0
	for despachados < total {
		enviado := false
		for p := 0; p < 4; p++ {
			select {
			case d := <-colas[p]:
				workers <- d
				despachados++
				enviado = true
			default:
			}
			if enviado {
				break
			}
		}
		if !enviado {
			runtime.Gosched()
		}
	}
	close(workers)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	os.Remove("prioridad0.txt")
	os.Remove("prioridad1.txt")

	var colas [4]chan Dato
	for i := range colas {
		colas[i] = make(chan Dato, numDatos)
	}

	for i := 0; i < numDatos; i++ {
		d := Dato{numero: rand.Intn(10000), prioridad: rand.Intn(4)}
		colas[d.prioridad] <- d
	}

	workers := make(chan Dato, numWorkers)
	var acumulado int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for id := 0; id < numWorkers; id++ {
		wg.Add(1)
		go worker(id, workers, &acumulado, &mu, &wg)
	}

	scheduler(colas, workers, numDatos)

	wg.Wait()
	fmt.Println("\nProcesamiento terminado. Acumulado final (prioridad 3):", acumulado)
}
