# Ejercicio Obligatorio 2 — Cajas de supermercado

## Enunciado

Simular la atención de clientes en las cajas de un supermercado. La atención de cada cliente se simula con un timer entre 0 y 1 segundo.

**a)** Una **cola global**: los clientes esperan en una única cola y van a la caja disponible.

**b)** **Colas individuales por caja** con asignación **round-robin**: cada cliente nuevo va a la siguiente caja en orden circular.

**c)** **Colas individuales por caja** con asignación a la **caja con cola más corta**.

**d)** Imprimir los tiempos de ejecución de cada implementación y compararlos.

*Objetivo: channels, goroutines, patrones de distribución de trabajo*

---

## Lógica de resolución

### Parámetros comunes

```go
const numCajas    = 3
const numClientes = 12
```

### a) Cola global (channel único)

```go
cola := make(chan int, numClientes)  // canal = cola global

// cajeros: cada caja lee del mismo canal
for id := 0; id < numCajas; id++ {
    go func(id int) {
        for cliente := range cola {
            dur := time.Duration(rand.Intn(1000)) * time.Millisecond
            time.Sleep(dur)
            fmt.Printf("Caja %d atendió cliente %d\n", id, cliente)
            wg.Done()
        }
    }(id)
}

// clientes
for i := 0; i < numClientes; i++ { cola <- i }
close(cola)
wg.Wait()
```

### b) Round-robin

```go
cajas := make([]chan int, numCajas)
for i := range cajas { cajas[i] = make(chan int, numClientes) }

for i, cliente := range clientes {
    cajas[i % numCajas] <- cliente
}
```

### c) Cola más corta

```go
// usar len(cajas[i]) para saber cuántos clientes están esperando
minCola := 0
for i := 1; i < numCajas; i++ {
    if len(cajas[i]) < len(cajas[minCola]) { minCola = i }
}
cajas[minCola] <- cliente
```

### d) Medir tiempos

```go
start := time.Now()
// ... ejecutar variante ...
fmt.Printf("Tiempo: %v\n", time.Since(start))
```

> Con cola global los cajeros nunca están ociosos si hay clientes esperando, por eso suele ser la más eficiente. Round-robin puede desequilibrarse si los tiempos de atención varían.
