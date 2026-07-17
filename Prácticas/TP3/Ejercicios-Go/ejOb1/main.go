package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"
)

// numero mayor que 1 con 2 dividores, el mismo y 1
func esPrimo(n int) bool {
	if n < 2 { 
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// a) una única goroutine
func primosSecuencial(N int) []int {
	var primos []int
	for i := 2; i <= N; i++ {
		if esPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos
}

// tamaño de cada lote que se reparte por el canal de tareas (prueba)
const tamLote = 1000

type lote struct{ inicio, fin int }

// b) en vez de darle a cada goroutine un rango fijo de antemano, se arma un
// canal con muchos lotes chicos (como una pila de tareas en el medio de la
// mesa) y cada goroutine saca el siguiente lote apenas termina el anterior.

func primosParalelo(N int, numGoroutines int) []int {
	if numGoroutines < 1 {
		numGoroutines = 1
	}

	tareas := make(chan lote, numGoroutines)
	resultados := make(chan []int, numGoroutines)
	var wg sync.WaitGroup

	for w := 0; w < numGoroutines; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for l := range tareas {
				var parciales []int
				for n := l.inicio; n <= l.fin; n++ {
					if esPrimo(n) {
						parciales = append(parciales, n)
					}
				}
				resultados <- parciales
			}
		}()
	}

	// productor de tareas: parte [2, N] en lotes chicos y los va mandando
	go func() {
		for inicio := 2; inicio <= N; inicio += tamLote {
			fin := inicio + tamLote - 1
			if fin > N {
				fin = N
			}
			tareas <- lote{inicio, fin}
		}
		close(tareas)
	}()

	go func() {
		wg.Wait()
		close(resultados)
	}()

	var todos []int
	for parcial := range resultados {
		todos = append(todos, parcial...)
	}
	sort.Ints(todos)
	return todos
}

// c) mide cuánto tarda la versión secuencial vs la paralela para el mismo N
func medirSpeedup(N int) {
	p := runtime.NumCPU()

	inicio1 := time.Now()
	primosSeq := primosSecuencial(N)
	t1 := time.Since(inicio1)

	inicioP := time.Now()
	primosPar := primosParalelo(N, p)
	tp := time.Since(inicioP)

	speedup := float64(t1) / float64(tp)
	fmt.Printf("N=%d secuencial=%v paralelo(p=%d)=%v primos=%d/%d S(p)=%.2f\n",
		N, t1, p, tp, len(primosSeq), len(primosPar), speedup)
}

func main() {
	// Verificar que ingresa argumento
	if len(os.Args) > 1 {
		N, err := strconv.Atoi(os.Args[1])
		if err != nil || N < 0 {
			fmt.Println("Uso: ejOb1 <N entero positivo>")
			os.Exit(1)
		}

		fmt.Println("a) Una sola goroutine:")
		fmt.Println(primosSecuencial(N))

		fmt.Println("\nb) Múltiples goroutines:")
		fmt.Println(primosParalelo(N, runtime.NumCPU()))
		return
	}

	fmt.Println("c) Speed-up para distintos N (sin argumento se corre el benchmark):")
	for _, N := range []int{1_000, 100_000, 1_000_000} {
		medirSpeedup(N)
	}
}
