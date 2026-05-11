# Ejercicio 7 — Conteo de Caracteres

## Enunciado

Leer una secuencia de caracteres que finaliza con CR (Enter) e informar la cantidad de letras, números y caracteres especiales.

**a)** Contar mayúsculas y minúsculas por separado.

**b)** Contar de forma separada las ocurrencias de cada dígito decimal usando un `Map`.

---

## Lógica de resolución

### Lectura de caracteres (runas)

```go
var letras, numeros, especiales int

scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
linea := scanner.Text()

for _, r := range linea {
    // r es de tipo rune
}
```

En Go, iterar un `string` con `range` da runas (`rune` = `int32`), no bytes. Es la forma correcta de manejar caracteres Unicode.

### Clasificación de runas

```go
import "unicode"

for _, r := range linea {
    switch {
    case unicode.IsLetter(r):
        letras++
    case unicode.IsDigit(r):
        numeros++
    default:
        especiales++
    }
}
```

### Parte a — mayúsculas y minúsculas

```go
var mayusculas, minusculas int

for _, r := range linea {
    switch {
    case unicode.IsUpper(r):
        mayusculas++
    case unicode.IsLower(r):
        minusculas++
    case unicode.IsDigit(r):
        numeros++
    default:
        especiales++
    }
}
```

### Parte b — conteo de dígitos con Map

```go
digitos := make(map[rune]int)

for _, r := range linea {
    if unicode.IsDigit(r) {
        digitos[r]++
    }
}

for d := '0'; d <= '9'; d++ {
    fmt.Printf("'%c': %d\n", d, digitos[d])
}
```

> Si una clave no existe en el map, Go retorna el **zero value** del tipo valor (0 para `int`), así que `digitos[r]++` funciona incluso si `r` aparece por primera vez.
