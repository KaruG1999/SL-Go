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

// b) múltiples goroutines, cada una procesa un rango
func primosParalelo(N int, numGoroutines int) []int {
	if numGoroutines < 1 {
		numGoroutines = 1
	}
	ch := make(chan []int, numGoroutines)
	var wg sync.WaitGroup

	tamRango := N / numGoroutines
	if tamRango < 1 {
		tamRango = 1
	}

	for i := 0; i < numGoroutines; i++ {
		inicio := i*tamRango + 1
		if inicio > N {
			break
		}
		fin := inicio + tamRango - 1
		if i == numGoroutines-1 || fin > N {
			fin = N
		}
		wg.Add(1)
		go func(inicio, fin int) {
			defer wg.Done()
			var parciales []int
			for n := inicio; n <= fin; n++ {
				if esPrimo(n) {
					parciales = append(parciales, n)
				}
			}
			ch <- parciales
		}(inicio, fin)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var todos []int
	for parcial := range ch {
		todos = append(todos, parcial...)
	}
	sort.Ints(todos)
	return todos
}

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
