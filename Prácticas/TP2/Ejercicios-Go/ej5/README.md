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

## Lógica de resolución (como está en `main.go`)

```go
const N = 3
type Vector [N]float64

func Initialize(v *Vector, f float64) {
    for i := 0; i < N; i++ {
        v[i] = f
    }
}

// Retorna un vector nuevo, los originales no cambian
func Sum(v1, v2 Vector) Vector {
    var v Vector
    for i := 0; i < N; i++ {
        v[i] = v1[i] + v2[i]
    }
    return v
}

// El resultado se guarda en v1 (pasaje por referencia con *)
func SumInPlace(v1 *Vector, v2 Vector) {
    for i := 0; i < N; i++ {
        v1[i] = v1[i] + v2[i]
    }
}
```

## Observaciones

- Los arreglos en Go se pasan por valor (se copian). Por eso `Initialize` y `SumInPlace` necesitan `*Vector`: sin el puntero, los cambios quedarían en una copia local y no se verían afuera.
- `Sum` en cambio recibe todo por valor a propósito: no toca ni `v1` ni `v2`, arma un `Vector` nuevo y lo devuelve. Por eso no necesita puntero.
- `N` es una constante y `Vector` un arreglo (no slice), así que el tamaño queda fijo en tiempo de compilación — no se puede tener un Vector de tamaño variable con este tipo.
