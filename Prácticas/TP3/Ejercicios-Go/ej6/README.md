# Ejercicio 6 — Productores y Consumidores

## Enunciado

Realizar un programa con **2 productores** y **2 consumidores**, cada uno siendo una goroutine. Cada productor produce 3 números aleatorios (entre 0 y 100), esperando un tiempo aleatorio entre 0 y 1 segundo antes de cada producción. Cada consumidor consume 3 números e imprime cuál consumidor lo procesó.

Resolver la comunicación de **tres formas distintas**:
1. Memoria compartida con locks
2. Canal unbuffered
3. Canal buffered

Evaluar ventajas y desventajas. Luego modificar para que los consumidores terminen recién cuando se terminó toda la producción y que la cantidad de productores y consumidores pueda ser distinta.

*Objetivo: Locks, Buffered Channels, WaitGroups*

---

## Lógica de resolución

### Imports necesarios

```go
import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)
```

### Forma 1 — Memoria compartida con Mutex

```go
var mu sync.Mutex
var buffer []int

// productor escribe bajo lock
mu.Lock()
buffer = append(buffer, rand.Intn(100))
mu.Unlock()

// consumidor lee bajo lock
mu.Lock()
if len(buffer) > 0 {
    val := buffer[0]
    buffer = buffer[1:]
    mu.Unlock()
    fmt.Printf("Consumidor %d procesó %d\n", id, val)
}
```

### Forma 2 — Canal unbuffered

```go
ch := make(chan int)
// productor: ch <- valor
// consumidor: val := <-ch
```

Se bloquea hasta que haya un receptor listo. Sincronización punto a punto.

### Forma 3 — Canal buffered

```go
ch := make(chan int, 6) // capacidad = total de items
// productor: ch <- valor (no bloquea si hay espacio)
// consumidor: val := <-ch
```

### Terminación dinámica con WaitGroup + close

```go
var wg sync.WaitGroup
ch := make(chan int, capacidad)

// productores
for i := 0; i < numProductores; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        // producir...
    }(i)
}

// cerrar canal cuando todos los productores terminen
go func() { wg.Wait(); close(ch) }()

// consumidores leen hasta que el canal se cierre
for val := range ch {
    fmt.Printf("Consumidor procesó %d\n", val)
}
```

> Con `range ch`, el consumidor itera hasta que el canal esté cerrado y vacío, sin necesidad de saber de antemano cuántos items habrá.
