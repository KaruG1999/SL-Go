# Go — Práctica 1: Sintaxis, tipos, control y E/S
`Seminario de Lenguajes opción Go`

> 🔗 Hilo previo: tipos básicos, variables, constantes, operadores, estructuras de control (`for`, `if`, `switch`), funciones, package `fmt`.
> ⚠️ Práctica del año anterior — puede tener cambios en la cursada 2026.

---

## Ejercicio 1 — Interpretado vs. compilado

**¿De qué concepto sale?** → Teoría general de lenguajes

**Resolución:**

| | Interpretado | Compilado |
|---|---|---|
| **Cómo funciona** | El código fuente se ejecuta línea por línea en tiempo real por un intérprete | El código fuente se traduce completo a código máquina antes de ejecutarse |
| **Velocidad** | Más lento en ejecución | Más rápido en ejecución |
| **Portabilidad** | El intérprete corre en cualquier plataforma | El ejecutable es específico a la arquitectura destino |
| **Detección de errores** | En tiempo de ejecución | En tiempo de compilación |
| **Ejemplos** | Python, JavaScript, Ruby, PHP | C, C++, Rust, Go |

**Go es compilado.** Genera un ejecutable nativo para la plataforma destino. Se compila con `go build` o se compila y ejecuta en un paso con `go run`.

---

## Ejercicio 2 — Hello World

**¿De qué concepto sale?** → Estructura básica de un programa Go + compilación

**Resolución:**

```go
// hola.go
package main

import "fmt"

func main() {
    fmt.Println("Hola Mundo!")
}
```

```bash
# Compilar y ejecutar en un paso
go run hola.go

# Compilar generando ejecutable
go build -o hola hola.go
./hola
```

**Por qué:** Todo ejecutable pertenece al package `main` y su punto de entrada es `func main()`. `fmt.Println` requiere importar el package `fmt`.

---

## Ejercicio 3 — Corrección de declaraciones

**¿De qué concepto sale?** → Sintaxis de declaración de variables y constantes en Go

**Código original con errores:**

```go
func main() {
    var zz int = 0A        // ❌ 0A no es un literal entero válido en Go
    var z int := x         // ❌ := no se combina con var; además x no está declarada
    x := 10;               // ❌ x ya fue referenciada antes de declararla
    var y int8 = x + 1     // ❌ x es int, y es int8: tipos distintos, no se pueden sumar directamente
    const n := 5001        // ❌ las constantes no usan :=
    const c int := 5001    // ❌ ídem
    /* float //             // ❌ comentario mal cerrado (mezcla /* con //)
    var e float32 := 6     // ❌ := no se combina con var
    f float32 = e          // ❌ falta var o :=
}
```

**a) Versión corregida con modificaciones mínimas:**

```go
package main

import "fmt"

func main() {
    var zz int = 0                  // 0A → 0 (literal válido)
    x := 10                         // declarar x primero
    var z int = x                   // var con = (no :=)
    var y int8 = int8(x) + 1        // conversión explícita int → int8
    const n = 5001                  // const sin tipo ni :=
    const c int = 5001              // const con tipo usa =, no :=
    /* float */
    var e float32 = 6               // var con =
    var f float32 = e               // var con =

    // b) mostrar todos los valores
    fmt.Println("zz:", zz)
    fmt.Println("x:", x)
    fmt.Println("z:", z)
    fmt.Println("y:", y)
    fmt.Println("n:", n)
    fmt.Println("c:", c)
    fmt.Println("e:", e)
    fmt.Println("f:", f)
}
```

**Por qué:** Resumen de reglas: `var` usa `=`, `:=` es solo para declaración corta sin `var`, las constantes usan `=` (nunca `:=`), y toda conversión entre tipos distintos es explícita.

---

## Ejercicio 4 — Suma de pares hasta 250

**¿De qué concepto sale?** → Estructuras de control `for`, constantes, E/S numérica

**a) Iterando hacia adelante:**

