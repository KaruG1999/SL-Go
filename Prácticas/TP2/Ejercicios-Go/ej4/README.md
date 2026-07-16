# Ejercicio 4 — Tres Sucesiones

## Enunciado

Leer tres sucesiones de números enteros, cada una de longitud N (constante): `x1…xn`, `y1…yn`, `z1…zn`, almacenarlas en sus respectivos arreglos y calcular:

```
R = (Σ 1/xi − Π zi³) × maxmin(yi)
```

Usar funciones para la productoria, la sumatoria y el máximo-mínimo. La función `maxmin` retorna el máximo y el mínimo de la serie, y luego ambos se multiplican por el resto de la ecuación.

---

## Lógica de resolución (como está en `main.go`)

```go
const N = 5

func maxmin(serie [N]int) (int, int) {
    max := serie[0]
    min := serie[0]
    for _, v := range serie {
        if v > max { max = v }
        if v < min { min = v }
    }
    return max, min
}

func productoria(serie [N]int) int {
    p := 1
    for _, v := range serie {
        p *= (v * v * v) // zi³
    }
    return p
}

func sumatoria(serie [N]int) float64 {
    var s float64
    for _, v := range serie {
        s += 1.0 / float64(v) // 1/xi, necesita float para no truncar
    }
    return s
}

func main() {
    var x, y, z [N]int
    // leer x, y, z ...

    s := sumatoria(x)        // Σ 1/xi
    p := productoria(z)      // Π zi³
    max, min := maxmin(y)    // maxmin de y

    resto := s - float64(p)
    resultado := resto * float64(max) * float64(min)

    fmt.Printf("El resultado es: %.4f\n", resultado)
}
```

- `x` va a la sumatoria (de recíprocos), `z` a la productoria (al cubo), `y` al maxmin. El enunciado los pide en ese orden.

## Observaciones

- Todo el ejercicio pasa por el tipado fuerte: `sumatoria` devuelve `float64` para no perder la parte decimal de `1/xi`; el resto de las cuentas fuerza casting explícito a `float64` antes de mezclarlas.
- Si algún valor de `x` es 0, `1.0/float64(0)` da `+Inf` en vez de explotar (es división de floats, no de enteros). No hay chequeo para ese caso — vale la pena mencionarlo si preguntan qué pasa con un 0 en la serie.
- `productoria` eleva al cubo cada `z` antes de multiplicar: con series de N=5 y valores grandes esto puede desbordar un `int` fácilmente, no hay protección contra overflow.
- El resultado final multiplica el resto por `max` y por `min` por separado (no por `max*min` en conjunto), pero da lo mismo matemáticamente.
