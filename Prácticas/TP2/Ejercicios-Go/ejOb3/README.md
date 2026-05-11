# Ejercicio Obligatorio 3 — OptimumSlice ⭐ (a entregar)

## Enunciado

Implementar funciones para manejar slices de enteros que tienen muchas rachas de números repetidos. La estructura guarda cada número junto a su cantidad de ocurrencias consecutivas (compresión Run-Length Encoding).

```go
func New(s []int) OptimumSlice
func IsEmpty(o OptimumSlice) bool
func Len(o OptimumSlice) int
func FrontElement(o OptimumSlice) int
func LastElement(o OptimumSlice) int
func Average(o OptimumSlice) float64
func Occurrences(o OptimumSlice, element int) int
func IndexOf(o OptimumSlice, element int) int        // primera aparición
func Mode(o OptimumSlice) int                        // el número que más se repite
func Insert(o OptimumSlice, element int, position int) OptimumSlice
func SliceArray(o OptimumSlice) []int
```

**Ejemplo de Insert:**
- `o` representa `{3×5, 1×6, 23×6, 3×8, 7×1, 5×3}` → arreglo `[3,3,3,3,3,1,1,1,1,1,1,1,23,23,23,23,23,23,3,3,3,3,3,3,3,3,7,5,5,5]`
- `Insert(o, 9, 6)` → insertar el 9 en la posición 6

**Restricción:** no se puede convertir a `[]int`, insertar y volver a convertir.

---

## Lógica de resolución

### Estructura

```go
type run struct {
    valor     int
    ocurrencias int
}

type OptimumSlice []run
```

Cada `run` representa una racha: `{valor: 3, ocurrencias: 5}` significa cinco 3 seguidos.

### New — construir desde slice

```go
func New(s []int) OptimumSlice {
    if len(s) == 0 { return OptimumSlice{} }
    result := OptimumSlice{{s[0], 1}}
    for _, v := range s[1:] {
        last := &result[len(result)-1]
        if v == last.valor {
            last.ocurrencias++
        } else {
            result = append(result, run{v, 1})
        }
    }
    return result
}
```

### Len — longitud total del arreglo expandido

```go
func Len(o OptimumSlice) int {
    total := 0
    for _, r := range o {
        total += r.ocurrencias
    }
    return total
}
```

### FrontElement y LastElement

```go
func FrontElement(o OptimumSlice) int { return o[0].valor }
func LastElement(o OptimumSlice) int  { return o[len(o)-1].valor }
```

### Average

```go
func Average(o OptimumSlice) float64 {
    var suma float64
    for _, r := range o {
        suma += float64(r.valor) * float64(r.ocurrencias)
    }
    return suma / float64(Len(o))
}
```

### Occurrences — total de ocurrencias de un valor

```go
func Occurrences(o OptimumSlice, element int) int {
    total := 0
    for _, r := range o {
        if r.valor == element {
            total += r.ocurrencias
        }
    }
    return total
}
```

### IndexOf — primera aparición

```go
func IndexOf(o OptimumSlice, element int) int {
    pos := 0
    for _, r := range o {
        if r.valor == element { return pos }
        pos += r.ocurrencias
    }
    return -1
}
```

### Mode — valor más repetido

```go
func Mode(o OptimumSlice) int {
    best := o[0]
    for _, r := range o[1:] {
        if r.ocurrencias > best.ocurrencias {
            best = r
        }
    }
    return best.valor
}
```

### Insert — la operación más compleja

La idea: localizar en qué `run` cae la `position` y dividirlo si hace falta.

```go
func Insert(o OptimumSlice, element int, position int) OptimumSlice {
    pos := 0
    for i, r := range o {
        if pos+r.ocurrencias > position {
            // la posición cae dentro del run i
            offset := position - pos   // cuántos elementos del run van antes del nuevo

            // ¿el elemento insertado es igual al valor del run?
            if element == r.valor {
                o[i].ocurrencias++
                return o
            }

            // Dividir el run en: [0..offset) | nuevo | [offset..r.ocurrencias)
            antes := run{r.valor, offset}
            nuevo := run{element, 1}
            despues := run{r.valor, r.ocurrencias - offset}

            result := make(OptimumSlice, 0, len(o)+2)
            result = append(result, o[:i]...)
            if antes.ocurrencias > 0    { result = append(result, antes) }
            result = append(result, nuevo)
            if despues.ocurrencias > 0  { result = append(result, despues) }
            result = append(result, o[i+1:]...)
            return result
        }
        pos += r.ocurrencias
    }
    // position >= Len(o): agregar al final
    last := &o[len(o)-1]
    if last.valor == element {
        last.ocurrencias++
        return o
    }
    return append(o, run{element, 1})
}
```

### SliceArray — expandir de vuelta a []int

```go
func SliceArray(o OptimumSlice) []int {
    result := make([]int, 0, Len(o))
    for _, r := range o {
        for i := 0; i < r.ocurrencias; i++ {
            result = append(result, r.valor)
        }
    }
    return result
}
```

### Casos borde a tener en cuenta

- **Insert en la frontera entre dos runs**: si se inserta entre el run `i` y el run `i+1`, `offset == r.ocurrencias`, el `antes` toma todo el run y `despues` tiene 0 elementos → no se agrega.
- **Insert de un valor igual al vecino**: si el elemento insertado es igual al run antes o después del punto de corte, conviene fusionar runs para mantener la compresión.
- **Insert al final**: cuando `position >= Len(o)`.

> Este ejercicio es el que se entrega. La clave es entender que cada operación trabaja sobre runs comprimidos, sin expandir nunca el arreglo completo. `Insert` es la más delicada: hay que ubicar el run correcto, calcular el offset dentro de ese run y dividirlo correctamente.