```go
package main

import "fmt"

const TOPE = 250

func main() {
    suma := 0
    for i := 2; i <= TOPE; i += 2 {
        suma += i
    }
    fmt.Println("Suma:", suma) // 15750
}
```

**b) Iterando hacia atrás (mismo resultado):**

```go
suma := 0
for i := TOPE; i >= 2; i -= 2 {
    suma += i
}
fmt.Println("Suma:", suma) // 15750
```

**Por qué:** El resultado es el mismo porque la suma es conmutativa. La constante `TOPE` reemplaza el literal `250` — si cambia el rango, se modifica en un solo lugar.

> **gofmt**: herramienta oficial de Go que formatea el código automáticamente según el estilo del lenguaje. Se corre con `gofmt -w hola.go` (reescribe el archivo) o `gofmt hola.go` (muestra el diff).

---

## Ejercicio 5 — Función por tramos

**¿De qué concepto sale?** → `if/else if`, `switch`, operaciones aritméticas, E/S

La función es:
- `|x|` si x ∈ (-∞, -18)
- `x mod 4` si x ∈ [-18, -1]
- `x²` si x ∈ [1, 20)
- `-x` si x ∈ [20, +∞)

**Resolución con `if/else if`:**

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    var x int
    fmt.Scan(&x)

    var resultado int
    if x < -18 {
        resultado = int(math.Abs(float64(x)))
    } else if x >= -18 && x <= -1 {
        resultado = x % 4
    } else if x >= 1 && x < 20 {
        resultado = x * x
    } else if x >= 20 {
        resultado = -x
    }
    // x == 0 no cae en ningún caso

    fmt.Println("f(x) =", resultado)
}
```

**a) El caso x = 0:** El 0 no pertenece a ninguno de los intervalos definidos (`(-∞,-18)`, `[-18,-1]`, `[1,20)`, `[20,+∞)`). La función no está definida en 0 → no se necesita `default`/`else` porque no hay un caso que cubrir. Sí conviene agregar un mensaje de error o manejo explícito.

**Re-escrito con `switch` sin selector:**

```go
switch {
case x < -18:
    resultado = int(math.Abs(float64(x)))
case x >= -18 && x <= -1:
    resultado = x % 4
case x >= 1 && x < 20:
    resultado = x * x
case x >= 20:
    resultado = -x
}
```

**b) Con punto flotante:**

```go
var x float64
fmt.Scan(&x)

var resultado float64
switch {
case x < -18:
    resultado = math.Abs(x)
case x >= -18 && x <= -1:
    resultado = math.Mod(x, 4)
case x >= 1 && x < 20:
    resultado = math.Pow(x, 2)
case x >= 20:
    resultado = -x
}
fmt.Printf("f(%.2f) = %.2f\n", x, resultado)
```

**Por qué:** Con floats, `%` no funciona — hay que usar `math.Mod`. El valor absoluto pasa de conversión manual a `math.Abs`. La potencia usa `math.Pow`.

---

## Ejercicio 6 — División entre mayor y menor

**¿De qué concepto sale?** → E/S, comparación, conversión de tipos, división por cero

**Con enteros:**

```go
package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    mayor, menor := a, b
    if b > a {
        mayor, menor = b, a
    }

    if menor == 0 {
        fmt.Println("Error: división por cero")
    } else {
        fmt.Println("Resultado:", mayor/menor)
    }
}
```

**Con enteros sin signo (`uint`):**

```go
var a, b uint
fmt.Scan(&a, &b)
// igual lógica, pero no hay negativos — mayor/menor siempre >= 0
```

> ⚠️ Con `uint`, si el usuario ingresa un número negativo el comportamiento es indefinido (underflow). Go no lo detecta automáticamente.

**Con floats:**

```go
var a, b float64
fmt.Scan(&a, &b)

mayor, menor := a, b
if b > a {
    mayor, menor = b, a
}

