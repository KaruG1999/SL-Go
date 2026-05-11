# Ejercicio 8 — Conversión de Base

## Enunciado

Implementar la función:

```go
func Convert(v int, b int) string
```

Convierte el entero `v` a un string en su representación en base `b` (entre 2 y 36). Los dígitos disponibles son: `"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"`.

Ejemplo: `Convert(23, 2)` retorna `"10111"`.

**b)** Contemplar que `v` puede ser negativo, usando el símbolo `-` para la representación.

---

## Lógica de resolución

### Idea base

La conversión a base b se realiza obteniendo el resto de la división sucesiva por b, y los restos leídos de atrás hacia adelante forman el número. En código se puede hacer de forma iterativa o recursiva.

### Versión iterativa

```go
const digitos = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Convert(v int, b int) string {
    if b < 2 || b > 36 {
        panic("base inválida")
    }
    if v == 0 {
        return "0"
    }

    negativo := v < 0
    if negativo {
        v = -v
    }

    result := ""
    for v > 0 {
        result = string(digitos[v%b]) + result  // preponer el dígito
        v /= b
    }

    if negativo {
        result = "-" + result
    }
    return result
}
```

### Versión recursiva

```go
func convertRec(v int, b int) string {
    if v == 0 {
        return ""
    }
    return convertRec(v/b, b) + string(digitos[v%b])
}

func Convert(v int, b int) string {
    if v == 0 { return "0" }
    if v < 0  { return "-" + convertRec(-v, b) }
    return convertRec(v, b)
}
```

### Uso

```go
func main() {
    fmt.Println(Convert(23, 2))   // "10111"
    fmt.Println(Convert(255, 16)) // "FF"
    fmt.Println(Convert(-10, 2))  // "-1010"
    fmt.Println(Convert(0, 10))   // "0"
}
```

> `string(digitos[v%b])` convierte el byte en posición `v%b` del string `digitos` a un string de un carácter. El truco de **preponer** (`result = char + result`) construye el string en el orden correcto sin necesidad de invertirlo al final.
