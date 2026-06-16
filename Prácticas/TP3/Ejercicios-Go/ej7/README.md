# Ejercicio 7 — Lectores y Escritores

## Enunciado

Simular el problema de **Lectores/Escritores** de forma que no exista interferencia según las reglas del problema. Plantear dos soluciones:

1. Con **memoria compartida** (RWMutex)
2. Con **canales**

*Objetivo: Locks, Conditions, Canales con guardas*

---

## Lógica de resolución

### Reglas del problema

- Múltiples lectores pueden leer simultáneamente.
- Solo un escritor puede escribir a la vez.
- Mientras hay un escritor activo, ningún lector puede leer (y viceversa).

### Solución 1 — RWMutex

Go tiene `sync.RWMutex` que modela exactamente estas reglas:

```go
var mu sync.RWMutex
var data int

// lector
func leer(id int) {
    mu.RLock()           // permite concurrencia con otros lectores
    defer mu.RUnlock()
    fmt.Printf("Lector %d leyó: %d\n", id, data)
}

// escritor
func escribir(id, val int) {
    mu.Lock()            // exclusión total
    defer mu.Unlock()
    data = val
    fmt.Printf("Escritor %d escribió: %d\n", id, val)
}
```

### Solución 2 — Con canales

Usar una goroutine "coordinadora" que serializa el acceso mediante canales:

```go
type request struct {
    write bool
    val   int
    resp  chan int
}

func coordinador(req chan request) {
    data := 0
    for r := range req {
        if r.write {
            data = r.val
            r.resp <- 0
        } else {
            r.resp <- data
        }
    }
}
```

Los lectores y escritores envían requests al coordinador y esperan respuesta:

```go
resp := make(chan int)
req <- request{write: false, resp: resp}
val := <-resp
```

> Con canales, la exclusión está implícita: el coordinador procesa un request a la vez. La solución con `RWMutex` es más eficiente porque permite paralelismo real entre lectores.

---

## Conceptos de Teoría

**Problema Lectores/Escritores:** problema clásico de concurrencia. Las invariantes son: (1) múltiples lectores pueden leer simultáneamente sin interferencia entre sí; (2) un escritor necesita acceso exclusivo — ni lectores ni otros escritores pueden acceder mientras escribe.

**`sync.RWMutex`:** variante del mutex que distingue lecturas de escrituras. `RLock()`/`RUnlock()` permiten concurrencia entre lectores; `Lock()`/`Unlock()` dan exclusión total. Es la abstracción directa del problema en Go y la solución más eficiente para cargas con muchas lecturas.

**Modelo CSP (Communicating Sequential Processes):** alternativa al modelo de memoria compartida. En lugar de proteger datos con locks, los goroutines se comunican enviando mensajes por canales. El lema de Go: *"No comuniques compartiendo memoria; compartí memoria comunicando"*.

**Goroutine coordinadora:** patrón CSP donde un único goroutine "posee" el dato y lo expone mediante un canal de requests. La exclusión es implícita porque el coordinador procesa un mensaje a la vez. Desventaja frente a `RWMutex`: serializa también las lecturas, perdiendo el paralelismo entre lectores.

**Canal de respuesta (`resp chan int`):** para modelar llamadas síncronas sobre canales, cada request lleva su propio canal de respuesta. El solicitante se bloquea leyendo de ese canal hasta que el coordinador responde.
