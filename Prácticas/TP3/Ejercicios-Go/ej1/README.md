# Ejercicio 1 — Map genérico

## Enunciado

Realizar un programa que haga **dos usos con distintos tipos** del siguiente tipo genérico:

```go
type Map[K comparable, V any] map[K]V
```

*Objetivo: tipos genéricos*

---

## Lógica de resolución

### Definir el tipo

```go
type Map[K comparable, V any] map[K]V
```

El parámetro `K` debe ser `comparable` (condición para ser clave de map en Go). `V` puede ser cualquier tipo.

### Uso 1 — clave string, valor int (como en `main.go`)

```go
inventario := make(Map[string, int])
inventario["Arroz"] = 15
inventario["Fideos"] = 4
```

### Uso 2 — clave int, valor string

```go
usuarios := make(Map[int, string])
usuarios[1024] = "Karen"
usuarios[2048] = "Alejandro"
```

> El objetivo es ver que el mismo tipo genérico `Map[K, V]` funciona igual independientemente de los tipos concretos que se le pasen, sin duplicar código.

---

## Nota: tipos genéricos en Go

Los genéricos permiten escribir código que funciona con distintos tipos sin repetirlo. Se declaran con **parámetros de tipo** entre corchetes `[T constraint]`.

### Declaración de tipo genérico

```go
// tipo genérico simple
type Stack[T any] []T

// con múltiples parámetros
type Pair[K comparable, V any] struct {
    Key   K
    Value V
}
```

### Función genérica

```go
func Contiene[T comparable](slice []T, elem T) bool {
    for _, v := range slice {
        if v == elem { return true }
    }
    return false
}
```

### Uso (instanciación)

```go
// el tipo se pasa entre corchetes
p := Pair[string, int]{Key: "edad", Value: 30}

// en funciones, Go puede inferirlo automáticamente
Contiene([]int{1, 2, 3}, 2)           // infiere T = int
Contiene([]string{"a", "b"}, "a")     // infiere T = string
```

### Constraints más comunes

| Constraint | Significado |
|---|---|
| `any` | cualquier tipo (alias de `interface{}`) |
| `comparable` | tipos que soportan `==` y `!=` (necesario para claves de map) |
| `cmp.Ordered` | tipos ordenables con `<`, `>` (int, float, string) |