resultado := mayor / menor
fmt.Printf("%.4f\n", resultado)
// Con floats, dividir por 0.0 da +Inf (no panic)
```

**Por qué:** En Go, la división entera por cero produce un **panic** en tiempo de ejecución. La división float por cero produce `+Inf` o `-Inf` según el signo — no hace panic. Hay que manejar explícitamente el caso `menor == 0` para enteros.

---

## Ejercicio 7 — Temperaturas de pacientes

**¿De qué concepto sale?** → `for`, `if`, floats, acumuladores, casting

```go
package main

import "fmt"

func main() {
    var alta, normal, baja int
    var maxTemp, minTemp float64
    var primera bool = true

    for i := 0; i < 10; i++ {
        var temp float64
        fmt.Scan(&temp)

        if primera {
            maxTemp, minTemp = temp, temp
            primera = false
        }
        if temp > maxTemp {
            maxTemp = temp
        }
        if temp < minTemp {
            minTemp = temp
        }

        if temp > 37.5 {
            alta++
        } else if temp >= 36 {
            normal++
        } else {
            baja++
        }
    }

    fmt.Printf("Alta:   %.1f%%\n", float64(alta)/10*100)
    fmt.Printf("Normal: %.1f%%\n", float64(normal)/10*100)
    fmt.Printf("Baja:   %.1f%%\n", float64(baja)/10*100)
    fmt.Printf("Promedio (max+min)/2: %d\n", int((maxTemp+minTemp)/2))
}
```

**a) `switch` con tipos reales:** En Go (y en la mayoría de los lenguajes), el `switch` con selector **no funciona con floats** porque la comparación de igualdad exacta con floats es poco confiable. Se usa `switch` sin selector (equivalente a `if/else if`) o directamente `if`.

**b) Conversión float ↔ entero:** En Go, siempre explícita: `int(f)` trunca hacia cero, `float64(i)` convierte entero a float. No hay conversión implícita. En otros lenguajes como Java o C también es explícita con cast `(int)f`; en Python es implícita en muchos contextos.

---

## Ejercicio 8 — Punto cardinal del viento

**¿De qué concepto sale?** → `switch` con `default`, E/S de strings/chars

```go
package main

import "fmt"

func main() {
    var direccion string
    fmt.Scan(&direccion)

    switch direccion {
    case "N", "n":
        fmt.Println("El viento se dirige hacia el Sur")
    case "S", "s":
        fmt.Println("El viento se dirige hacia el Norte")
    case "E", "e":
        fmt.Println("El viento se dirige hacia el Oeste")
    case "O", "o":
        fmt.Println("El viento se dirige hacia el Este")
    default:
        fmt.Println("Dirección no válida")
    }
}
```

**Por qué:** El `default` captura cualquier entrada que no matchee ningún `case` — equivale al `else` final. En otros lenguajes: `default` en Java/C/C#, `else` en Python (en `match`), `otherwise` en algunos lenguajes funcionales.

---

## Ejercicicio Obligatorio 1 — Reemplazar "jueves" por "martes"

**¿De qué concepto sale?** → Manipulación de strings, package `strings`, respeto de casing

```go
package main

import (
    "fmt"
    "strings"
)

func reemplazarConCasing(frase, original, reemplazo string) string {
    originalLower := strings.ToLower(original)
    var resultado strings.Builder
    fraseLower := strings.ToLower(frase)
    i := 0

    for i < len(frase) {
        if strings.HasPrefix(fraseLower[i:], originalLower) {
            // aplicar casing de la ocurrencia original al reemplazo
            for j, c := range reemplazo {
                if j < len(original) {
                    orig := rune(frase[i+j])
                    if orig >= 'A' && orig <= 'Z' {
                        resultado.WriteRune(c - 32) // mayúscula
                    } else {
                        resultado.WriteRune(c)
                    }
                } else {
                    resultado.WriteRune(c)
                }
            }
            i += len(original)
        } else {
            resultado.WriteByte(frase[i])
            i++
        }
    }
    return resultado.String()
}

