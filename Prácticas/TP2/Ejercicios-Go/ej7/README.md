# Ejercicio 7 — Conteo de Caracteres

## Enunciado

Leer una secuencia de caracteres que finaliza con CR (Enter) e informar la cantidad de letras, números y caracteres especiales.

**a)** Contar mayúsculas y minúsculas por separado.

**b)** Contar de forma separada las ocurrencias de cada dígito decimal usando un `Map`.

---

## Lógica de resolución (como está en `base/`, `a/` y `b/`)

### Base

```go
var letras, numeros, especiales int

for _, caracter := range texto {
    if caracter == '\n' || caracter == '\r' {
        continue // no contar el Enter que termina la entrada
    }
    if unicode.IsLetter(caracter) {
        letras++
    } else if unicode.IsDigit(caracter) {
        numeros++
    } else {
        especiales++
    }
}
```

### Parte a — mayúsculas y minúsculas

```go
var mayusculas, minusculas, numeros, especiales int

for _, c := range texto {
    switch {
    case unicode.IsUpper(c):
        mayusculas++
    case unicode.IsLower(c):
        minusculas++
    case unicode.IsDigit(c):
        numeros++
    default:
        especiales++
    }
}
```

### Parte b — conteo de dígitos con map

```go
digitos := make(map[rune]int)

for _, c := range texto {
    if unicode.IsDigit(c) {
        digitos[c]++
    }
}

for d := '0'; d <= '9'; d++ {
    fmt.Printf("'%c': %d\n", d, digitos[d])
}
```

## Observaciones

- `range` sobre un string da runas, no bytes — importante si el texto tiene tildes o ñ, cada carácter se cuenta como una unidad aunque ocupe más de un byte.
- La lectura con `bufio.Scanner` ya corta en el salto de línea, pero igual se filtra `\n`/`\r` explícitamente por las dudas.
- En la parte b, si una clave todavía no apareció en el map, `digitos[d]` devuelve 0 (zero value), así que el `for` final imprime los 10 dígitos aunque algunos no hayan aparecido nunca.
