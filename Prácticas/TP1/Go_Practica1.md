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

## Ejercicio 9 — Reemplazar "jueves" por "martes"

**¿De qué concepto sale?** → Manipulación de strings, packages `strings` y `unicode`, respeto de casing

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// reemplazarConCasing busca todas las ocurrencias de `original` dentro de
// `frase` de forma case-insensitive y las reemplaza por `reemplazo`,
// copiando el patrón de mayúsculas/minúsculas de la ocurrencia encontrada
// posición a posición.
//
// Se trabaja con []rune (y no con bytes) porque en Go un string es una
// secuencia de bytes UTF-8; caracteres como é, ñ, ó ocupan 2 bytes, por lo
// que iterar con índice de bytes daría posiciones incorrectas.
// Al convertir a []rune cada elemento representa un carácter completo.
//
// En este ejercicio ambas palabras son ASCII puro ("jueves"/"martes"),
// así que rune y byte coinciden; pero se usa []rune desde ya para que la
// función sea reutilizable con palabras acentuadas (ver Obligatorio 1).
func reemplazarConCasing(frase, original, reemplazo string) string {
	fraseRunes := []rune(frase)
	origRunes := []rune(strings.ToLower(original))
	reemplRunes := []rune(strings.ToLower(reemplazo))
	n := len(origRunes)

	var sb strings.Builder // Builder acumula el resultado sin alocar un string nuevo en cada paso

	i := 0
	for i < len(fraseRunes) {
		// ¿Hay suficientes caracteres por delante para que quepa `original`?
		if i+n <= len(fraseRunes) {
			// Comparamos el segmento en minúsculas con la palabra buscada.
			segmento := strings.ToLower(string(fraseRunes[i : i+n]))
			if segmento == string(origRunes) {
				// Ocurrencia encontrada: escribir `reemplazo` con el casing
				// de cada posición de la ocurrencia original.
				for j, r := range reemplRunes {
					if unicode.IsUpper(fraseRunes[i+j]) {
						sb.WriteRune(unicode.ToUpper(r))
					} else {
						sb.WriteRune(r)
					}
				}
				i += n
				continue
			}
		}
		// Carácter normal: copiarlo tal cual.
		sb.WriteRune(fraseRunes[i])
		i++
	}
	return sb.String()
}

