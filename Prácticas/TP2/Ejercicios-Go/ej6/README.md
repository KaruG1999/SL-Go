# Ejercicio 6 — Operaciones con Slices

## Enunciado

Implementar dos funciones sobre slices de enteros:

```go
func Sum(a, b []int) []int   // retorna slice con la suma elemento a elemento; tamaño = mínimo de las longitudes
func Avg(a []int) int        // retorna el promedio entero de los elementos
```

**a)** Re-implementar `Avg` para que retorne un `float64`.

---

## Lógica de resolución (como está en `main.go`)

```go
func Sum(s1, s2 []int) []int {
    n := len(s1)
    if len(s2) < n {
        n = len(s2)
    }
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = s1[i] + s2[i]
    }
    return res
}

func Prom(s []int) float64 {
    if len(s) == 0 {
        return 0
    }
    suma := 0
    for _, v := range s {
        suma += v
    }
    return float64(suma) / float64(len(s))
}
```

## Observaciones

- La función quedó directamente como `Prom` devolviendo `float64` (la parte a del enunciado). No está la versión base que pide `Avg` devolviendo `int` — si preguntan por esa, es la misma cuenta pero con `suma / len(s)` en enteros (trunca en vez de redondear).
- `make([]int, n)` crea el slice del tamaño del resultado ya en 0, sin necesidad de un `append` en el loop.
- `Sum` no valida que los slices tengan al menos un elemento: si ambos vienen vacíos, `n` es 0 y devuelve un slice vacío sin problema, no hace panic.
