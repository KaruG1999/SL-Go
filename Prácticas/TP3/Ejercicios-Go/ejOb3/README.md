# Ejercicio Obligatorio 3 — Planificador con prioridades

## Enunciado

> **Este es el ejercicio a entregar.**

Desarrollar un programa que implemente un sistema de planificación (en inglés, "scheduler" — en el código la función se llama `planificador`) utilizando **5 goroutines** (el main y 4 trabajadores). El programa genera números enteros aleatorios, cada uno con una prioridad aleatoria entre 0 y 3 (0 = más alta, 3 = más baja).

### Reglas del planificador

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

Usar un canal por nivel de prioridad para que el planificador pueda elegir siempre el de mayor prioridad disponible:

```go
colas := [4]chan Dato{
    make(chan Dato, 100),
    make(chan Dato, 100),
    make(chan Dato, 100),
    make(chan Dato, 100),
}
```

### Planificador con select priorizando (como está en `main.go`)

El select de Go no garantiza orden entre casos. Para respetar prioridades, se revisa la cola 0 primero, y solo si no hay nada ahí se pasa a la 1, luego la 2, luego la 3 — con select no bloqueante (`default`) en cada una:

```go
// trabajadores es "chan<- Dato": el "<-" después de chan significa que desde
// esta función el canal solo se puede escribir, no leer.
func planificador(colas [4]chan Dato, trabajadores chan<- Dato, total int) {
    despachados := 0
    for despachados < total {
        enviado := false
        for p := 0; p < 4; p++ {
            // select con "default": si colas[p] no tiene nada listo en este
            // instante, no se queda esperando, cae al default y sigue con
            // la próxima prioridad. Sin el default, quedaría trabado ahí.
            select {
            case d := <-colas[p]:
                trabajadores <- d
                despachados++
                enviado = true
            default:
            }
            if enviado {
                break // ya despachó de la cola de mayor prioridad disponible, vuelve a arrancar desde p=0
            }
        }
        if !enviado {
            // no había nada en ninguna cola: runtime.Gosched() le cede el
            // turno a otras goroutines en vez de reintentar en loop cerrado
            runtime.Gosched()
        }
    }
    close(trabajadores)
}
```

En vez de un canal `done` aparte, el planificador sabe que terminó contando cuántos datos ya despachó (`despachados`) contra el total esperado (`total`) — como se generan todos los datos de antemano en `main`, ese total se conoce desde el principio.

### Trabajadores

Los archivos de prioridad 0 y 1 ya vienen abiertos desde `main` (una sola vez para todo el programa) y se los pasa como parámetro:

```go
// ch es "<-chan Dato": el "<-" antes de chan significa que desde acá el
// canal solo se puede leer, no escribir (al revés que en planificador).
func trabajador(id int, ch <-chan Dato, acum *int, mu *sync.Mutex, wg *sync.WaitGroup, archP0, archP1 *os.File) {
    defer wg.Done()
    for d := range ch {
        switch d.prioridad {
        case 0:
            res := sumarDigitos(d.numero)
            escribirArchivo(archP0, fmt.Sprintf("(0, %d)\n", res))
        case 1:
            res := invertirDigitos(d.numero)
            escribirArchivo(archP1, fmt.Sprintf("(1, %d)\n", res))
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

// escribe en un archivo ya abierto; el mutex protege la escritura porque
// varios trabajadores pueden escribir al mismo archivo al mismo tiempo
func escribirArchivo(f *os.File, texto string) {
    fileMu.Lock()
    defer fileMu.Unlock()
    f.WriteString(texto)
}
```

`ch <-chan Dato` y `trabajadores chan<- Dato` son el mismo tipo de canal (`chan Dato`) visto desde dos lados distintos: la flecha antes de `chan` (`<-chan`) marca "solo lectura", la flecha después (`chan<-`) marca "solo escritura". No cambian el comportamiento del canal en sí, son una anotación en la firma de la función para dejar claro qué hace cada una con el canal — no es algo que se use todo el tiempo, pero vale la pena reconocerlo si aparece.

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

    archP0, _ := os.Create("prioridad0.txt")
    defer archP0.Close()
    archP1, _ := os.Create("prioridad1.txt")
    defer archP1.Close()

    // generar los datos aleatorios y cargarlos en su cola de prioridad
    for i := 0; i < numDatos; i++ {
        d := Dato{numero: rand.Intn(10000), prioridad: rand.Intn(4)}
        colas[d.prioridad] <- d
    }

    // lanzar el planificador y los trabajadores, esperar con WaitGroup
}
```

---

## Devolución de los profesores

> Buena implementación en general, cumple con lo solicitado. Distribuye dinámicamente el trabajo mediante un canal compartido, mantiene el orden de llegada dentro de cada prioridad, utiliza correctamente mecanismos de sincronización (Mutex y WaitGroup) para proteger los recursos compartidos.
>
> Los archivos se cierran y abren por tarea, se podría hacer una sola vez.

**Corregido:** antes, `escribirArchivo` hacía `os.OpenFile` + `defer f.Close()` en cada llamada — es decir, una apertura y cierre por cada número de prioridad 0 o 1 procesado. Ahora `prioridad0.txt` y `prioridad1.txt` se abren una sola vez en `main` (con `os.Create`, que además reemplaza el `os.Remove` que se hacía antes para limpiar el archivo viejo) y los trabajadores reciben los `*os.File` ya abiertos como parámetro. Se cierran una sola vez al final, con `defer` en `main`.

El mutex (`fileMu`) se mantuvo igual: sigue haciendo falta porque varios trabajadores pueden escribir al mismo archivo al mismo tiempo, y eso no tiene que ver con cuántas veces se abre el archivo.