func main() {
	// bufio.Scanner lee la línea completa (incluyendo espacios).
	// fmt.Scan solo lee hasta el primer espacio, lo que partiría la frase.
	fmt.Print("Ingrese una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println(reemplazarConCasing(frase, "jueves", "martes"))
}
```

**Puntos clave:**
- `[]rune`: convierte el string a un slice de puntos de código Unicode. Cada elemento es un carácter completo, sin importar cuántos bytes ocupe. Necesario para no corromper caracteres acentuados al indexar.
- `strings.ToLower`: normaliza a minúsculas solo para comparar; el string original se preserva intacto para copiar el casing.
- `strings.Builder`: acumula el string de salida de forma eficiente, sin crear un string nuevo por cada carácter concatenado.
- `unicode.IsUpper` / `unicode.ToUpper` / `unicode.ToLower`: funciones del package `unicode` que funcionan con cualquier carácter del estándar Unicode, no solo el alfabeto ASCII.

---

## Ejercicio Obligatorio 1 — Reemplazar "miércoles" por "automóvil"

**¿De qué concepto sale?** → Mismo patrón que el Ejercicio 9 aplicado a palabras con caracteres Unicode (tildes)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// reemplazarConCasing es la misma función del Ejercicio 9, reutilizable
// para cualquier par de palabras.
//
// Impacto de pasar de "jueves"/"martes" a "miércoles"/"automóvil":
//   - Ambas palabras tienen 9 runes ("miércoles": m-i-é-r-c-o-l-e-s,
//     "automóvil": a-u-t-o-m-ó-v-i-l), así que el reemplazo posición
//     a posición sigue siendo válido.
//   - La diferencia es que ahora los strings tienen bytes extra por las
//     tildes (é = 2 bytes, ó = 2 bytes). Si usáramos índice de bytes,
//     len("miércoles") == 11 (bytes) pero tiene 9 caracteres reales.
//     Por eso es imprescindible trabajar con []rune.
func reemplazarConCasing(frase, original, reemplazo string) string {
	fraseRunes := []rune(frase)
	origRunes := []rune(strings.ToLower(original))
	reemplRunes := []rune(strings.ToLower(reemplazo))
	n := len(origRunes)

	var sb strings.Builder

	i := 0
	for i < len(fraseRunes) {
		if i+n <= len(fraseRunes) {
			segmento := strings.ToLower(string(fraseRunes[i : i+n]))
			if segmento == string(origRunes) {
				for j, r := range reemplRunes {
					if unicode.IsUpper(fraseRunes[i+j]) {
						sb.WriteRune(unicode.ToUpper(r))
					} else {
						sb.WriteRune(r)
					}
				}
				i += n
				continue
			}
		}
		sb.WriteRune(fraseRunes[i])
		i++
	}
	return sb.String()
}

func main() {
	fmt.Print("Ingrese una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println(reemplazarConCasing(frase, "miércoles", "automóvil"))
}
```

**Impacto de usar palabras con tildes:**

| | "jueves" / "martes" | "miércoles" / "automóvil" |
|---|---|---|
| `len(palabra)` en bytes | 6 / 6 | 11 / 10 |
| `len([]rune(palabra))` | 6 / 6 | 9 / 9 |
| ¿Funciona con índice de bytes? | Sí (todo ASCII) | No (é, ó ocupan 2 bytes) |
| ¿Funciona con `[]rune`? | Sí | Sí |

La función no necesitó modificarse: al estar escrita con `[]rune` desde el principio, absorbe el cambio sin errores.

---

## Ejercicio Obligatorio 2 — Invertir palabras en posiciones impares

**¿De qué concepto sale?** → `strings.Fields`, `strings.Join`, `[]rune`, algoritmo de inversión in-place

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// invertirPalabra devuelve la palabra con sus caracteres en orden inverso.
// Usa []rune para que caracteres Unicode multi-byte (é, ñ, etc.) se
// inviertan como una unidad y no byte a byte (lo que produciría UTF-8 inválido).
func invertirPalabra(s string) string {
	runes := []rune(s)
	// Intercambio desde los extremos hacia el centro (algoritmo in-place).
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Print("Ingrese una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	// strings.Fields divide la frase por cualquier secuencia de espacios en
	// blanco y devuelve un slice de strings con cada palabra.
	palabras := strings.Fields(frase)

	for i := range palabras {
		// Las posiciones se cuentan desde 1, por eso chequeamos (i+1)%2 == 1.
		// Posiciones impares: 1, 3, 5, … → índices 0, 2, 4, …
		if (i+1)%2 == 1 {
			palabras[i] = invertirPalabra(palabras[i])
		}
	}

	// strings.Join vuelve a unir el slice con un espacio entre cada elemento.
	fmt.Println(strings.Join(palabras, " "))
}
```

**Puntos clave:**
- `strings.Fields(s)`: divide `s` por espacios (uno o varios) y devuelve `[]string`. Equivalente a `split` en otros lenguajes.
- `strings.Join(slice, sep)`: une un `[]string` insertando `sep` entre cada elemento.
- Algoritmo de inversión in-place: dos punteros `i` (inicio) y `j` (fin) se acercan intercambiando posiciones hasta cruzarse. No necesita array auxiliar.
- La puntuación pegada a la palabra (como el `.` en `hoy.`) se invierte junto con ella: `hoy.` → `.yoh`. Es el comportamiento esperado para una inversión de caracteres estricta.

---

## Ejercicio Obligatorio 3 — Invertir mayúsculas/minúsculas de ocurrencias

**¿De qué concepto sale?** → Búsqueda case-insensitive con `[]rune`, inversión de casing con package `unicode`

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// invertirCasing recibe un slice de runes y devuelve un string donde cada
// mayúscula se convirtió en minúscula y cada minúscula en mayúscula.
// Usa el package unicode para que funcione con Ñ, É, Ó, etc.
func invertirCasing(runes []rune) string {
	var sb strings.Builder
	for _, c := range runes {
		if unicode.IsUpper(c) {
			sb.WriteRune(unicode.ToLower(c))
		} else {
			sb.WriteRune(unicode.ToUpper(c))
		}
	}
	return sb.String()
}

func main() {
	// Leemos la palabra a buscar con fmt.Scan (lee hasta el primer espacio).
	var palabra string
	fmt.Print("Ingrese la palabra a buscar: ")
	fmt.Scan(&palabra)

	// Leemos la frase completa con bufio.Scanner para capturar los espacios.
	fmt.Print("Ingrese una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	// Convertimos ambos a []rune para trabajar carácter a carácter con Unicode.
	fraseRunes := []rune(frase)
	// Guardamos la palabra en minúsculas para comparar sin importar el casing.
	palabraLower := strings.ToLower(palabra)
	palabraRunes := []rune(palabraLower)
	n := len(palabraRunes)

	var sb strings.Builder
	i := 0

	for i < len(fraseRunes) {
		// Si quedan al menos n caracteres, comparamos el segmento con la palabra.
		if i+n <= len(fraseRunes) {
			segmento := strings.ToLower(string(fraseRunes[i : i+n]))
			if segmento == palabraLower {
				// Ocurrencia encontrada: escribimos el segmento con casing invertido.
				sb.WriteString(invertirCasing(fraseRunes[i : i+n]))
				i += n
				continue
			}
		}
		// No hay ocurrencia en esta posición: copiamos el carácter tal cual.
		sb.WriteRune(fraseRunes[i])
		i++
	}

	fmt.Println(sb.String())
}
```

**Puntos clave:**
- La estructura del `for` es idéntica al Ejercicio 9: avanzar por `fraseRunes`, comparar segmento en minúsculas, actuar si hay match. La única diferencia es qué se hace cuando hay match: en ej9 se copia el casing, acá se invierte.
- `invertirCasing`: recorre el segmento original (con su casing real) y para cada carácter aplica la lógica opuesta con `unicode.IsUpper`.
- La búsqueda es de subcadena, no de palabra completa: encuentra `peqUEño` dentro de `peqUEño,` aunque tenga la coma — invierte solo los caracteres de la palabra y deja la coma intacta.

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
