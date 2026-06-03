# Ejercicio Obligatorio 3 — Scheduler con prioridades

## Enunciado

> **Este es el ejercicio a entregar.**

Desarrollar un programa que implemente un sistema de planificación (scheduler) utilizando **5 goroutines** (el main y 4 workers). El programa genera números enteros aleatorios, cada uno con una prioridad aleatoria entre 0 y 3 (0 = más alta, 3 = más baja).

### Reglas del scheduler

**a)** Procesar los datos en orden de prioridad: 0, luego 1, 2 y 3.

**b)** Mientras haya datos de prioridad 0, procesarlos exclusivamente.

**c)** Si no hay datos de prioridad 0 y hay goroutines disponibles, asignarles datos de menor prioridad.

**d)** Una vez sin datos de prioridad 0, pasar a prioridad 1, y así sucesivamente.

**e)** El main genera los datos con prioridades aleatorias y distribuye el trabajo a las goroutines disponibles manteniendo el orden por prioridad.

**f)** Procesamiento por prioridad:
- **Prioridad 0**: sumar los dígitos del número → guardar en `prioridad0.txt` como `(0, resultado)`.
- **Prioridad 1**: invertir los dígitos del número → guardar en `prioridad1.txt` como `(1, resultado)`.
- **Prioridad 2**: multiplicar el número por 10 → imprimir en consola.
- **Prioridad 3**: acumular los números → mostrar el acumulado en consola cada vez que se procesa uno.

*Tip*: usar `math/rand` para generar números aleatorios.

*Objetivo: goroutines, channels, select, prioridades*

---

## Lógica de resolución

### Estructura de un dato

```go
type Dato struct {
    numero    int
    prioridad int
}
```

### Canales por prioridad

Usar un canal por nivel de prioridad para que el scheduler pueda elegir siempre el de mayor prioridad disponible:

```go
colas := [4]chan Dato{
    make(chan Dato, 100),
    make(chan Dato, 100),
    make(chan Dato, 100),
    make(chan Dato, 100),
}
```

### Scheduler con select priorizando

El select de Go no garantiza orden entre casos. Para respetar prioridades, intentar primero la cola 0 con un select no bloqueante:

```go
func scheduler(colas [4]chan Dato, workers chan Dato, done chan struct{}) {
    for {
        select {
        case d := <-colas[0]:
            workers <- d
            continue
        default:
        }
        select {
        case d := <-colas[1]:
            workers <- d
            continue
        default:
        }
        // ... prioridades 2 y 3 ...
        select {
        case <-done:
            return
        default:
            runtime.Gosched() // ceder CPU brevemente
        }
    }
}
```

### Workers

```go
func worker(id int, ch chan Dato, acum *int, mu *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    for d := range ch {
        switch d.prioridad {
        case 0:
            res := sumarDigitos(d.numero)
            escribirArchivo("prioridad0.txt", fmt.Sprintf("(0, %d)\n", res))
        case 1:
            res := invertirDigitos(d.numero)
            escribirArchivo("prioridad1.txt", fmt.Sprintf("(1, %d)\n", res))
        case 2:
            fmt.Printf("P2: %d × 10 = %d\n", d.numero, d.numero*10)
        case 3:
            mu.Lock()
            *acum += d.numero
            fmt.Printf("P3: acumulado = %d\n", *acum)
            mu.Unlock()
        }
    }
}
```

### Funciones auxiliares

```go
func sumarDigitos(n int) int {
    suma := 0
    for n > 0 { suma += n % 10; n /= 10 }
    return suma
}

func invertirDigitos(n int) int {
    inv := 0
    for n > 0 { inv = inv*10 + n%10; n /= 10 }
    return inv
}
```

### Main: generación y distribución

```go
func main() {
    rand.Seed(time.Now().UnixNano())
    // generar N datos aleatorios
    for i := 0; i < N; i++ {
        d := Dato{numero: rand.Intn(10000), prioridad: rand.Intn(4)}
        colas[d.prioridad] <- d
    }
    // lanzar scheduler y 4 workers
    // esperar con WaitGroup
}
```
