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

---

## Conceptos de Teoría

**Modelo de memoria compartida:** múltiples goroutines acceden a la misma variable. Requiere sincronización explícita para evitar race conditions (lecturas y escrituras concurrentes sobre el mismo dato).

**`sync.Mutex`:** garantiza exclusión mutua. Solo un goroutine puede estar dentro de `Lock()`/`Unlock()` a la vez. El consumidor con mutex necesita hacer *polling* (revisar si hay datos en un loop), lo que consume CPU innecesariamente.

**Canal sin buffer como rendezvous:** productor y consumidor se sincronizan en cada envío/recepción. Garantiza que el valor es entregado antes de continuar. Desventaja: si no hay consumidor listo, el productor se bloquea.

**Canal con buffer:** desacopla productor y consumidor — el productor puede avanzar hasta llenar la capacidad sin esperar. Más eficiente pero requiere dimensionar bien el buffer.

**`sync.WaitGroup`:** contador de goroutines activos. `Add(1)` antes de lanzar, `Done()` al terminar (con `defer`), `Wait()` para bloquear hasta que todos lleguen a cero.

**`close(ch)` + `range`:** patrón idiomático para terminación dinámica. Cerrar el canal señaliza a los consumidores que no habrá más datos; `range ch` itera hasta que el canal esté cerrado y vacío. Permite que productores y consumidores tengan cardinalidades distintas.
