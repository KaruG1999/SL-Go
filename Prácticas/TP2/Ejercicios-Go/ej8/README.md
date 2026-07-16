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

## Lógica de resolución (como está en `base/` y `b/`)

```go
const digitos = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Convert(v int, b int) string {
    if v == 0 { return "0" }
    if b < 2 || b > 36 { return "Base inválida" }

    resultado := ""
    for v > 0 {
        resto := v % b
        digito := digitos[resto]
        resultado = string(digito) + resultado // se antepone
        v = v / b
    }
    return resultado
}
```

Parte b agrega el manejo del signo:

```go
negativo := false
if v < 0 {
    negativo = true
    v = -v
}
// ... mismo loop ...
if negativo { resultado = "-" + resultado }
```

## Observaciones

- El chequeo de base válida originalmente decía `b > 32`, pero `digitos` tiene 36 caracteres (10 dígitos + 26 letras), así que las bases 33 a 36 son válidas. Ya está corregido a `b > 36` en el código — ojo si en algún momento se vuelve a tocar ese límite.

- El truco de anteponer el dígito (`resultado = digito + resultado`) evita tener que invertir el string al final.
- `digitos[resto]` da un `byte`, por eso hay que convertirlo con `string(...)` antes de concatenar.
- El caso `v == 0` se resuelve aparte porque el `for v > 0` nunca entraría y devolvería un string vacío en vez de `"0"`.
- No se usa `strings.Builder` acá, y cambiarlo tal cual no mejora nada: Builder es eficiente para *appendear*, pero el algoritmo *prepende* el dígito en cada vuelta (`digito + resultado`), que es justo lo que Builder no acelera. Para aprovecharlo de verdad hay que appendear los dígitos en el orden que salen (unidades primero) y invertir el string una sola vez al final — queda comentado como alternativa en el código.
