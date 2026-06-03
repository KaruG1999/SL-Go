# Ejercicio 4 — Goroutines básicas

## Enunciado

Ejecutar el siguiente programa:

```go
package main

import "fmt"

func main() {
    fmt.Println("Inicia Goroutine del main")
    go hello()
    fmt.Println("Termina Goroutine del main")
}

func hello() {
    fmt.Println("Inicia Goroutine de hello")
    for i := 0; i < 3; i++ {
        fmt.Println(i, " Hello world")
    }
    fmt.Println("Termina Goroutine de hello")
}
```

**a)** ¿Cuántas veces se imprime Hello world?  
**b)** ¿Cuántas Goroutines tiene el programa?  
**c)** ¿Cómo cambiaría el programa (con la misma cantidad de Goroutines) para que imprima 3 veces Hello world?
- **i)** Usando `time.Sleep`
- **ii)** Usando Channel Synchronization

*Objetivo: goroutines*

---

## Lógica de resolución

### Respuestas a y b

- **a)** Probablemente **0 veces**: el `main` termina antes de que `hello()` pueda ejecutarse, porque `go hello()` lanza la goroutine pero no espera a que termine.
- **b)** **2 goroutines**: la del `main` y la de `hello`.

### c-i) Con `time.Sleep`

```go
import "time"

func main() {
    fmt.Println("Inicia Goroutine del main")
    go hello()
    time.Sleep(time.Second) // dar tiempo a que hello() termine
    fmt.Println("Termina Goroutine del main")
}
```

Desventaja: el tiempo de sleep es arbitrario, no garantiza correctitud.

### c-ii) Con Channel Synchronization

```go
func hello(done chan bool) {
    fmt.Println("Inicia Goroutine de hello")
    for i := 0; i < 3; i++ {
        fmt.Println(i, " Hello world")
    }
    fmt.Println("Termina Goroutine de hello")
    done <- true // señalizar que terminó
}

func main() {
    fmt.Println("Inicia Goroutine del main")
    done := make(chan bool)
    go hello(done)
    <-done // bloquear hasta recibir la señal
    fmt.Println("Termina Goroutine del main")
}
```

> El canal actúa como barrera de sincronización. Es la forma idiomática en Go frente a `time.Sleep`.
