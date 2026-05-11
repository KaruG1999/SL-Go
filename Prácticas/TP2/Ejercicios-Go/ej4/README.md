# Ejercicio 4 — Tres Sucesiones

## Enunciado

Leer tres sucesiones de números enteros, cada una de longitud N (constante): `x1…xn`, `y1…yn`, `z1…zn`, almacenarlas en sus arreglos y calcular:

```
resultado = (Σx / Πy) + máx(z) - mín(z)
```

Usar funciones para la productoria, la sumatoria y el máximo-mínimo. La función `maxmin` retorna el máximo y el mínimo de la serie.

---

## Lógica de resolución

### Constante y tipos

```go
const N = 5

type Serie [N]int
```

### Funciones auxiliares

```go
func sumatoria(s Serie) int {
    total := 0
    for _, v := range s {
        total += v
    }
    return total
}

func productoria(s Serie) int {
    total := 1
    for _, v := range s {
        total *= v
    }
    return total
}

func maxmin(s Serie) (int, int) {
    max, min := s[0], s[0]
    for _, v := range s[1:] {
        if v > max { max = v }
        if v < min { min = v }
    }
    return max, min
}
```

### Cálculo principal

```go
func main() {
    var x, y, z Serie
    // leer x, y, z ...

    sumX  := sumatoria(x)
    prodY := productoria(y)
    maxZ, minZ := maxmin(z)

    resultado := float64(sumX)/float64(prodY) + float64(maxZ-minZ)
    fmt.Printf("Resultado: %.4f\n", resultado)
}
```

> Se necesita casting a `float64` para la división, ya que `sumX/prodY` en enteros truncaría el resultado. El ejercicio pone foco en el **tipado fuerte de Go**: no se puede mezclar `int` y `float64` sin conversión explícita.
