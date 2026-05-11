# Ejercicio 5 — Vector de Flotantes

## Enunciado

Definir el tipo vector de flotantes de tamaño fijo (constante) con las operaciones:

```go
func Initialize(v Vector, f float64)
func Sum(v1, v2 Vector) Vector
func SumInPlace(v1, v2 Vector)
```

`SumInPlace` guarda el resultado de la suma en el primer vector (a diferencia de `Sum` que retorna un nuevo vector).

---

## Lógica de resolución

### Tipo

```go
const SIZE = 5

type Vector [SIZE]float64
```

### Initialize

```go
func Initialize(v *Vector, f float64) {
    for i := range v {
        v[i] = f
    }
}
```

> En Go, los arreglos se pasan **por valor** (se copian). Para modificar el original desde dentro de la función se necesita un **puntero** `*Vector`. Sin el puntero, los cambios no se reflejan en el caller.

### Sum — retorna un nuevo vector

```go
func Sum(v1, v2 Vector) Vector {
    var result Vector
    for i := range v1 {
        result[i] = v1[i] + v2[i]
    }
    return result
}
```

Acá los arreglos se pasan por valor (copias). La función trabaja sobre copias y retorna un nuevo arreglo.

### SumInPlace — modifica el primer vector

```go
func SumInPlace(v1 *Vector, v2 Vector) {
    for i := range v1 {
        v1[i] += v2[i]
    }
}
```

### Uso

```go
func main() {
    var a, b Vector
    Initialize(&a, 1.0)
    Initialize(&b, 2.0)

    c := Sum(a, b)
    fmt.Println("Sum:", c)

    SumInPlace(&a, b)
    fmt.Println("SumInPlace a:", a)
}
```

> La diferencia clave entre `Sum` y `SumInPlace` es **semántica de propiedad**: `Sum` es pura (no modifica nada), `SumInPlace` muta el primer argumento. Para mutar, se pasa puntero.
