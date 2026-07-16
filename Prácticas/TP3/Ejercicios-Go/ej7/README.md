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

### Solución 2 — Con canales (como está en `ej7b/main.go`)

Una goroutine "coordinadora" gobierna quién puede entrar, usando canales de pedido/permiso/fin para cada rol y `select` con **guardas dinámicas** (canales que se ponen en `nil` para "apagar" ese case cuando no corresponde):

```go
func coordinador(pedirLeer, finLeer, pedirEscribir, finEscribir chan int, permisoLeer, permisoEscribir chan bool) {
    lectoresActivos := 0
    escribiendo := false

    for {
        var lecturaGuard, escrituraGuard chan int

        if !escribiendo {
            lecturaGuard = pedirLeer // solo se puede entrar a leer si nadie escribe
        }
        if lectoresActivos == 0 && !escribiendo {
            escrituraGuard = pedirEscribir // solo se puede escribir si no hay nadie leyendo ni escribiendo
        }

        select {
        case id := <-lecturaGuard:
            lectoresActivos++
            permisoLeer <- true
        case id := <-finLeer:
            lectoresActivos--
        case id := <-escrituraGuard:
            escribiendo = true
            permisoEscribir <- true
        case id := <-finEscribir:
            escribiendo = false
        }
    }
}
```

Lectores y escritores piden permiso y avisan cuando terminan:

```go
pedirLeer <- id
<-permisoLeer
// ... leer ...
finLeer <- id
```

> A diferencia de un coordinador que serializa todo (un request a la vez), acá el `select` con guardas sí permite varios lectores concurrentes: mientras `escribiendo` sea `false`, todos los pedidos de lectura se aceptan sin esperar. Solo se bloquea la lectura cuando hay un escritor activo, y solo se bloquea la escritura cuando hay lectores o un escritor activo. Es más código que la versión con `RWMutex`, pero replica el mismo comportamiento sin locks, a mano.

---

## Conceptos de Teoría

**Problema Lectores/Escritores:** problema clásico de concurrencia. Las invariantes son: (1) múltiples lectores pueden leer simultáneamente sin interferencia entre sí; (2) un escritor necesita acceso exclusivo — ni lectores ni otros escritores pueden acceder mientras escribe.

**`sync.RWMutex`:** variante del mutex que distingue lecturas de escrituras. `RLock()`/`RUnlock()` permiten concurrencia entre lectores; `Lock()`/`Unlock()` dan exclusión total. Es la abstracción directa del problema en Go y la solución más eficiente para cargas con muchas lecturas.

**Modelo CSP (Communicating Sequential Processes):** alternativa al modelo de memoria compartida. En lugar de proteger datos con locks, los goroutines se comunican enviando mensajes por canales. El lema de Go: *"No comuniques compartiendo memoria; compartí memoria comunicando"*.

**Goroutine coordinadora:** patrón CSP donde un único goroutine "posee" el estado (`lectoresActivos`, `escribiendo`) y decide quién entra. No hace falta que sea un cuello de botella para las lecturas: con guardas dinámicas en el `select`, el coordinador solo serializa las decisiones de *entrada/salida* (que son rápidas), mientras que la lectura en sí ocurre en paralelo en cada goroutine lectora.

**Guardas dinámicas (`nil` channels en `select`):** un `case` de un `select` que lee de un canal `nil` nunca está listo, así que ese `case` queda "apagado". Poniendo o sacando el canal real según el estado (`escribiendo`, `lectoresActivos`), el mismo `select` habilita o deshabilita cada camino en cada vuelta del loop.

**Canal de permiso (`permisoLeer`, `permisoEscribir`):** el coordinador no devuelve el dato en sí, solo un "adelante" — el lector/escritor hace su trabajo por su cuenta y después avisa por otro canal (`finLeer`/`finEscribir`) que terminó, para que el coordinador actualice su estado.
