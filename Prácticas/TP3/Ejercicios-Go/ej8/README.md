# Ejercicio 8 — Select con tres canales

## Enunciado

Realizar un programa que utilice `select` para recibir valores desde **tres canales diferentes**. Cada canal recibe una secuencia de números enteros. El programa debe:

- Recibir un valor de cada canal usando `select` para determinar cuál tiene un valor disponible.
- Continuar hasta haber recibido todos los valores enviados a cada canal.
- Al final, mostrar el total de valores recibidos desde cada canal.

*Objetivo: select*

---

## Lógica de resolución

### Estructura general

```go
ch1 := make(chan int)
ch2 := make(chan int)
ch3 := make(chan int)
```

Lanzar una goroutine por canal que envíe una secuencia de valores y luego cierre el canal:

```go
go func() {
    for _, v := range []int{1, 2, 3} { ch1 <- v }
    close(ch1)
}()
// igual para ch2 y ch3
```

### Loop con select, poniendo el canal en nil al cerrarse (como está en `main.go`)

```go
count1, count2, count3 := 0, 0, 0

for ch1 != nil || ch2 != nil || ch3 != nil {
    select {
    case v, ok := <-ch1:
        if !ok { ch1 = nil; continue }
        fmt.Println("ch1:", v)
        count1++
    case v, ok := <-ch2:
        if !ok { ch2 = nil; continue }
        fmt.Println("ch2:", v)
        count2++
    case v, ok := <-ch3:
        if !ok { ch3 = nil; continue }
        fmt.Println("ch3:", v)
        count3++
    }
}

fmt.Printf("Totales - ch1: %d, ch2: %d, ch3: %d\n", count1, count2, count3)
```

> Cuando un canal cerrado se usa en `select`, `ok` vale `false`. Poner el canal en `nil` lo excluye por completo del `select` (un canal `nil` nunca está listo). Esto es importante: si solo se usara un flag booleano para cortar el loop *sin* tocar el canal, ese `case` seguiría "ganando" el `select` en cada vuelta — leer de un canal cerrado no bloquea, devuelve el cero del tipo con `ok=false` inmediatamente — y el programa quedaría girando en banda sobre ese canal hasta que los otros dos también cierren. Poniendo `chN = nil` se evita ese busy-spin.

---

## Conceptos de Teoría

**`select`:** permite esperar sobre múltiples operaciones de canal a la vez. Cuando hay más de un caso listo simultáneamente, Go elige uno al azar — no hay prioridad implícita.

**Detección de canal cerrado:** la forma `v, ok := <-ch` permite saber si el canal fue cerrado (`ok == false`). Leer de un canal cerrado devuelve el valor cero del tipo sin bloquear, así que sin este chequeo el loop sería infinito.

**Canal `nil` en `select`:** asignar `nil` a un canal lo excluye de todos los casos de `select` — un canal nil bloquea para siempre. Es la forma idiomática de "desactivar" un caso sin reestructurar el select.

**Fan-in:** patrón donde múltiples canales de entrada se multiplexan en un único receptor usando `select`. Permite procesar eventos de distintas fuentes sin saber de antemano cuál tendrá un valor disponible primero.