func main() {
    var frase string
    fmt.Println("Ingresá la frase:")
    fmt.Scanln(&frase)

    resultado := reemplazarConCasing(frase, "jueves", "martes")
    fmt.Println(resultado)
}
```

**Por qué:** La clave es recorrer el string buscando la palabra en minúsculas (para detectarla sin importar el casing), y luego copiar el casing posición a posición al reemplazo. `strings.Builder` es eficiente para construir strings carácter por carácter.

---

## Ejercicio Obligatorio 2 — Reemplazar "miércoles" por "automóvil"

**¿De qué concepto sale?** → Mismo patrón que Obligatorio 1, pero con strings Unicode (tildes)

**Resolución:** Mismo código que el ejercicio anterior, cambiando los argumentos:

```go
resultado := reemplazarConCasing(frase, "miércoles", "automóvil")
```

**Impacto de usar palabras con tildes:** El problema es que en Go los strings son `[]byte` (UTF-8), y las tildes ocupan más de 1 byte. El índice `i` avanza en bytes, no en caracteres. Hay que trabajar con `[]rune` para manejar correctamente Unicode:

```go
fraseRunes := []rune(frase)
originalRunes := []rune(strings.ToLower(original))
reemplazoRunes := []rune(reemplazo)
// iterar sobre fraseRunes con índice de runes
```

**Por qué:** Con ASCII puro (Ej. 1) `len(string)` == cantidad de caracteres. Con Unicode (Ej. 2) `len(string)` == cantidad de bytes, que puede ser mayor. Usar `[]rune` convierte el string a slice de puntos de código Unicode, donde cada elemento es un carácter real.

---

## Ejercicio Obligatorio 3 — Invertir mayúsculas/minúsculas de ocurrencias

**¿De qué concepto sale?** → Búsqueda case-insensitive + inversión de casing por carácter

```go
package main

import (
    "fmt"
    "strings"
    "unicode"
)

func invertirCasing(s string) string {
    var resultado strings.Builder
    for _, c := range s {
        if unicode.IsUpper(c) {
            resultado.WriteRune(unicode.ToLower(c))
        } else {
            resultado.WriteRune(unicode.ToUpper(c))
        }
    }
    return resultado.String()
}

func main() {
    var palabra string
    fmt.Println("Ingresá la palabra a buscar:")
    fmt.Scan(&palabra)
    fmt.Println("Ingresá la frase:")
    var frase string
    fmt.Scan(&frase)

    palabraLower := strings.ToLower(palabra)
    palabraRunes := []rune(palabraLower)
    fraseRunes := []rune(frase)
    var resultado strings.Builder
    i := 0

    for i < len(fraseRunes) {
        // verificar si hay match case-insensitive en posición i
        fin := i + len(palabraRunes)
        if fin <= len(fraseRunes) {
            segmento := strings.ToLower(string(fraseRunes[i:fin]))
            if segmento == palabraLower {
                resultado.WriteString(invertirCasing(string(fraseRunes[i:fin])))
                i = fin
                continue
            }
        }
        resultado.WriteRune(fraseRunes[i])
        i++
    }

    fmt.Println(resultado.String())
}
```

**Por qué:** La búsqueda se hace en minúsculas para ser case-insensitive. Cuando hay match, se toma el segmento **original** de la frase (con su casing real) y se invierte carácter por carácter con `unicode.IsUpper`/`ToLower`/`ToUpper`. Trabajar con `[]rune` es necesario para soportar caracteres como `Ñ`, `É`, etc.

---

## 🔑 Patrones Go que se repiten en esta práctica

| Situación | Patrón |
|---|---|
| Convertir entre tipos | `T(valor)` — siempre explícito |
| División segura (enteros) | Verificar `denominador != 0` antes |
| Floats y división por cero | Produce `+Inf`, no panic |
| Switch con rangos | `switch` sin selector + condiciones en cada `case` |
| Strings con tildes/Unicode | Convertir a `[]rune` antes de indexar |
| Invertir casing | `unicode.IsUpper` + `ToUpper`/`ToLower` |
