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

---

## Conceptos de Teoría

**Goroutine:** hilo de ejecución liviano gestionado por el runtime de Go (no por el SO). Se lanza con la palabra clave `go` y corre concurrentemente al goroutine que la invoca.

**Goroutine del main:** es el goroutine principal. Cuando termina, el programa termina y destruye todos los demás goroutines, sin importar si terminaron o no. Por eso sin sincronización `hello()` nunca llega a imprimirse.

**`time.Sleep` como sincronización:** funciona en la práctica pero no garantiza correctitud — si la máquina está lenta o la goroutine tarda más, el resultado cambia. No escala.

**Channel synchronization:** un canal sin buffer (`make(chan bool)`) usado como señal de finalización. El receptor bloquea hasta que la goroutine envíe, garantizando que el trabajo terminó. Es el mecanismo idiomático en Go para sincronización simple.
