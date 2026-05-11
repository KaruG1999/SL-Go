# Ejercicio 6 — Operaciones con Slices

## Enunciado

Implementar dos funciones sobre slices de enteros:

```go
func Sum(a, b []int) []int   // retorna slice con la suma elemento a elemento; tamaño = mínimo de las longitudes
func Avg(a []int) int        // retorna el promedio entero de los elementos
```

**a)** Re-implementar `Avg` para que retorne un `float64`.

---

## Lógica de resolución

### Sum

```go
func Sum(a, b []int) []int {
    n := len(a)
    if len(b) < n {
        n = len(b)
    }
    result := make([]int, n)
    for i := 0; i < n; i++ {
        result[i] = a[i] + b[i]
    }
    return result
}
```

> `make([]int, n)` crea un slice de longitud `n` con todos los elementos en 0. Los slices son referencias en Go: la función puede retornar un slice recién creado sin problemas.

### Avg — versión entera

```go
func Avg(a []int) int {
    if len(a) == 0 {
        return 0
    }
    total := 0
    for _, v := range a {
        total += v
    }
    return total / len(a)
}
```

### Avg — versión float (parte a)

```go
func AvgFloat(a []int) float64 {
    if len(a) == 0 {
        return 0.0
    }
    total := 0
    for _, v := range a {
        total += v
    }
    return float64(total) / float64(len(a))
}
```

### Uso

```go
func main() {
    a := []int{1, 2, 3, 4, 5}
    b := []int{10, 20, 30}

    fmt.Println("Sum:", Sum(a, b))       // [11 22 33]
    fmt.Println("Avg:", Avg(a))          // 3
    fmt.Println("AvgFloat:", AvgFloat(a)) // 3.0
}
```

> Los slices en Go no tienen tamaño fijo (a diferencia de los arreglos). Se pasan como referencias al header del slice (puntero + len + cap), por lo que las funciones ven los mismos datos sin copiarlos.
