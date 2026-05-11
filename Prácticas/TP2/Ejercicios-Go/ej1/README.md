# Ejercicio 1 — Temperaturas

## Enunciado

Las temperaturas de los pacientes de un hospital se dividen en 3 grupos:
- **Alta**: mayor de 37.5
- **Normal**: entre 36 y 37.5
- **Baja**: menor de 36

Se deben leer 10 temperaturas e informar el porcentaje de pacientes de cada grupo, y el promedio entero entre la temperatura máxima y la mínima. Resolver cargando primero todos los valores usando un arreglo y usando variables escalares como acumuladores/contadores. El programa debe poder leer desde archivo:

```bash
go run ej1.go < input_ej1.txt
```

**a)** Volver a resolver usando un arreglo o Map de 3 posiciones para acumular cada grupo.

**b)** Incluir un grupo de valores incorrectos (mayores a 50° o menores a 20°).

**c)** Escribir una función que convierta de Celsius a Fahrenheit usando nuevos tipos y aplicarla al arreglo de valores leídos. Fórmula: `F = (C × 9/5) + 32`.

---

## Lógica de resolución

### Estructura de datos

```go
const N = 10

type Celsius    float64
type Fahrenheit float64

var temps [N]Celsius
```

Usar un arreglo de tamaño fijo `[N]Celsius` para almacenar las lecturas antes de procesarlas.

### Lectura

```go
for i := 0; i < N; i++ {
    fmt.Scan(&temps[i])
}
```

La entrada estándar redirigida desde archivo funciona igual que el teclado con `fmt.Scan`.

### Clasificación (parte base — variables escalares)

```go
var alta, normal, baja int
var maxT, minT Celsius = temps[0], temps[0]

for _, t := range temps {
    if t > 37.5      { alta++ }
    else if t >= 36  { normal++ }
    else             { baja++ }
    if t > maxT { maxT = t }
    if t < minT { minT = t }
}

promedio := int((maxT + minT) / 2)
fmt.Printf("Alta: %.1f%%, Normal: %.1f%%, Baja: %.1f%%\n",
    float64(alta)/N*100, float64(normal)/N*100, float64(baja)/N*100)
fmt.Printf("Promedio (max+min)/2: %d\n", promedio)
```

### Parte a — con Map

```go
grupos := map[string]int{"alta": 0, "normal": 0, "baja": 0}

for _, t := range temps {
    switch {
    case t > 37.5: grupos["alta"]++
    case t >= 36:  grupos["normal"]++
    default:       grupos["baja"]++
    }
}
```

### Parte b — grupo de valores incorrectos

Agregar al switch un caso adicional: `case t > 50 || t < 20: grupos["error"]++`.

### Parte c — conversión de tipos

```go
func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

var fahrenheits [N]Fahrenheit
for i, t := range temps {
    fahrenheits[i] = CToF(t)
}
```

> El tipado fuerte de Go no permite operar directamente entre `Celsius` y `Fahrenheit` sin conversión explícita — eso es exactamente lo que el ejercicio quiere demostrar.
