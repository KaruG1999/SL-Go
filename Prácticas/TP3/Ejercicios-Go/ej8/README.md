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

### Loop con select y detección de cierre

```go
count1, count2, count3 := 0, 0, 0
open1, open2, open3 := true, true, true

for open1 || open2 || open3 {
    select {
    case v, ok := <-ch1:
        if !ok { open1 = false; break }
        fmt.Println("ch1:", v)
        count1++
    case v, ok := <-ch2:
        if !ok { open2 = false; break }
        fmt.Println("ch2:", v)
        count2++
    case v, ok := <-ch3:
        if !ok { open3 = false; break }
        fmt.Println("ch3:", v)
        count3++
    }
}

fmt.Printf("Totales — ch1: %d, ch2: %d, ch3: %d\n", count1, count2, count3)
```

> Cuando un canal cerrado se usa en `select`, la variable `ok` vale `false`. Hay que desactivar ese caso poniendo el canal en `nil` o usando flags booleanos para evitar recibirlo infinitamente.

### Alternativa: poner el canal en nil para excluirlo del select

```go
case v, ok := <-ch1:
    if !ok { ch1 = nil; break }
```

Un canal `nil` en `select` nunca está listo, así que queda excluido automáticamente.
