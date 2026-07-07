# Apunte de repaso — Seminario de Lenguajes opción Go

Apunte armado a partir de las diapositivas de teoría (Go-1 a Go-8 + Go-Package) y del repaso de TP3, para el escrito teórico-práctico y el oral.

## Índice

1. Go-1 — Básicos: variables, tipos, operadores
2. Go-2 — Control de flujo, funciones, package fmt
3. Go-3 — Arrays, Slices, Maps
4. Go-4 — Punteros, Tipos, Structs, Interfaces
5. Go-5 — Errores, funciones como valores, panic/recover
6. Go-6 — Genéricos
7. Go-7 — Concurrencia
8. Go-8 — Problemas clásicos de concurrencia
9. Go-Package — Packages, Modules, Dependencies
10. Preguntas frecuentes de examen (todos los temas)

---

# Go 1 — Programa, Package, Variables, Tipos, Operadores, Asignación

## Qué es Go

Go es un lenguaje de programación de Google (2007), de código abierto, que anda en cualquier sistema operativo. Es **compilado** (se traduce a binario antes de correr) y **fuertemente tipado** (cada variable tiene un tipo fijo, no cambia solo). La sintaxis se parece a C/C++ (como Java, PHP, Python, C#).

**Para qué se usa:** desarrollo web (backend), aplicaciones en red, apps multiplataforma, apps "cloud native" y programas concurrentes (que hacen varias cosas "al mismo tiempo").

**Por qué se usa:** es fácil de aprender, compila y corre rápido, soporta concurrencia y genéricos, tiene manejo automático de memoria (no hace falta liberarla a mano) y corre en Windows/Mac/Linux.

## Estructura de un programa

Todo programa Go tiene esta forma básica:

```go
package main

import ("fmt")

func main() {
    fmt.Println("Hello World!")
}
```

- `package main`: declara a qué paquete pertenece el archivo. Un programa ejecutable siempre es del package `main`.
- `import ("fmt")`: importa las librerías que se van a usar. Se puede importar más de una:

```go
import (
    "fmt"
    "math/rand"
)
```

- Go ignora las líneas en blanco.
- `func main() { ... }` es una función. Todo lo que está dentro de sus `{ }` se ejecuta.
- `fmt.Println(...)` es una función del package `fmt` que sirve para imprimir texto.

**Ojo:** todo ejecutable tiene que pertenecer al package `"main"`.

## Consideraciones generales

- Cada sentencia (instrucción) se separa con un salto de línea (Enter) o con `;`. Se puede escribir todo en una sola línea usando `;`, aunque no es lo normal:

```go
package main; import ("fmt"); func main() { fmt.Println("Hello World!");}
```

- **Ojo:** la llave `{` que abre un bloque (función o estructura de control) **no puede** ir sola al principio de la línea siguiente. Tiene que ir en la misma línea que el `func` o el `if`/`for`/etc.
- Comentarios:

```go
// Comentario de una línea

/* Comentario
   de varias líneas */
```

## Exportación de identificadores (mayúscula vs minúscula)

Si un nombre (función, variable, constante) declarado en un package empieza con **mayúscula**, queda "exportado", es decir, visible desde otros packages. Si empieza con minúscula, es privado de ese package.

```go
package main
import (
    "fmt"
    "math"
)
func main() {
    fmt.Println(math.Pi)
}
```

Acá `Println` es una función exportada del package `fmt`, y `Pi` es una constante exportada del package `math` (por eso ambas empiezan con mayúscula).

## Variables

Se declaran con la palabra clave `var`:

```go
var i int
var c, d, e bool

var (
    i int
    c, d, e bool
)
```

Se pueden declarar e inicializar juntas:

```go
var i int = 1
var i, j int = 1, 2
```

Si hay valores iniciales, Go puede "inferir" (adivinar) el tipo solo:

```go
var c, d, e = true, false, "Texto"
```

También se puede usar `:=` para que Go infiera el tipo (se llama "short assignment" o asignación corta):

```go
k := 3
```

**Ojo — trampa típica:** `:=` **solo se puede usar dentro de una función**. Adentro de un package (fuera de cualquier función) toda sentencia tiene que empezar con una palabra clave como `var` o `func`, así que ahí no se puede usar `:=`.

```go
package main
import "fmt"
var b, c, d bool     // esto está bien, es a nivel package
func main() {
    var i int        // acá adentro sí se podría usar i := 0
    fmt.Println(i, b, c, d)
}
```

Nombres de variables:
- Son *case-sensitive* (mayúscula y minúscula son distintas).
- Pueden tener letras, números y `_`, pero no pueden **empezar** con `_`.
- No pueden ser igual a una palabra reservada del lenguaje (como `var`, `func`, etc).

## Constantes

Se declaran igual que las variables pero con `const` en vez de `var`. Pueden ser character, string, boolean o numéricas.

```go
const Pi = 3.14
const A = 1

const (
    A int = 1
    B     = 3.14
)
```

**Ojo:** las constantes **no** se pueden declarar con `:=`.

## Tipos básicos

- `bool`: `true` / `false`. Valor por defecto: `false`.
- `string`: valor por defecto `""` (string vacío).
- `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`: valor por defecto `0`. **Conviene usar siempre `int`** salvo que haya una razón concreta para usar otro.
- `byte` es un alias de `uint8`.
- `rune` es un alias de `int32` (se usa para representar un carácter Unicode).
- `float32`, `float64`: valor por defecto `0`.
- `complex64`, `complex128`: valor por defecto `(0+0i)`.

## Conversión de tipos

Si `T` es un tipo y `v` es un valor, `T(v)` convierte el valor de `v` al tipo `T`. Se puede hacer entre tipos numéricos, o entre tipos que tengan el mismo "tipo subyacente" (la base sobre la que están construidos).

```go
i := 42
f := float64(i) + 0.000001   // 42.000001
u := uint(f)                  // 42
g := float64(int(f))          // 42
fmt.Println(i, f, u, g)
```

## Named types (tipos con nombre propio)

Se puede crear un tipo nuevo a partir de uno existente, para que el compilador evite mezclas sin querer entre cosas que en el fondo son el mismo tipo numérico pero significan cosas distintas:

```go
type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}
```

Estos tipos tienen la misma estructura y operaciones que su tipo subyacente (`float64`), y se pueden convertir entre sí porque comparten ese tipo subyacente.

## Operadores

**Aritméticos** (mismo tipo en ambos operandos):

| Operador | Descripción | Ejemplo |
|---|---|---|
| `+` | Suma | `x + y` |
| `-` | Resta | `x - y` |
| `*` | Multiplicación | `x * y` |
| `/` | División | `x / y` |
| `%` | Módulo (resto) | `x % y` |
| `-` | Menos unario | `-x` |
| `+` | Más unario | `+x` |

**Sobre strings:**

| Operador | Descripción | Ejemplo |
|---|---|---|
| `+` | Concatenación | `s = s1 + s2` |

**De comparación** (mismo tipo en ambos operandos):

| Operador | Descripción | Ejemplo |
|---|---|---|
| `==` | Igualdad | `x == y` |
| `!=` | Desigualdad | `x != y` |
| `>` | Mayor que | `x > y` |
| `<` | Menor que | `x < y` |
| `>=` | Mayor o igual | `x >= y` |
| `<=` | Menor o igual | `x <= y` |

**Lógicos** (sobre booleanos):

| Operador | Descripción | Ejemplo |
|---|---|---|
| `&&` | And | `x < 5 && x < 10` |
| `\|\|` | Or | `x < 5 \|\| x > 10` |
| `!` | Not | `!(x < 5 && x < 10)` |

La diapositiva pregunta "Long circuit ó short circuit ???" — es para pensar si `&&` y `||` evalúan siempre los dos lados o cortan apenas saben el resultado (esto es un punto para repasar/preguntar).

**Bitwise (bit a bit)**, sobre tipos numéricos:

| Operador | Descripción | Ejemplo |
|---|---|---|
| `&` | And | `x & y` |
| `\|` | Or | `x \| y` |
| `^` | Xor | `x ^ b` |
| `<<` | Shift a izquierda | `x << 2` |
| `>>` | Shift a derecha | `x >> 2` |

## Asignación

La asignación **no es un operador**, es una instrucción del lenguaje. Sirve para darle valor a variables y constantes (en su declaración).

```go
variable = expresión   // la expresión debe ser del mismo tipo que la variable
x = x + y * 10
```

Existen versiones "compuestas" de la asignación, que combinan la operación con la asignación:

| Instrucción | Ejemplo | Equivale a |
|---|---|---|
| `+=` | `x += 3` | `x = x + 3` |
| `-=` | `x -= 3` | `x = x - 3` |
| `*=` | `x *= 3` | `x = x * 3` |
| `/=` | `x /= 3` | `x = x / 3` |
| `%=` | `x %= 3` | `x = x % 3` |
| `&=` | `x &= 3` | `x = x & 3` |
| `\|=` | `x \|= 3` | `x = x \| 3` |
| `^=` | `x ^= 3` | `x = x ^ 3` |
| `>>=` | `x >>= 3` | `x = x >> 3` |
| `<<=` | `x <<= 3` | `x = x << 3` |

Y los operadores de incremento y decremento:

| Operador | Descripción | Ejemplo |
|---|---|---|
| `++` | Incremento | `x++` |
| `--` | Decremento | `x--` |

**Ojo:** el título de esta diapositiva dice "Operadores (Asignación??)" con signos de pregunta — el profesor está marcando que `++` y `--` tampoco son operadores clásicos sino instrucciones, en la misma línea que lo de la asignación.

## Posibles huecos / revisar

- Página 16 (tabla de operadores aritméticos): el pie de página tapa un poco la última fila de la tabla ("Más unario"), pero el contenido se entiende igual.
- Ninguna otra página presentó problemas de lectura.

---

# Go 2 — Estructuras de control, Funciones, Package "fmt"

## Estructuras de control: Secuencia

Es simplemente ejecutar instrucciones una después de la otra, en orden.

```go
x := 5
fmt.Println(x)
x++
fmt.Println(x)

// o todo junto separado por ";"
x := 5; fmt.Println(x); x++; fmt.Println(x)
```

## Estructuras de control: Iteración (for)

Go **solo tiene `for`** (no hay `while` aparte, se hace todo con `for`).

Formas del `for`:

```go
// for "infinito", hay que cortarlo con break adentro
for {
    // secuencia
}

// for tipo "while": solo la condición
sum := 1
for sum < 1000 {
    sum += sum
}

// for clásico: inicialización; condición; incremento
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}

// se puede omitir la inicialización
sum := 1
for ; sum < 1000; {
    sum += sum
}

// se puede omitir el incremento (se hace a mano adentro)
for i := 1; i < 500; {
    i += i
    fmt.Println(i)
}
```

También se pueden inicializar y actualizar **varias variables a la vez**, separadas por coma:

```go
for i, j := 0, 10; i <= j; i, j = i+1, j-1 {
    fmt.Println(n, ":", i, "-", j)
}
```

**Ojo:** el `for` de Go es muy flexible — cualquiera de las tres partes (inicialización, condición, incremento) puede faltar, pero los `;` que las separan solo se sacan si es el `for` tipo "while" (una sola condición) o el `for` infinito (sin nada).

### Iteración tipo repeat / do-while

Go no tiene `do-while`, pero se simula con un `for` infinito más un `if` con `break` adentro (o sea, primero se ejecuta el cuerpo y recién después se chequea si hay que cortar):

```go
i := 0
for {
    i++
    if i >= 10 {
        break
    }
}
```

## Estructuras de control: Selección "if"

```go
if x > y {
    fmt.Println(x)
    fmt.Println(y)
}

if x < y {
    fmt.Println(x)
} else {
    fmt.Println(y)
}

if x > y && x > z {
    fmt.Println("x")
} else if y > x && y > z {
    fmt.Println("y")
} else {
    fmt.Println("z")
}
```

### If con sentencia de inicialización

Se puede declarar una variable justo antes de la condición, separada por `;`. Esa variable solo existe dentro del `if` (y de sus `else`):

```go
x := 3.0
n := 2.0
lim := 10.0

if v := math.Pow(x, n); v < lim {
    fmt.Println(v)
} else {
    fmt.Println(lim)
}
// fmt.Println(v) acá afuera daría error, v ya no existe
```

**Ojo:** esta es una trampa típica de examen — la variable declarada en el `if` (acá `v`) **no existe fuera** del bloque `if/else`.

## Estructuras de control: Selección "switch"

```go
switch runtime.GOOS {
case "darwin" + "win":
    fmt.Println("OS X.")
case "linux" + "nux":
    fmt.Println("Linux.")
default:
    fmt.Println("Other")
}
```

(Nota: el `case` puede tener una expresión, no solo un valor fijo — en el ejemplo se concatenan strings, aunque en la práctica lo normal es comparar contra valores concretos como `"darwin"` o `"linux"`, como en el ejemplo con inicialización de abajo.)

### Switch con sentencia de inicialización

Igual que el `if`, se puede declarar una variable antes del selector:

```go
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    fmt.Println("Other")
}
```

### Switch sin selector

Si no se pone nada después de `switch`, es como poner `switch true` — cada `case` es una condición booleana y se ejecuta la primera que sea verdadera:

```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

**Ojo:** en Go, a diferencia de C/Java, el `switch` **no** hace *fallthrough* automático entre `case` (no hace falta poner `break`, cada `case` corta solo al terminar).

## Funciones

Se declaran con `func`:

```go
func nombreFuncion() {
    fmt.Println("Esta es una función")
}
nombreFuncion() // así se llama
```

Con parámetros:

```go
func add(x int, y int) {
    fmt.Println(x + y)
}
add(2, 3)
```

**Ojo (ahorro de tipeo):** si dos o más parámetros seguidos tienen el mismo tipo, se puede escribir el tipo una sola vez al final:

```go
func add(x, y int) {
    fmt.Println(x + y)
}
add(2, 3) // funciona igual
```

### Funciones con valor de retorno

```go
func add(x, y int) int {
    return x + y
}
z := add(2, 3)
```

Una función puede devolver **más de un valor**:

```go
func swap(x int, y int) (int, int) {
    return y, x
}
a, b = swap(a, b)
```

También se pueden nombrar los valores de retorno (named return values). En ese caso, `return` sin nada devuelve esos valores tal cual quedaron:

```go
func swap(x1 int, y1 int) (x2, y2 int) {
    x2, y2 = y1, x1
    return
}
a, b = swap(a, b)
```

## Package "fmt"

Ejemplo base de siempre:

```go
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}
```

### Print, Println, Printf

```go
const name, age = "Kim", 22

fmt.Print(name, " is ", age, " years old.\n")
// Kim is 22 years old.
// Print pone espacio entre argumentos, EXCEPTO entre strings

fmt.Println(name, "is", age, "years old.")
// Kim is 22 years old.
// Println pone espacio entre TODOS los argumentos (incluso strings) y agrega un salto de línea al final

fmt.Printf("%s is %d years old.\n", name, age)
// Kim is 22 years old.
// Printf usa "marcas" o "verbos" (%s, %d, etc.) dentro de un string de formato
```

**Ojo:** la diferencia entre `Print` y `Println` es sutil — `Print` solo agrega espacio entre dos argumentos si **ninguno de los dos** es un string; `Println` siempre pone espacio entre todos y además agrega un salto de línea al final automáticamente (`Print` no).

### Marcas / verbos de Printf — generales

Con `i := 42`, `s := "Pepe"`, `b := true`:

| Verbo | Qué hace | Ejemplo resultado |
|---|---|---|
| `%v` | Formato predeterminado | `42 Pepe` |
| `%#v` | Valor representado en sintaxis Go | `42 "Pepe"` |
| `%T` | Tipo representado en sintaxis Go | `int string` |
| `%%` | Un `%` literal | `Pepe 42 %` |
| `\n` / `\r\n` | Salto de línea | — |
| `\t` | Tabulación | — |

### Marcas para enteros

Con `i := 128578`:

| Verbo | Qué hace | Resultado |
|---|---|---|
| `%d` | Decimal | `128578` |
| `%b` | Binario | `11111011001000010` |
| `%x` | Hexadecimal minúscula | `1f642` |
| `%X` | Hexadecimal mayúscula | `1F642` |
| `%o` | Octal | `373102` |
| `%O` | Octal con prefijo `0o` | `0o373102` |
| `%c` | Carácter Unicode | el emoji correspondiente |
| `%q` | Carácter Unicode entre comillas simples | `'🙂'` |
| `%U` | Formato Unicode | `U+1F642` |

### Marcas para strings

Con `i := "Pepe"`:

| Verbo | Qué hace | Resultado |
|---|---|---|
| `%s` | Valor normal | `Pepe` |
| `%q` | Valor entre comillas | `"Pepe"` |
| `%x` | Base 16 minúscula | `50657065` |
| `%X` | Base 16 mayúscula | `50657065` |

### Marcas para floats

Con `pi := math.Pi`:

| Verbo | Qué hace | Resultado |
|---|---|---|
| `%e` | Notación científica (minúscula) | `3.141593e+00` |
| `%E` | Notación científica (mayúscula) | `3.141593E+00` |
| `%f` | Con decimales, sin exponente | `3.141593` |
| `%F` | Igual que `%f` | `3.141593` |
| `%g` | `%e` si el exponente es grande, si no `%f` | `3.141592653589793` |
| `%G` | Igual que `%g` pero con `%E` | `3.141592653589793` |

### Equivalencias de %v según el tipo

| Tipo | Verbo equivalente |
|---|---|
| `bool` | `%t` |
| `int`, `int8`, etc. | `%d` |
| `uint`, `uint8`, etc. | `%d` |
| `float32`, `complex32`, etc. | `%g` |
| `string` | `%s` |

O sea: `%v` en el fondo usa `%t` para booleanos, `%d` para enteros, `%g` para floats/complex, y `%s` para strings.

### Width y precision (ancho y precisión)

Con `i := 123` y `f := 123.12`, se puede controlar cuántos caracteres ocupa lo impreso y cuántos decimales mostrar, poniendo un número antes (o después de un punto) del verbo:

```go
fmt.Printf("%d\n", i)     // 123
fmt.Printf("%6d\n", i)    //    123   (ancho mínimo 6, rellena con espacios)
fmt.Printf("%06d\n", i)   // 000123   (ancho 6, rellena con ceros)
fmt.Printf("%f\n", f)     // 123.120000
fmt.Printf("%8.2f\n", f)  //   123.12  (ancho 8, 2 decimales)
fmt.Printf("%+d\n", i)    // +123     (el + siempre muestra el signo)
```

**Ojo:** el número antes del punto es el **ancho total** mínimo, y el número después del punto es la **cantidad de decimales** (solo aplica a floats).

### Flags para investigar

La diapositiva deja como tarea de investigación (no da la explicación en clase) el significado de estas marcas dentro de `Printf`:

- `"-"` (ej: `%-d`, `%-6d`, `%-06d`, `%-f`, `%-8.2f`, `%-08.2f`)
- `"#"`
- `" "` (un espacio)
- `"%.2f"`
- `"%9.f"`

**Ojo:** esto quedó como "para investigar" en la propia diapositiva — conviene repasarlo por separado antes del examen, porque no está desarrollado acá.

### Sprintf, Sprint, Sprintln

Son iguales a `Printf`, `Print` y `Println`, pero en vez de imprimir en pantalla devuelven el resultado como un `string`:

```go
const name, age = "Kim", 22

s := fmt.Sprintf("%s is %d years old.\n", name, age)
s := fmt.Sprint(name, " is ", age, " years old.\n")
s := fmt.Sprintln(name, "is", age, "years old.")
// Los tres arman: "Kim is 22 years old."
```

### Scan, Scanf, Scanln (leer datos de entrada)

Sirven para leer valores que ingresa el usuario. Necesitan que les pasemos **la dirección** de la variable (con `&`), porque tienen que escribir el valor adentro:

```go
func Scan(...) (n int, err error)
var mensaje string
n, e := fmt.Scan(&mensaje)

func Scanf(format string, ...) (n int, err error)
var (nom string; ape string; tel int)
n, e := fmt.Scanf("%s %s %d", &nom, &ape, &tel)
if e != nil {
    fmt.Printf("Error: %s", e)
} else {
    fmt.Printf("Todo bien, recibimos %d argumentos: %s, %s, %d", n, nom, ape, tel)
}

func Scanln(...) (n int, err error)
var nom, ape string
n, e := fmt.Scanln(&nom, &ape)
```

Todas devuelven cuántos valores lograron leer (`n`) y un posible error (`e`).

### Sscan, Sscanln, Sscanf (leer datos desde un string)

Son iguales a `Scan`/`Scanln`/`Scanf`, pero en vez de leer de la entrada del usuario, leen desde un `string` que le pasamos nosotros como primer argumento:

```go
func Sscan(str string, ...) (n int, err error)
func Sscanln(str string, ...) (n int, err error)
func Sscanf(str string, format string, ...) (n int, err error)

var x, y string
n, e := fmt.Sscan("100\n200", &x, &y)     // n:2 e:<nil>  x:100 y:200
n, e = fmt.Sscan("300 400", &x, &y)       // n:2 e:<nil>  x:300 y:400
n, e = fmt.Sscanf("500 600", "%s %s", &x, &y) // n:2 e:<nil>  x:500 y:600
n, e = fmt.Sscanln("700\n800", &x, &y)    // n:1 e:"unexpected newline"  x:700 y:600 (y no cambia)
n, e = fmt.Sscanln("900 1000\n", &x, &y)  // n:2 e:<nil>  x:900 y:1000
n, e = fmt.Sscanln("1100 1200", &x, &y)   // n:2 e:<nil>  x:1100 y:1200
```

**Ojo — trampa típica:** `Sscanln` (y `Scanln`) esperan que los valores estén en **una sola línea**. Si aparece un salto de línea (`\n`) antes de haber leído todos los valores esperados, corta ahí con error `"unexpected newline"` y la variable que no llegó a leer **se queda con el valor que tenía antes** (no se pisa). Por eso en el ejemplo `fmt.Sscanln("700\n800", &x, &y)` da error y `y` queda en `600` (el valor de la línea anterior), no en `800`.

## Posibles huecos / revisar

- Página 16 (tabla general de marcas de `Printf`): el pie de página tapa la última fila, que empezaba a mostrar la sección "Boolean" con su verbo (probablemente `%t`). No se ve el contenido completo de esa fila en la imagen, aunque se puede inferir por la tabla de equivalencias de `%v` de la página 20 (bool → `%t`). Conviene confirmar con el material original.
- Página 22 ("Package fmt (flags)"): la diapositiva solo lista flags para investigar por cuenta propia (`-`, `#`, espacio, `%.2f`, `%9.f`), sin explicar qué hace cada uno. No se completó esa parte porque no está desarrollada en la fuente.

---

# Go - Teoría 3: Arrays, Slices, Maps

Apunte simple para repasar antes del examen. Basado en las diapositivas "Go-3".

## Arrays

Un array es una lista de elementos, todos del mismo tipo, con tamaño fijo (no cambia nunca) y el primer índice es el 0.

```go
var x [5]int
x[4] = 100
fmt.Println(x) // [0 0 0 0 100]
```

Hay varias formas de escribir lo mismo (declarar y llenar un array):

```go
// Una por una
var x [5]float64
x[0] = 98
x[1] = 93
// ...

// Con llaves, en el mismo := 
x := [5]float64{98, 93, 77, 82, 83}

// Dejando que Go cuente el tamaño (con "...")
x := [...]float64{98, 93, 77, 82, 83}
```

Para recorrer un array podés usar el `for` de toda la vida, o `for range` (más cómodo):

```go
for i, value := range x {
    total += value
    fmt.Println(i, value)
}

for _, value := range x { // el _ descarta el índice
    total += value
}

for i := range x { // solo el índice, sin valor
    fmt.Println(i)
}
```

**Ojo:** también se puede inicializar solo algunas posiciones indicando el índice:

```go
arr := [5]int{1: 10, 2: 20, 3: 30}
fmt.Println(arr) // [0 10 20 30 0]
```
Las posiciones que no nombrás quedan en su "valor cero" (0 para números).

Para saber el tamaño: `len(arr)`.

### Arrays multidimensionales

Un array puede tener adentro otros arrays (una "matriz"):

```go
a := [2][2]string{
    {"Hello", "World"},
    {"Hola", "Mundo"},
}
fmt.Println(a[0], a[1]) // [Hello World] [Hola Mundo]
fmt.Println(a)          // [[Hello World] [Hola Mundo]]
```

## Slices

El problema de los arrays es que el tamaño es parte del tipo (un `[5]int` y un `[6]int` son tipos distintos), y eso es medio incómodo. Por eso en la práctica se usa mucho más el **slice**.

Un slice es como un "pedazo" (segmento) de un array. Se puede indexar y tiene una longitud, pero esa longitud **sí puede cambiar**.

Dos medidas importantes:
- `len(s)`: cuántos elementos tiene ahora.
- `cap(s)`: cuánto lugar hay disponible en el array de atrás, antes de tener que pedir más memoria.

Formas de crear un slice:

```go
a := [6]int{10, 11, 12, 13, 14, 15}
s := a[2:4]              // len(s)=2, cap(s)=4

var s1 []int              // len=0, cap=0 (nil)
s2 := []int{}              // len=0, cap=0
s3 := []int{1, 2, 3}       // len=3, cap=3

s1 := make([]int, 5, 10)   // len=5, cap=10
s2 := make([]int, 5)       // len=5, cap=5
```

Un slice sin inicializar (`var s []int`) es un slice **nil**:

```go
var s []int
fmt.Println(s, len(s), cap(s)) // [] 0 0
if s == nil {
    fmt.Println("nil!") // nil!
}
```

### Slices son referencias

Esto es lo más importante de todo el tema: un slice **no tiene su propia copia de los datos**, apunta al array de atrás. Si modificás el slice, modificás el array (¡y viceversa!).

```go
a := [6]int{10, 11, 12, 13, 14, 15}
s := a[2:4]
fmt.Println(a, s) // [10 11 12 13 14 15] [12 13]

s[1] = 31
a[2] = 21
fmt.Println(a, s) // [10 11 21 31 14 15] [21 31]
```

**Ojo:** re-cortar un slice (`s[1:4]`, `s[:2]`, etc.) puede cambiar `len` y `cap`, pero sigue siendo el mismo array de atrás:

```go
s := []int{1, 2, 3, 4, 5, 6}
t := s[1:4]
fmt.Println(t, len(t), cap(t)) // [2 3 4] 3 5
t = s[:2]
fmt.Println(t, len(t), cap(t)) // [1 2] 2 6
```

Y si reasignás un slice a otra variable (por ejemplo `a = z`), la variable `a` pasa a apuntar a los datos de `z` (no se mezclan).

### Agregar elementos: `append`

```go
func append(slice []Type, elems ...Type) []Type

slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...) // el "..." expande el otro slice
```

**Ojo:** cuando el slice se queda sin `cap` para crecer, Go crea un array nuevo más grande por atrás (típicamente duplica la capacidad) y copia todo ahí. Por eso a veces `cap` "salta" mucho más de lo que uno esperaría:

```go
var s []int
s = append(s, 0)          // len=1 cap=1  [0]
s = append(s, 1)          // len=2 cap=2  [0 1]
s = append(s, 2, 3, 4)    // len=5 cap=6  [0 1 2 3 4]
```

### Copiar slices: `copy`

```go
func copy(dst, src []Type) int
```

No hace falta que `dst` y `src` tengan el mismo tamaño (copia lo que entre en el más chico).

**Ojo (trampa clásica):** si en vez de `copy` hacés `dst := src[:n]`, seguís compartiendo el array de atrás. Modificar el array original (`numbers[2] = 100`) se refleja en cualquier slice sacado con `[:]`, pero **no** en uno creado con `copy` (porque `copy` sí hace una copia real de los valores):

```go
numbers := []int{1, 2, 3, ..., 20}
neededNumbers := numbers[:5]           // comparte array con numbers
numbersCopy := make([]int, len(neededNumbers))
copy(numbersCopy, neededNumbers)        // copia real, independiente

numbers[2] = 100
// numbers y neededNumbers ahora tienen el 100
// numbersCopy NO cambia, sigue con el valor viejo
```

### Slices multidimensionales

Al igual que los arrays, un slice puede tener adentro otros slices (`[][]string`, por ejemplo). Las diapositivas muestran un ejemplo de tablero de tres-en-raya (ta-te-ti) armado así, pero la idea central es la misma: cada "fila" es un slice independiente.

## Parámetros: por valor vs por copia

Algo para tener muy presente para el examen: **los parámetros de las funciones en Go siempre se pasan por valor** (se copian). Esto es distinto según el tipo:

- Si el parámetro es un **array**, se copia todo el array entero. Modificarlo adentro de la función no afecta al original.
- Si el parámetro es un **slice**, se copia el slice (que por dentro es solo un puntero + len + cap), pero apunta al mismo array de atrás. Por eso modificar los *elementos* de un slice adentro de una función sí afecta al original (aunque hacer `append` puede o no afectarlo, según si hay que crear un array nuevo).

```go
type Array600Int [600]int

func sumPrimes(arr Array600Int) (res int) {
    for _, e := range arr {
        res += e
    }
    arr[0] = 17           // esto NO se ve afuera, porque arr es una copia
    return
}
```

## Maps

Un map es una colección **no ordenada** de pares clave-valor (también llamados diccionarios, tablas hash, o arreglos asociativos).

- No permite claves repetidas.
- El valor por defecto de un map no inicializado es `nil`.
- Se puede agregar, modificar y borrar elementos, **excepto si el map es nil** (ahí explota en tiempo de ejecución).

```go
var a map[string]string   // a es nil
// a["clave"] = "Valor"   // ERROR EN TIEMPO DE EJECUCIÓN !!!

var b = make(map[string]string) // así sí se puede usar
b["clave"] = "Valor"

c := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
```

**Ojo:** `fmt.Println` de un map siempre lo muestra ordenado alfabéticamente por clave, aunque el map en sí no tenga ningún orden interno.

### Qué puede ser clave y qué puede ser valor

- Clave: cualquier tipo que tenga definida la comparación `==` (booleanos, números, strings, arrays...).
- **No puede ser clave:** slices, maps, ni funciones (justamente porque no se pueden comparar con `==`).
- Valor: cualquier tipo, sin restricciones.

### Operaciones básicas

```go
m[key] = value       // agregar o modificar
delete(m, key)       // eliminar

elem := m[key]        // si key no está, elem es el "zero value" del tipo
elem, ok := m[key]    // "ok" te dice si la clave existía o no (comma-ok)
```

### Los maps también son referencias

Igual que los slices: si asignás un map a otra variable, ambas apuntan a los mismos datos.

```go
b := a
b["year"] = "1970"
// tanto a como b ahora muestran year:1970
```

### Recorrer un map con `range`

```go
for k, v := range m {
    fmt.Printf("(%s:%s) ", k, v)
}
```

**Ojo:** el orden en que `range` recorre un map **no está garantizado** (puede cambiar cada vez que se ejecuta el programa), a diferencia de lo que se ve al hacer `fmt.Println(m)` directamente (que sí ordena).

## Posibles huecos / revisar

No se detectaron páginas ilegibles o con contenido cortado en este PDF. La página 13 (slices multidimensionales, ejemplo de tres en raya) tiene bastante código superpuesto/denso, pero se pudo leer completo — si al repasar algo no cierra, vale la pena mirar esa diapositiva directamente.

---

# Go - Teoría 4: Pointers, Tipos, Structs e Interfaces

Apunte simple para repasar antes del examen. Basado en las diapositivas "Go-4".

## Pointers (punteros)

Un puntero es la dirección de memoria de un valor. `*T` significa "puntero a un valor de tipo T". El valor por defecto (zero value) de un puntero es `nil`.

```go
var p *int
// fmt.Println(*p) // ¡ERROR en ejecución! (p es nil, no apunta a nada)

if p != nil {
    fmt.Println(*p)
} else {
    fmt.Println("nil") // esto es lo que se imprime
}

i := 42
p = &i          // & = "dame la dirección de i"
fmt.Println(*p) // 42, * = "dame el valor al que apunta p"
```

**Ojo:** desreferenciar (`*p`) un puntero `nil` es un error en tiempo de ejecución, no de compilación. Siempre conviene chequear `if p != nil` antes.

### El allocador `new(T)`

`new(T)` reserva memoria para un valor de tipo `T` (inicializado en su zero value) y devuelve un puntero a él.

```go
p := new(int)
q := new(int)
*p = 10
*q = 5
fmt.Println(*p, *q) // 10 5

q = p // ahora p y q apuntan a lo mismo
fmt.Println(*p, *q) // 10 10
```

### Garbage collector

Cuando ya nadie tiene un puntero apuntando a cierto valor, ese valor queda "huérfano" y el recolector de basura (garbage collector) lo libera solo. En el ejemplo de arriba, después de `q = p`, el valor `5` que tenía `q` antes queda sin nadie que lo referencie, así que se recolecta.

### Parámetros con punteros

Los parámetros de una función **siempre son por valor**, incluso los punteros. Lo que se copia es la dirección, no lo que está adentro. Por eso, si querés que una función modifique algo "de afuera", tenés que pasarle un puntero:

```go
func zero(xPtr *int) {
    *xPtr = 0
}

func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // 0 (sí cambió, porque le pasamos la dirección)
}
```

## Declaración de Tipos (`type`)

Con `type` podés crear un tipo "nombrado" propio, basado en otro tipo (el "tipo subyacente"):

```go
type MyInt int
type MyFloat float64
type MyString string
type MyArray [5]int
type MySlice []int
```

Un tipo nombrado tiene los mismos valores y operaciones que su tipo subyacente, **pero no es compatible con él ni con otros tipos nombrados del mismo subyacente**, aunque sean "iguales por dentro":

```go
var f float64 = 1.5
var mf MyFloat = 2.5
var mf2 MyFloat2 = 3.5  // otro tipo nombrado distinto, mismo subyacente

f = mf   // ERROR: type mismatch
mf = f   // ERROR: type mismatch
// tampoco se puede comparar f > mf ni f == mf
```

**Ojo:** esto confunde bastante — aunque `MyFloat` "es" un `float64` por dentro, Go los trata como tipos totalmente distintos e incompatibles entre sí.

### Casting (conversión de tipos)

Para pasar de un tipo a otro compatible, se usa la sintaxis `T(x)`:

```go
type MyString *string
type MyString2 *string

ms2 = MyString2(ms) // conversión explícita, ahora sí funciona
```

**Regla:** una conversión `T(x)` es válida solo si el tipo origen y el tipo destino tienen el **mismo tipo subyacente**.

Ejemplo típico (conversión de temperaturas):

```go
type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
```

## Métodos

Go no tiene clases, pero permite definirle métodos a los tipos nombrados. Un método es básicamente una función con un argumento extra especial llamado **receiver**, que va antes del nombre de la función:

```go
func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}
```

Esto es casi lo mismo que una función común (`func CToF(c Celsius) Fahrenheit`), solo que se llama con notación de punto: `c.CToF()` en vez de `CToF(c)`.

### El receiver por valor no modifica el original

Como el receiver actúa igual que un parámetro común, si el receiver es por valor, los cambios que hagas adentro del método **no se ven afuera**:

```go
type MyFloat float64

func (f MyFloat) Scale(s float64) {
    f = f * MyFloat(s) // esto NO cambia el f de afuera
}
```

### Para modificar el receiver: usar puntero

```go
type MyFloat float64

func (f *MyFloat) Scale(s float64) {
    *f = *f * MyFloat(s) // ahora sí modifica el valor original
}
```

**Ojo (importante):** Go te deja llamar un método con receiver puntero tanto desde una variable puntero como desde una variable valor (`f.Scale(2)` funciona aunque `f` no sea un puntero) — Go agrega el `&` automáticamente por vos. Por eso en las diapositivas los tres ejemplos (con puntero explícito, con `:=` y puntero, y directo con valor) terminan dando el mismo resultado.

### El receiver puede ser nil

Si el receiver es un puntero, puede llegar como `nil`, y hay que chequearlo adentro del método si hace falta:

```go
func (mv *MySlice) Add() (res int) {
    if *mv == nil {
        return
    }
    for _, e := range *mv {
        res += e
    }
    return
}
```

## Structs

Un struct agrupa varios campos, que pueden ser de tipos distintos:

```go
type Person struct {
    firstname string
    lastname  string
    age       int
}

var p1 Person                          // zero value: {"" "" 0}
p2 := Person{"Pepe", "Sargento", 25}    // inicialización posicional
p3 := new(Person)                       // p3 es un *Person
p4 := Person{lastname: "Larralde", firstname: "José"} // inicialización nombrada
```

Se accede a los campos con notación de punto, y funciona igual con punteros (Go hace la desreferencia sola):

```go
p1.firstname = "John"
p3.age = 28   // aunque p3 sea un *Person, no hace falta escribir (*p3).age
```

Los structs pueden anidarse (un campo que es otro struct), y se accede encadenando puntos:

```go
type Date struct{ day, month, year int }
type Person struct {
    firstname, lastname string
    birthdate            Date
}

p3.birthdate.day = 28
```

**Ojo:** al igual que los arrays, los structs se pasan **por copia** a las funciones. Modificar un campo adentro de una función no afecta al struct original:

```go
func circleArea(c Circle) float64 {
    c.r++ // esto NO se ve afuera
    return math.Pi * c.r * c.r
}
```

### Comparación e igualdad

Dos structs son comparables con `==` si **todos** sus campos son comparables (y se comparan campo por campo):

```go
type Point struct{ X, Y int }
p := Point{1, 2}
q := Point{2, 1}
fmt.Println(p == q) // false
```

Como consecuencia, un tipo struct comparable se puede usar como **clave de un map**:

```go
type address struct {
    hostname string
    port     int
}
hits := make(map[address]int)
hits[address{"golang.org", 443}]++
```

## Struct embedding (composición)

Se puede evitar repetir campos entre structs parecidos usando un campo **anónimo** (le ponés solo el tipo, sin nombre):

```go
type Point struct{ X, Y int }

type Circle struct {
    Point // campo anónimo: "está embebido"
    Radius int
}

type Cylinder struct {
    Circle // también embebido
    Height int
}
```

Esto permite acceder a los campos del tipo embebido de dos formas: la completa (`cy.Circle.Point.X`) o la "promovida", más corta (`cy.X` directo):

```go
var cy Cylinder
cy.Circle.Point.X = 1  // {{{1 0} 0} 0}
cy.Circle.X = 2        // forma intermedia, también funciona
cy.X = 3               // forma promovida, más corta
```

**Ojo:** esto se parece a la herencia de otros lenguajes, pero **no lo es**. Es composición: el tipo de afuera tiene "adentro" una copia del tipo embebido, y Go simplemente te deja acceder a sus campos sin escribir el camino completo.

## Interfaces

Un tipo `interface` se define por un conjunto de firmas de métodos (nombre + parámetros + retorno, sin implementación). Un valor de tipo interface puede ser **cualquier cosa** que implemente esos métodos.

**Ojo clave:** en Go no existe la palabra `implements`. Un tipo "implementa" una interfaz automáticamente, con solo tener los métodos que la interfaz pide (esto se llama implementación implícita).

```go
type Abser interface {
    Abs() float64
}

type MyFloat float64
func (f MyFloat) Abs() float64 { ... }

type Vertex struct{ X, Y float64 }
func (v *Vertex) Abs() float64 { ... } // ojo: el receiver es *Vertex, no Vertex

var a Abser
a = MyFloat(123.45) // OK, MyFloat implementa Abser
a = &Vertex{3, 4}    // OK, *Vertex implementa Abser

a = Vertex{3, 4} // ERROR: Vertex (sin puntero) NO implementa Abser,
                 // porque el método Abs está definido sobre *Vertex
```

**Ojo (trampa muy típica de examen):** si un método está definido con receiver puntero (`func (v *Vertex) Abs()`), entonces **solo el puntero** (`*Vertex`) implementa la interfaz. El valor "pelado" (`Vertex`) no la implementa.

### Cómo es "por dentro" un valor de interfaz

Se puede pensar como un par `(value, type)`: el tipo concreto que tiene adentro, y el valor concreto de ese tipo. Llamar a un método de la interfaz en realidad ejecuta el método del tipo concreto que tiene guardado.

- El valor concreto puede ser `nil` (por ejemplo, un puntero nulo de cierto tipo).
- Pero un valor de interfaz que es `nil` de verdad (sin ningún tipo asignado) es distinto: no tiene ni tipo ni valor.

**Ojo:** esta es una trampa clásica de Go — una interfaz que "contiene" un puntero nil **no es lo mismo** que una interfaz nil pura. Por ejemplo, si guardás un `*T` nulo dentro de una interfaz, la interfaz en sí ya no es `== nil`, porque tiene un tipo asociado (aunque el valor sea nil).

Ejemplo completo (el que trae la teoría):
```go
type I interface{ M() }

type T struct{ S string }
func (t *T) M() {
    if t == nil {
        fmt.Println("nil!")
        return
    }
    fmt.Println(t.S)
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}

var i I
describe(i)          // (<nil>, <nil>)  -- acá sí es una interfaz nil "de verdad"
i.M()                // runtime error, no hay tipo concreto ni método que llamar

var t *T
i = t
describe(i)          // (<nil>, *main.T) -- ya tiene un tipo (*T), aunque el valor sea nil
i.M()                // imprime "nil!" -- el método SÍ se puede llamar, porque M() chequea t == nil por dentro

i = &T{"hello"}
describe(i)          // (&{hello}, *main.T)
i.M()                // hello
```
La clave: `describe(i)` con `i = t` (un `*T` nil) muestra `(<nil>, *main.T)` — tiene tipo `*main.T` aunque el valor sea nil. Es distinto de la primera `describe(i)`, que muestra `(<nil>, <nil>)` porque ahí la interfaz no tiene ningún tipo asignado. Por eso llamar a un método sobre la primera revienta (`runtime error`) pero sobre la segunda funciona (porque `M()` está preparado para recibir un receiver nil y lo chequea explícitamente).

### Empty interface (`interface{}`)

Una interfaz vacía no pide ningún método, así que puede guardar un valor de **cualquier tipo**. Por eso funciones como `fmt.Print` aceptan cualquier cosa: reciben parámetros `interface{}`.

```go
func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

var i interface{}
describe(i)        // (<nil>, <nil>)
i = 42
describe(i)        // (42, int)
i = "hello"
describe(i)        // (hello, string)
```

### Type assertion

Sirve para "recuperar" el valor concreto que está adentro de una interfaz:

```go
var i interface{} = "hello"

s := i.(string)       // si el tipo no coincide, esto genera un panic (error runtime)
s, ok := i.(string)    // forma segura: ok te dice si funcionó, sin panic
f, ok := i.(float64)   // acá ok da false, porque i guarda un string, no un float64
```

**Ojo:** siempre que no estés 100% seguro del tipo, usá la forma con `ok` para evitar que el programa explote en tiempo de ejecución.

### Type switch

Es como un `switch`, pero sobre el tipo concreto que hay adentro de una interfaz:

```go
func do(i interface{}) {
    switch v := i.(type) {
    case nil:
        fmt.Printf("Nil: %v\n", v)
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}
```

### Stringer interface

Es una interfaz muy usada, ya definida en el paquete `fmt`, que solo pide un método `String() string`. Si tu tipo lo tiene, `fmt.Println` (y similares) lo usan automáticamente para mostrarlo lindo:

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 42}
    fmt.Println(a) // usa el String() de Person automáticamente
}
```

## Posibles huecos / revisar

Ninguno — la página 25 (interfaz con puntero nil) se revisó y completó a mano, el resto se leyó completo sin problemas.

---

# Go - Manejo de Errores y Funciones (Go-5)

Este deck tiene dos grandes bloques: **manejo de errores** (cómo Go maneja los errores esperables sin excepciones) y **funciones como valores** (function values, funciones anónimas, variádicas y defer), y cierra con **panic/recover** que es el mecanismo para errores "no esperables".

## Manejo de Errores

En Go los errores son considerados algo normal y esperable, no algo excepcional. Por eso, las funciones que pueden fallar simplemente devuelven un **valor extra** además de su resultado normal.

Si el error tiene una sola causa posible, alcanza con devolver un booleano (o similar):

```go
value, ok := cache.Lookup(key)
if !ok {
    // ...cache[key] no existe...
}
```

Si el error puede tener varias causas distintas, se devuelve un valor de tipo `error`. `error` es una interfaz (built-in) con un solo método:

```go
type error interface {
    Error() string
}
```

Ejemplo típico de chequeo de error:

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
fmt.Println("Entero convertido: ", i)
```

**Ojo:** el patrón `if err != nil { ... }` aparece todo el tiempo en Go. Es LA forma estándar de chequear errores, no hay try/catch.

## Estrategias de Manejo de Errores

Cuando una función devuelve un error, es responsabilidad de quien la llama (el "llamador") chequearlo y decidir qué hacer. Go no obliga a nada automáticamente, así que hay varias estrategias posibles:

### 1. Propagación del error

Simplemente devolver el mismo error hacia arriba, para que lo maneje quien llamó a la función:

```go
func AddStr(s1, s2 string) (string, error) {
    i1, err := strconv.Atoi(s1)
    if err != nil {
        return "", err
    }
    i2, err := strconv.Atoi(s2)
    if err != nil {
        return "", err
    }
    return strconv.Itoa(i1 + i2), nil
}
```

```go
func main() {
    s, err := AddStr("42", "28a")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Println("Suma: ", s)
}
// Salida: Error: strconv.Atoi: parsing "28a": invalid syntax
```

### 2. Reemplazo del error

En vez de propagar el error tal cual, se crea un error nuevo con un mensaje más claro usando `fmt.Errorf`:

```go
func AddStr(s1, s2 string) (string, error) {
    i1, err := strconv.Atoi(s1)
    if err != nil {
        return "", fmt.Errorf("convirtiendo %s", s1)
    }
    i2, err := strconv.Atoi(s2)
    if err != nil {
        return "", fmt.Errorf("convirtiendo %s", s2)
    }
    return strconv.Itoa(i1 + i2), nil
}
// Salida: Error: convirtiendo 28a
```

`fmt.Errorf` arma el mensaje de error usando `fmt.Sprintf` y devuelve un `error` nuevo.

### 3. Reintentar la operación

Si el error puede ser transitorio (por ejemplo, un servidor que todavía no está listo), se puede reintentar con una espera creciente entre intentos (exponential back-off):

```go
func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil // éxito
        }
        log.Printf("server not responding (%s); retrying...", err)
        time.Sleep(time.Second << uint(tries)) // espera exponencial
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
```

### 4. Terminación controlada

Si es imposible seguir sin recuperarse del error, se corta la ejecución del programa:

```go
if err := WaitForServer(url); err != nil {
    fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
    os.Exit(1)
}

// o de forma más corta con log.Fatalf:
if err := WaitForServer(url); err != nil {
    log.Fatalf("Site is down: %v\n", err)
}
```

### 5. Registrar el error y continuar

A veces el error no es tan grave, entonces se anota (log) y el programa sigue, quizás deshabilitando alguna funcionalidad:

```go
if err := Ping(); err != nil {
    log.Printf("ping failed: %v; networking disabled", err)
}
```

### 6. Ignorar el error

También se puede directamente ignorar el error y seguir con un valor por defecto:

```go
content, err := os.ReadFile("data.json")
if err != nil {
    content = []byte("Datos a usar en caso de haber error al leer el archivo")
}
fmt.Println(string(content))
```

**Ojo:** ignorar el error es válido en algunos casos, pero hay que hacerlo a propósito y sabiendo el riesgo — no por olvidarse de chequear `err`.

## Package errors

El paquete `errors` es el que provee el constructor más simple para crear errores, `errors.New`. Por dentro es muy simple: un struct con un string y un método `Error()`:

```go
package errors

func New(text string) error { return &errorString{text} }
type errorString struct { text string }
func (e *errorString) Error() string { return e.text }
```

Y `fmt.Errorf` (que ya vimos arriba) está construido encima de `errors.New`:

```go
package fmt

import "errors"

func Errorf(format string, args ...interface{}) error {
    return errors.New(Sprintf(format, args...))
}
```

## Function Values

En Go, las funciones son valores como cualquier otro: tienen un tipo, se pueden guardar en variables, pasar como parámetro o devolver como resultado.

```go
func square(n int) int   { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

f := square
fmt.Println(f(3)) // "9"

f = negative
fmt.Println(f(3))       // "-3"
fmt.Printf("%T\n", f)   // "func(int) int"

f = product // ERROR de compilación: no se puede asignar func(int, int) int a func(int) int
```

**Ojo:** el tipo de la función importa. `f` tiene tipo `func(int) int`, así que no le podés asignar `product` (que es `func(int, int) int`) — no compila.

El valor por defecto de una variable de tipo función es `nil`:

```go
var f func(int) int
f(3) // runtime error: call of nil function

// por eso conviene chequear antes de llamar:
var f func(int) int
if f != nil {
    f(3)
}
```

**Ojo:** llamar a una función `nil` no da error de compilación, sino un **error en tiempo de ejecución** (panic). Hay que chequear `!= nil` antes si no estás seguro de que la variable tenga una función asignada.

Ejemplo de función que recibe otras funciones como parámetro (muy común para recorrer estructuras, como un árbol):

```go
func forEachNode(n *Node, pre, post func(n *Node)) {
    if pre != nil {
        pre(n)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    if post != nil {
        post(n)
    }
}
```

## Anonymous Functions

Una función anónima es una función sin nombre, que se puede escribir "en el lugar" donde se necesita, por ejemplo como argumento de otra función:

```go
s := strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")

// o guardarla en una variable para usarla después:
rot13 := func(r rune) rune {
    switch {
    case r >= 'A' && r <= 'Z':
        return 'A' + (r-'A'+13)%26
    case r >= 'a' && r <= 'z':
        return 'a' + (r-'a'+13)%26
    }
    return r
}
s := strings.Map(rot13, "HAL-9000")
```

Un detalle importante: una función anónima puede "capturar" variables del entorno donde fue creada, y esas variables persisten mientras la función exista (esto se llama clausura o *closure*):

```go
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```

**Ojo:** cada vez que llamás a `squares()` obtenés una `x` nueva e independiente. Pero si guardás el resultado en `f` y llamás a `f()` varias veces, todas esas llamadas comparten la misma `x` (por eso el resultado va aumentando: 1, 4, 9, 16 en vez de repetir 1).

## Variadic Functions

Una función variádica es una que puede recibir una cantidad variable de parámetros del mismo tipo, usando `...` antes del tipo:

```go
func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

fmt.Println(sum())           // "0"
fmt.Println(sum(3))          // "3"
fmt.Println(sum(1, 2, 3, 4)) // "10"

values := []int{1, 2, 3, 4}
fmt.Println(sum(values...))  // "10"
```

**Ojo:** si ya tenés un slice (`[]int`) y querés pasarlo a una función variádica, hay que agregarle `...` al final (`values...`), si no da error de compilación.

## Deferred Function Calls

`defer` sirve para decir "ejecutá esto, pero recién cuando la función actual esté por terminar" (ya sea que termine normal o por un panic). Se usa muchísimo para liberar recursos, como cerrar un archivo:

```go
// función ReadFile del package "os"
func ReadFile(name string) ([]byte, error) {
    f, err := Open(name)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    // calcular el tamaño
    ...
    // leer el contenido del archivo
    ...
}
```

**Ojo:** `defer f.Close()` se escribe justo después de abrir el archivo (aunque se ejecute al final), así es imposible olvidarse de cerrarlo, sin importar por dónde salga la función (incluso si hay un `return` en el medio o un panic).

## Panic

Cuando Go detecta un error en tiempo de ejecución que no puede manejar, entra en **pánico** (panic): la ejecución normal se detiene, se ejecutan las funciones diferidas (`defer`) pendientes, y el programa termina mostrando un mensaje de error con el valor del panic.

El programador también puede generar un panic a propósito con la función built-in `panic`, que acepta un parámetro de cualquier tipo:

```go
func Reset(x *Buffer) {
    if x == nil {
        panic("x is nil") // Innecesario acá, salvo que sea para dar un mejor mensaje
    }
    x.elements = nil
}
```

**Ojo (regla de oro):** usar `error` para errores "esperables" (los que forman parte del comportamiento normal de la función), y usar `panic` solo para errores "no esperables" (bugs, situaciones que no deberían pasar nunca). No hay que abusar de `panic` como si fuera un manejo de errores normal.

## Recover

Ante un `panic`, es posible "recuperar" la ejecución (o al menos dejar todo prolijo antes de terminar) usando `recover` dentro de una función diferida:

```go
func Parse(input string) (s *Syntax, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error: %v", p)
        }
    }()
    // ...parser...
}
```

Reglas clave sobre `recover`:
- Si `recover` se invoca dentro de una función diferida (`defer`), y la función que contiene ese `defer` entra en pánico, `recover` frena el estado de pánico y devuelve el valor que se le pasó a `panic`.
- La función que había entrado en pánico no continúa desde donde estaba; termina "normalmente" (como si hubiera hecho un return).
- Si `recover` se invoca en cualquier otro momento (fuera de un defer, o sin que haya panic activo), no tiene ningún efecto y devuelve `nil`.

**Ojo:** `recover` solo sirve dentro de una función `defer`. Si lo llamás directamente en el medio del código normal, no hace nada.

## Posibles huecos / revisar

- **Página 21 (Go-5.pdf):** la diapositiva solo muestra el título "Panic" sin ningún contenido visible (aparenta ser una slide de transición o con una animación de PowerPoint que no se capturó al convertir a PDF). El contenido real de "Panic" parece estar en la página 22, así que probablemente no se perdió nada importante, pero convendría chequear el .pptx original por las dudas.

---

# Go-6: Genéricos

## El problema que resuelven
Sin genéricos, si necesitás la misma lógica para distintos tipos, tenés que duplicar la función una vez por tipo:
```go
func SumInts(m map[string]int64) int64 { ... }
func SumFloats(m map[string]float64) float64 { ... }
```
Son casi idénticas, solo cambia el tipo. Los genéricos evitan esa duplicación: escribís la función una sola vez, y el tipo se completa recién cuando la usás.

## Funciones genéricas
```go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```
- `[K comparable, V int64 | float64]` son los **parámetros de tipo** (type parameters), entre corchetes, antes de los parámetros normales.
- `comparable`: constraint (restricción) que dice "cualquier tipo que se pueda comparar con `==`/`!=`". Se le pide esto a `K` porque `K` es la clave de un `map`, y en Go las claves de un map siempre tienen que ser comparables.
- `int64 | float64`: constraint que dice "el tipo tiene que ser `int64` **o** `float64`, nada más" (unión de tipos permitidos con `|`).
- `any`: constraint que acepta cualquier tipo (equivalente a no poner restricción).

**Llamada explícita vs inferida:**
```go
SumIntsOrFloats[string, int64](ints)   // explícito: le decís vos los tipos
SumIntsOrFloats(ints)                   // Go infiere K=string, V=int64 mirando el argumento
```
Casi siempre alcanza con la forma inferida (sin corchetes) — Go deduce los tipos mirando qué le pasaste como argumento.

**Otros ejemplos de constraints:**
- `[T any]`: cualquier tipo.
- `[T comparable]`: cualquier tipo comparable con `==`.
- `[T Stringer]`: cualquier tipo que implemente la interfaz `Stringer`.
- `[T int | int16 | int32 | int64 | int8 | float32 | float64]`: solo tipos numéricos de esa lista.
- `[T Stringer | error]`: tipos que implementen `Stringer` **o** `error`.

## Tipos genéricos
También se puede hacer genérico un `struct` completo, no solo una función.

### Lista genérica
```go
type List[T any] struct {
    first, last *node[T]
}
type node[T any] struct {
    val  T
    next *node[T]
}
```
**¿Por qué `node` también tiene que ser `[T any]` y no alcanza con que `List` lo sea?** Porque `node` guarda un valor de tipo `T` (`val T`) y un puntero a otro `node` del mismo tipo (`next *node[T]`). Si `node` no fuera genérico, tendría que fijar de antemano qué tipo va en `val` — y entonces `List` no podría ser realmente genérica, porque por dentro dependería de un tipo fijo. Los dos tienen que "viajar juntos" con el mismo parámetro de tipo `T`.

Uso:
```go
list := List[int]{}
list.PutOnFront(10)
list.PutOnTail(20)
fmt.Println(list.GetAll())  // [10 20] (según orden de inserción)
```
Los métodos de un tipo genérico se escriben repitiendo el parámetro de tipo: `func (l *List[T]) PutOnFront(v T) { ... }`.

### Árbol binario genérico
```go
type Tree[T any] struct {
    val         T
    left, right *Tree[T]
}
```
Como el árbol necesita comparar valores para decidir si van a la izquierda o la derecha, y `T any` no garantiza que se pueda comparar con `<`, la solución que usa la teoría es pasar la función de comparación como parámetro:
```go
func lt(x, y int) bool { return x <= y }

func (t *Tree[T]) insert(v T, f func(T, T) bool) *Tree[T] {
    if t == nil {
        return &Tree[T]{val: v}
    }
    if f(v, t.val) {
        t.left = t.left.insert(v, f)
    } else {
        t.right = t.right.insert(v, f)
    }
    return t
}
```
Así el árbol funciona con cualquier tipo `T`, y quien lo usa decide cómo se comparan los valores (acá con `lt` para enteros).

## Resumen para no confundir
- **Función genérica**: la lógica es la misma, cambia el tipo de dato que procesa (ej: sumar).
- **Tipo genérico**: la estructura de datos es la misma, cambia el tipo de lo que guarda adentro (ej: lista, árbol).
- Los dos usan la misma sintaxis de corchetes `[T constraint]`, la diferencia es si se la ponés a una función o a un `type`.

---

# Go-7: Concurrencia

## Concurrencia vs Paralelismo
- **Concurrencia**: poder estructurar un programa en tareas que están "en progreso" al mismo tiempo, pero que no necesariamente se ejecutan en el mismo instante — pueden ir turnándose en un solo core.
- **Paralelismo**: tareas que se ejecutan literalmente en el mismo instante. Para esto hacen falta varios cores.
- Podés tener concurrencia sin paralelismo (todo en 1 core, intercalado), pero no paralelismo sin concurrencia.
- Esto NO tiene que ver con si comparten o no recursos — esa es otra discusión (ver Exclusión mutua más abajo).

**Ejemplo simple:** dos amigos cocinando en la misma cocina turnándose la única hornalla = concurrencia sin paralelismo. Dos amigos cocinando cada uno en su propia hornalla al mismo tiempo = paralelismo.

## Goroutines
- Una goroutine es una función que se puede ejecutar "en simultáneo" con otras, usando la palabra `go`.
- `go miFuncion()` lanza la función y sigue de largo sin esperarla.
- El programa **no espera** a las goroutines que lanza — si `main()` termina, se cortan todas, hayan terminado o no.

```go
func main() {
    fmt.Println("Inicia main")
    go hello()
    fmt.Println("Termina main")
}
func hello() {
    fmt.Println("Hello desde goroutine")
}
```
Acá puede pasar que "Hello desde goroutine" ni se llegue a imprimir, porque `main` puede terminar antes de que la goroutine arranque. Por eso hace falta algo que sincronice — ahí entra WaitGroup.

## WaitGroup
- Es un contador de goroutines pendientes.
- `Add(n)`: suma `n` al contador (se llama antes de lanzar las goroutines).
- `Done()`: resta 1 al contador (se llama al final de cada goroutine, generalmente con `defer`).
- `Wait()`: bloquea hasta que el contador llegue a 0.

```go
var wg sync.WaitGroup

func trabajo(id int) {
    defer wg.Done()
    fmt.Println("goroutine", id)
}

func main() {
    wg.Add(3)
    go trabajo(1)
    go trabajo(2)
    go trabajo(3)
    wg.Wait()
    fmt.Println("todas terminaron")
}
```
**Ojo:** si te olvidás un `Done()` en algún camino del código, el contador nunca llega a 0 y `Wait()` se queda esperando para siempre (deadlock silencioso, no tira error).

## Channels (canales)
- Mecanismo para que las goroutines se comuniquen y se sincronicen entre sí.
- Se declaran con `make`: `ch := make(chan int)`.
- El "zero value" de un canal es `nil` (una variable canal sin `make` vale `nil`).
- **Send**: `ch <- valor` (mandar un valor al canal).
- **Receive**: `valor := <-ch` (recibir un valor del canal).
- Por defecto (canal sin buffer / "unbuffered"), tanto el send como el receive **bloquean** a quien los ejecuta hasta que el otro lado esté listo — es un punto de encuentro entre las dos goroutines.

```go
ch := make(chan string)
go func() { ch <- "hola" }()   // se bloquea hasta que alguien reciba
msg := <-ch                     // se bloquea hasta que alguien mande
fmt.Println(msg)
```

### Cerrar canales
- `close(ch)` cierra el canal: avisa que no van a llegar más valores.
- Al recibir, `v, ok := <-ch` — si `ok` es `false`, el canal está cerrado y ya no tiene más valores.
- `for v := range ch { ... }` recibe valores repetidamente hasta que el canal se cierra (deja de iterar solo).
- Solo la goroutine "emisora" (la que manda) debería cerrar el canal, nunca la que recibe.

### Canales unidireccionales
- `chan<- int`: canal solo para mandar (send-only).
- `<-chan int`: canal solo para recibir (receive-only).
- Sirve para dejar explícito en la firma de una función qué hace cada lado (ej: `func productor(out chan<- int)`).
- Intentar cerrar un canal receive-only es error de compilación.

### Canal `nil`
- Cualquier operación (send o receive) sobre un canal `nil` se bloquea **para siempre** (no da error, simplemente nunca avanza).
- Truco muy usado: en un `select`, asignarle `nil` a la variable de un canal "apaga" ese `case` (nunca va a estar listo), sin tener que reescribir el select. Se usa para guardas condicionales (ej: "solo puedo aceptar lectores si no hay nadie escribiendo").

```go
var guardaLectura chan int
if !hayEscritorActivo {
    guardaLectura = pedirLeer   // habilitado
} // si no, queda nil = deshabilitado
select {
case id := <-guardaLectura:
    // solo entra acá si guardaLectura no es nil
}
```

## Buffered channels
- `make(chan string, 3)`: canal con una cola interna de capacidad 3.
- El `send` solo bloquea si la cola ya está llena (3 elementos sin consumir).
- El `receive` solo bloquea si la cola está vacía.
- `cap(ch)` devuelve la capacidad total, `len(ch)` cuántos elementos hay en este momento esperando ser leídos.
- Útil para desacoplar productor y consumidor cuando no hace falta sincronización estricta en cada elemento (ej: productor-consumidor con buffer).

## Select
- `select` permite que una goroutine espere sobre **varios canales a la vez**, y elige el que tenga algo listo (send o receive).
- Si más de un `case` está listo al mismo tiempo, Go elige uno al azar — no hay orden de prioridad entre los cases.
- **`default`**: si se agrega un `case default`, el `select` deja de bloquear — si ningún canal está listo en ese instante, ejecuta el `default` en vez de esperar. Sirve para hacer sends/receives "no bloqueantes".

```go
select {
case val := <-ch1:
    fmt.Println("de ch1:", val)
case val := <-ch2:
    fmt.Println("de ch2:", val)
default:
    fmt.Println("nadie tenía nada listo ahora")
}
```

## Exclusión mutua — Mutex
- Problema: dos goroutines modificando la misma variable sin coordinarse → resultado impredecible (race condition).
- `sync.Mutex`: sección crítica exclusiva. Solo una goroutine a la vez puede estar entre un `Lock()` y su `Unlock()` correspondiente.
- `Lock()`: si ya está tomado, bloquea a quien lo llama hasta que se libere.
- `Unlock()`: libera. Si se llama sobre un mutex no bloqueado, es un error en tiempo de ejecución.
- **Un Mutex NO es reentrante**: si una goroutine ya tiene el lock y vuelve a llamar `Lock()` (aunque sea indirectamente, desde otra función), se bloquea a sí misma → deadlock.

**Ejemplo clásico de deadlock por doble lock (muy preguntado):**
```go
func Deposit(amount int) {
    mu.Lock()
    balance += amount
    mu.Unlock()
}
func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    Deposit(-amount)          // Deposit intenta Lock() de nuevo -> DEADLOCK
    if balance < 0 {
        Deposit(amount)
        return false
    }
    return true
}
```
**Solución:** separar la lógica en una función interna sin lock (`deposit`, minúscula) y que las funciones públicas tomen el lock una sola vez:
```go
func deposit(amount int) { balance += amount }   // sin lock, uso interno

func Deposit(amount int) {
    mu.Lock(); defer mu.Unlock()
    deposit(amount)
}
func Withdraw(amount int) bool {
    mu.Lock(); defer mu.Unlock()
    deposit(-amount)
    if balance < 0 { deposit(amount); return false }
    return true
}
```

## RWMutex (un escritor, múltiples lectores)
- `sync.RWMutex`: permite que **muchos lectores entren a la vez**, mientras nadie esté escribiendo.
- `RLock()` / `RUnlock()`: para lectores. Varios pueden tener el RLock simultáneamente.
- `Lock()` / `Unlock()`: para el escritor, exclusivo total (bloquea a todos, lectores y escritores).
- Conviene usarlo cuando hay muchas más lecturas que escrituras — con un `Mutex` común, hasta los lectores tendrían que hacer fila uno por uno.

```go
func Balance() int {
    mu.RLock()
    defer mu.RUnlock()
    return balance
}
func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    balance += amount
}
```

## Nota sobre el orden en Lock()
- Un Mutex/RWMutex **no es un canal** — no tiene buffer, no usa send/receive. El runtime de Go maneja internamente una cola de espera para las goroutines bloqueadas en `Lock()`.
- **No hay garantía estricta de orden (FIFO)** entre las goroutines que esperan un lock — el runtime puede dejar pasar a una que recién llega antes que a una que ya esperaba (prioriza throughput). Sí evita que alguna quede esperando para siempre (mecanismo anti-starvation).

## Problema de los fumadores (ejemplo de select + canales)
- 3 fumadores, cada uno tiene ilimitado de un ingrediente (papel, tabaco, fósforo). Un "dealer" pone en la mesa los dos ingredientes que le faltan a un fumador elegido al azar.
- **Por qué se usan canales y no variables compartidas:** si el dealer solo escribiera en variables comunes, no habría garantía de que el fumador lea el dato recién después de que el dealer terminó de escribirlo (race condition). Los canales dan dato + sincronización juntos: el fumador que hace `<-ingrediente` se bloquea hasta que el dealer efectivamente mande el valor, así nunca lee "antes de tiempo". Además, como cada fumador escucha su propio canal, solo se despierta el que corresponde.

```go
ingrediente := make(chan string)
go func() { ingrediente <- "papel" }()   // dealer
val := <-ingrediente                      // fumador, espera hasta que llegue
```

---

# Go-8: Problemas clásicos de concurrencia

## Condiciones para que exista un deadlock
Para que se dé un deadlock tienen que cumplirse **las 4 a la vez**:
1. **Exclusión mutua**: el recurso lo usa uno solo a la vez.
2. **Retención y espera**: un proceso tiene un recurso agarrado y, sin soltarlo, pide otro más.
3. **No apropiación**: nadie te puede sacar un recurso a la fuerza, solo lo soltás vos cuando querés.
4. **Espera circular**: A espera un recurso que tiene B, B espera uno que tiene C, ..., y el último espera uno que tiene A — se cierra el círculo.

Si rompés **cualquiera** de las 4, no puede haber deadlock. Las soluciones de los filósofos de abajo rompen la "espera circular".

## Los filósofos (Dining Philosophers)
- 5 filósofos en una mesa redonda, un tenedor entre cada par de filósofos adyacentes (5 tenedores en total).
- Para comer, un filósofo necesita **los dos** tenedores de al lado (izquierda y derecha).
- Cada tenedor es un recurso exclusivo (un `sync.Mutex`).

**Versión con deadlock (todos agarran primero el de la izquierda):**
```go
var forks = [5]sync.Mutex{}

func philosopher(id int, forkL, forkR *sync.Mutex) {
    for i := 0; i < 50; i++ {
        forkL.Lock()
        forkR.Lock()
        // come
        forkL.Unlock()
        forkR.Unlock()
    }
    dining.Done()
}

func main() {
    dining.Add(5)
    for i := range philos {
        go philosopher(i, &forks[i], &forks[(i+1)%5])
    }
    dining.Wait()
}
```
Problema: si los 5 agarran **al mismo tiempo** su tenedor de la izquierda, todos quedan esperando el de la derecha, que está en manos del vecino — nadie suelta nada. Es la "espera circular" en persona.

**Solución (romper la simetría):** que no todos agarren en el mismo orden. Por ejemplo, un filósofo agarra primero el de la derecha en vez de la izquierda (alternando según si el `id` es par o impar):
```go
go philosopher(i, &forks[(i+i%2)%5], &forks[(i+1-i%2)%5])
```
Al romper la simetría (no todos piden en el mismo orden), se corta la posibilidad de que se cierre el círculo de espera.

**Inanición (starvation):** distinto del deadlock — acá el programa no se cuelga, pero un filósofo en particular puede quedarse *sin comer nunca* porque sus vecinos siempre le "ganan" los tenedores. La solución de la teoría le agrega una penalización: si un filósofo comió mucho menos que sus vecinos, se le da prioridad (duerme menos tiempo antes de volver a intentar).

**Diferencia clave para el examen:** deadlock = todo se traba, nadie avanza más. Starvation = el programa sigue corriendo y otros avanzan, pero a uno en particular nunca le toca.

## El barbero durmiente (Sleeping Barber)
Planteo: un barbero con 1 silla de corte y una sala de espera con `n` sillas.
- Si el barbero está libre (durmiendo) y llega un cliente: lo despierta y lo atiende.
- Si el barbero está ocupado y hay silla libre en la sala de espera: el cliente se sienta a esperar.
- Si no hay silla libre: el cliente se va (no espera).

**Canales usados:**
- `sillas := make(chan string, n)`: buffered, representa las `n` sillas de la sala de espera. Si está lleno, un cliente no puede "sentarse" más.
- `despertar := make(chan string)`: unbuffered, un cliente lo usa para despertar al barbero si estaba dormido.
- `listo := make(chan bool)`: unbuffered, el barbero avisa por acá cuando terminó de cortar.

**La parte más importante para entender (el `select` con dos niveles del cliente):**
```go
select {
case despertar <- nombre:
    // el barbero estaba dormido, lo desperté y me atiende directo
default:
    select {
    case sillas <- nombre:
        // el barbero estaba ocupado, pero conseguí lugar en la sala de espera
        <-listo
    default:
        // no había lugar, me voy sin cortar
    }
}
```
Esto funciona porque un `select` con `default` es **no bloqueante**: el cliente primero intenta despertar al barbero (`despertar <- nombre`) sin esperar — si nadie está escuchando ese canal en este instante (porque el barbero no está durmiendo esperando ahí), cae al `default`. Ahí intenta sentarse en la sala de espera (`sillas <- nombre`), que como es buffered, entra sin bloquear mientras haya lugar — si está lleno, cae a su propio `default` y el cliente se va.

Esta es la aplicación más clara de "select + default" para modelar decisiones de "probá esto, si no se puede probá lo otro, si no se puede hacé un tercer camino" sin que ninguna goroutine quede bloqueada esperando algo que quizás nunca pase.

---

# Go - Packages, Modules y Dependencies

Este apunte junta las 3 ideas del deck: **packages** (paquetes), **modules** (módulos) y **dependencies** (dependencias). Son temas que van de la mano: un paquete es la unidad de organización del código, un módulo es la unidad que agrupa paquetes y sus versiones, y las dependencias son paquetes externos que tu módulo usa.

## Packages

Un **package** (paquete) es simplemente una carpeta con archivos `.go` que comparten la misma primera línea `package <nombre>`. Todo lo que está en esos archivos se puede usar entre sí sin necesidad de importar nada (están en el mismo paquete). Para usar un paquete desde otro lado, hay que importarlo.

Ejemplo: un proyecto con un paquete `tempconv` (con dos archivos, `tempconv.go` y `conv.go`) y un `main.go` que lo usa:

```go
// tempconv.go
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    BoilingC      Celsius = 100
)

func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g°F", f)
}
```

```go
// conv.go (mismo paquete tempconv)
package tempconv

// CToF convierte de Celsius a Fahrenheit.
func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

// FToC convierte de Fahrenheit a Celsius.
func FToC(f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}
```

```go
// main.go
package main

import (
    "fmt"
    tc "projecttempconv/tempconv" // le pusimos el alias "tc"
)

func main() {
    c := tc.BoilingC
    fmt.Println(c, "=", tc.CToF(c))
}
```

**Ojo:** si intentás correr `main.go` sin haber configurado el módulo, Go tira un error tipo:

```
main.go:5:2: package projecttempconv/tempconv is not in GOROOT (...)
```

Esto pasa porque Go necesita saber cuál es la "ruta base" de tu proyecto para poder resolver el import `"projecttempconv/tempconv"`. Para eso hace falta un **módulo** (ver abajo).

## Modules

Un **module** (módulo) es un conjunto de paquetes que se versionan y distribuyen juntos. Se define con el comando:

```
go mod init projecttempconv
```

Esto crea un archivo `go.mod` en la raíz del proyecto:

```
module projecttempconv

go 1.20
```

El `go.mod` le dice a Go "todo lo que está en esta carpeta para abajo se llama `projecttempconv`, y por eso el import `projecttempconv/tempconv` tiene sentido". Una vez que existe el `go.mod`, correr `go run .\main.go` funciona bien y muestra:

```
100°C = 212°F
```

**Ojo:** sin `go.mod` no hay módulo, y sin módulo Go no puede resolver los imports de tus propios paquetes ni instalar dependencias externas. Es el primer paso obligatorio de cualquier proyecto Go moderno.

## Dependencies

Las **dependencies** (dependencias) son paquetes de terceros (que no escribiste vos) que tu código importa. Se agregan con `go get`:

```
go get -u rsc.io/quote
```

Y después simplemente se importan como cualquier otro paquete:

```go
package main

import (
    "fmt"
    tc "projecttempconv/tempconv"
    "rsc.io/quote" // <- dependencia externa
)

func main() {
    c := tc.BoilingC
    fmt.Println(c, "=", tc.CToF(c), quote.Hello())
}
```

Cuando agregás una dependencia, Go actualiza el `go.mod` agregando un bloque `require` con el nombre del paquete y la versión exacta que se está usando:

```
module projecttempconv

go 1.20

require (
    golang.org/x/text v0.9.0 // indirect
    rsc.io/quote v1.5.2 // indirect
    rsc.io/sampler v1.99.99 // indirect
)
```

**Ojo:** vas a ver que aparecen dependencias que vos no pediste explícitamente (como `golang.org/x/text` o `rsc.io/sampler`), marcadas como `// indirect`. Son dependencias de tus dependencias (dependencias transitivas): `rsc.io/quote` necesita esos paquetes para funcionar, y Go los anota igual en el `go.mod` aunque tu código nunca los importe directamente.

---

# Posibles preguntas de examen — Todos los temas

---

## Go-1: Básicos, variables, tipos, operadores

**P: ¿Cuándo se puede usar `:=` y cuándo hay que usar `var`?**
R: `:=` solo funciona **dentro de una función**. A nivel de package (afuera de cualquier función), toda declaración tiene que empezar con una palabra clave (`var`, `func`, `const`, etc.), así que ahí no se puede usar `:=`.

**P: ¿Por qué `Celsius` y `Fahrenheit` (definidos como `type Celsius float64` / `type Fahrenheit float64`) no se pueden sumar directamente entre sí ni con un `float64` normal?**
R: Un tipo nombrado (named type) es un tipo **distinto** de su tipo subyacente, aunque tengan la misma estructura interna. Go no hace conversión automática — hay que convertir explícitamente con `T(v)`.

**P: ¿`&&` y `||` en Go son short-circuit (cortan la evaluación apenas se sabe el resultado)?**
R: Sí. En `a && b`, si `a` es `false`, ni siquiera se evalúa `b` (ya se sabe que el resultado es `false`). En `a || b`, si `a` es `true`, no se evalúa `b`. Esto importa cuando el segundo operando tiene efectos secundarios o podría fallar (ej: `p != nil && p.Valor > 0` — si `p` es nil, ni se intenta leer `p.Valor`).

**P: ¿Cuál es el tipo numérico "por defecto" que conviene usar en Go si no hay una razón especial para otro?**
R: `int`.

---

## Go-2: Control de flujo, funciones, fmt

**P: ¿Qué pasa con una variable declarada en la sentencia de inicialización de un `if` (`if v := algo(); v < x {...}`) fuera del bloque `if/else`?**
R: Deja de existir. Su alcance (scope) es solo ese `if` y sus `else` asociados.

**P: ¿El `switch` de Go hace fallthrough automático entre casos, como en C/Java?**
R: No. Cada `case` corta solo al terminar, no hace falta `break`. Si de verdad querés pasar al siguiente caso, hay que pedirlo explícitamente con la palabra `fallthrough`.

**P: Diferencia entre `fmt.Print` y `fmt.Println`.**
R: `Println` siempre pone espacio entre todos los argumentos y agrega salto de línea al final. `Print` solo agrega espacio entre dos argumentos si **ninguno de los dos** es un string, y no agrega salto de línea.

**P: ¿Qué devuelven las funciones `Scan`/`Scanf`/`Scanln` (y las variantes `Sscan...`)?**
R: Dos valores: cuántos elementos lograron leer (`n int`) y un posible error (`err error`).

**P: ¿Qué problema tiene usar `Scanln`/`Sscanln` si el texto tiene un salto de línea antes de terminar de leer todos los valores esperados?**
R: Corta con error `"unexpected newline"` y las variables que faltaban leer se quedan con el valor que tenían antes (no se pisan con nada).

---

## Go-3: Arrays, Slices, Maps

**P: Diferencia fundamental entre array y slice.**
R: Un array tiene tamaño fijo (el tamaño es parte de su tipo: `[5]int` y `[6]int` son tipos distintos). Un slice es una "ventana" sobre un array de atrás, con `len` (tamaño actual) y `cap` (espacio disponible antes de necesitar un array nuevo), y puede crecer con `append`.

**P: Si tengo `s := a[2:4]` (un slice de un array `a`) y modifico `s[0]`, ¿se modifica `a`?**
R: Sí. Un slice no tiene copia propia de los datos, apunta al mismo array de atrás. Modificar el slice modifica el array original, y viceversa.

**P: ¿Qué diferencia hay entre `dst := src[:n]` y hacer `copy(dst, src[:n])`?**
R: `src[:n]` sigue compartiendo el array de atrás con `src` (modificar uno afecta al otro). `copy` hace una copia real e independiente de los valores.

**P: ¿Qué pasa si intentás escribir en un map que vale `nil` (nunca se le hizo `make`)?**
R: Error en tiempo de ejecución. Hay que inicializarlo con `make(map[K]V)` antes de poder escribir.

**P: ¿Qué tipos NO pueden ser clave de un map, y por qué?**
R: Slices, maps y funciones — porque no se pueden comparar con `==`, y las claves de un map necesitan ser comparables.

**P: ¿Está garantizado el orden en que `range` recorre un map?**
R: No. Puede cambiar entre ejecuciones. (Ojo: esto es distinto de `fmt.Println(m)`, que sí muestra el map ordenado alfabéticamente por clave, pero eso es solo cómo se imprime, no cómo se recorre internamente).

---

## Go-4: Punteros, Tipos, Structs, Interfaces

**P: ¿Por qué si tengo `func (v *Vertex) Abs() float64`, un valor `Vertex{3,4}` (sin puntero) no puede usarse donde se espera una interfaz que pide `Abs()`?**
R: Porque el método está definido con receiver puntero (`*Vertex`), así que solo `*Vertex` implementa esa interfaz — el valor "pelado" `Vertex` no.

**P: Si un método tiene receiver puntero, ¿hace falta escribir `(&v).Metodo()` para llamarlo desde una variable valor `v`?**
R: No, Go agrega el `&` automáticamente (siempre que `v` sea una variable direccionable). Por eso `v.Metodo()` funciona igual que `(&v).Metodo()`.

**P: ¿Qué diferencia hay entre una interfaz `nil` y una interfaz que contiene un puntero `nil` de algún tipo (ej: guardaste un `*T` nulo adentro)?**
R: Son distintas. Una interfaz `nil` de verdad no tiene ni tipo ni valor asociado (`(<nil>, <nil>)`). Una interfaz que guarda un `*T` nulo sí tiene un tipo asociado (`*T`), aunque el valor sea nil (`(<nil>, *main.T)`). Por eso llamar a un método sobre la primera da error en tiempo de ejecución, pero sobre la segunda funciona si el método está preparado para recibir un receiver nil.

**P: Diferencia entre type assertion (`i.(T)`) y type switch (`switch v := i.(type)`).**
R: La type assertion chequea/extrae un tipo concreto puntual (y puede hacer panic si no coincide, salvo que uses la forma con `, ok`). El type switch compara contra varios tipos posibles a la vez, como un switch normal pero sobre el tipo dinámico de la interfaz.

**P: ¿Struct embedding es lo mismo que herencia?**
R: No. Es composición: el struct de afuera tiene "adentro" una copia del struct embebido. Go simplemente te deja acceder a los campos/métodos del tipo embebido sin escribir el camino completo (forma "promovida"), pero no hay polimorfismo de herencia real.

---

## Go-5: Errores, funciones como valores, panic/recover

**P: ¿Cómo maneja Go los errores, a diferencia de lenguajes con try/catch?**
R: No hay excepciones. Una función que puede fallar devuelve un valor extra de tipo `error` (o un booleano si hay una sola causa posible), y quien la llama decide qué hacer chequeando `if err != nil`.

**P: Nombrá las estrategias para manejar un error una vez que lo recibiste.**
R: (1) Propagarlo tal cual hacia arriba. (2) Reemplazarlo por uno más claro con `fmt.Errorf`. (3) Reintentar la operación (con espera creciente). (4) Terminación controlada (`os.Exit`/`log.Fatalf`). (5) Registrarlo (log) y continuar. (6) Ignorarlo a propósito.

**P: ¿Qué pasa si llamás a una variable de tipo función que vale `nil`?**
R: Error en tiempo de ejecución (panic), no error de compilación.

**P: En el ejemplo de `squares()` que devuelve una función que cuenta (closure), ¿por qué llamar `f()` varias veces da 1, 4, 9, 16 en vez de repetir 1 cada vez?**
R: Porque la función anónima "capturó" la variable `x` de su entorno (closure), y esa `x` persiste asociada a esa instancia de función devuelta. Cada llamada a `squares()` crea una `x` nueva, pero múltiples llamadas a la misma `f` comparten la misma `x`.

**P: ¿Cómo se pasa un slice ya existente a una función variádica?**
R: Agregándole `...` al final: `sum(values...)`. Sin eso, no compila.

**P: ¿Cuándo conviene usar `panic` en vez de devolver un `error`?**
R: `error` es para fallas esperables (parte del comportamiento normal). `panic` es solo para errores que no deberían pasar nunca (bugs). No hay que usar panic como manejo de errores normal.

**P: ¿Dónde funciona `recover()`?**
R: Solo dentro de una función `defer`. Si se llama afuera de un defer, o sin que haya un panic activo, no hace nada y devuelve `nil`.

---

## Go-8: Deadlock y problemas clásicos

**P: ¿Cuáles son las 4 condiciones necesarias para que exista un deadlock?**
R: Exclusión mutua, retención y espera, no apropiación, y espera circular. Rompiendo cualquiera de las 4 se evita el deadlock.

**P: En los filósofos, ¿por qué se traba si todos agarran primero el tenedor de la izquierda?**
R: Porque los 5 pueden agarrar su tenedor izquierdo al mismo tiempo, y todos quedan esperando el derecho (que tiene el vecino) — se cierra un círculo de espera circular.

**P: ¿Cómo se evita ese deadlock en la solución de la teoría?**
R: Rompiendo la simetría — no todos piden los tenedores en el mismo orden (por ejemplo alternando según si el filósofo tiene id par o impar).

**P: ¿Qué es la inanición (starvation) y en qué se diferencia del deadlock?**
R: Starvation es cuando el programa sigue corriendo normalmente (otros avanzan), pero un proceso en particular nunca consigue lo que necesita. Deadlock es que todo se traba, nadie avanza más.

**P: En el barbero durmiente, ¿para qué se usa un `select` con `default` anidado en el lado del cliente?**
R: Para probar alternativas sin bloquear: primero intenta despertar al barbero (si nadie escucha ahí, cae al `default`), después intenta sentarse en la sala de espera (si está llena, cae a otro `default` y se va). Así ninguna goroutine queda esperando algo que quizás nunca pase.

---

## Go-Package: Packages, Modules, Dependencies

**P: ¿Qué es un package en Go?**
R: Una carpeta con archivos `.go` que comparten la misma primera línea `package <nombre>`. Todo lo que está en esos archivos se puede usar entre sí sin importar nada.

**P: ¿Para qué sirve `go mod init`?**
R: Crea el archivo `go.mod`, que define el nombre del módulo (la "ruta base" del proyecto) para que Go pueda resolver los imports entre tus propios paquetes y manejar las dependencias externas.

**P: ¿Qué significa que una dependencia en `go.mod` esté marcada `// indirect`?**
R: Que no la importaste vos directamente — es una dependencia de una de tus dependencias (dependencia transitiva), y Go la anota igual para saber exactamente qué versiones está usando todo el árbol de paquetes.

---

## Concurrencia y Genéricos

(Preguntas de Go-6 y Go-7, ya cubiertas en la sesión de repaso)

**P: ¿Cuál es la diferencia entre concurrencia y paralelismo?**
R: Concurrencia es poder estructurar tareas que están "en progreso" al mismo tiempo, sin que necesariamente se ejecuten en el mismo instante (pueden turnarse en un solo core). Paralelismo es que se ejecuten literalmente en simultáneo, lo cual requiere varios cores. Podés tener concurrencia sin paralelismo, pero no al revés.

**P: ¿Qué es una goroutine?**
R: Una función que se lanza con `go` y corre "en paralelo/concurrente" con el resto del programa, sin que quien la lanzó espere a que termine.

**P: ¿Qué hace `sync.WaitGroup`? ¿Qué pasa si te olvidás un `Done()`?**
R: Es un contador de goroutines pendientes. `Add(n)` suma, `Done()` resta 1, `Wait()` bloquea hasta que llegue a 0. Si falta un `Done()`, el contador nunca llega a 0 y `Wait()` queda esperando para siempre (deadlock silencioso).

**P: ¿Qué diferencia hay entre un canal buffered y uno unbuffered?**
R: Unbuffered: el send se bloquea hasta que haya alguien recibiendo en ese mismo instante (sincronización). Buffered (capacidad N): el send solo se bloquea si ya hay N elementos sin consumir; si hay lugar, no espera a nadie.

**P: ¿Qué pasa si leés o escribís un canal `nil`?**
R: La operación se bloquea para siempre (no da error). Se usa a propósito en un `select` para "apagar" un case: si el canal de ese case es `nil`, ese case nunca puede estar listo, entonces el select nunca lo elige.

**P: ¿Para qué sirve `default` en un `select`?**
R: Hace que el select sea no bloqueante: si ningún otro case está listo en ese instante, se ejecuta `default` en vez de quedarse esperando.

**P: Explicá este código y qué problema tiene:**
```go
func Deposit(amount int) {
    mu.Lock()
    balance += amount
    mu.Unlock()
}
func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    Deposit(-amount)
    if balance < 0 {
        Deposit(amount)
        return false
    }
    return true
}
```
R: Deadlock por mutex no reentrante. `Withdraw` toma el lock y, sin soltarlo, llama a `Deposit`, que intenta tomar el mismo lock de nuevo → se bloquea a sí mismo esperando algo que nunca va a pasar (solo él mismo podría liberarlo, pero está trabado). Solución: separar la lógica en una función interna sin lock, y que Deposit/Withdraw tomen el lock una sola vez.

**P: ¿Diferencia entre `sync.Mutex` y `sync.RWMutex`? ¿Cuándo conviene el segundo?**
R: Mutex: exclusión total, uno a la vez (sea para leer o escribir). RWMutex: permite muchos lectores simultáneos (`RLock`) mientras nadie escriba; el escritor (`Lock`) sigue siendo exclusivo total. Conviene cuando hay muchas más lecturas que escrituras.

**P: ¿Está garantizado el orden en que las goroutines bloqueadas obtienen un `Lock()`?**
R: No hay garantía estricta de FIFO (el runtime puede dejar pasar a una que recién llega antes que una que esperaba desde antes), pero sí evita que alguna quede esperando para siempre.

**P: ¿Cuáles son las 4 condiciones necesarias para que exista un deadlock?**
R: Exclusión mutua, retención y espera, no apropiación, espera circular. Alcanza con romper una sola para evitarlo.

**P: En el problema de los filósofos, ¿por qué se cuelga si todos agarran primero el tenedor de la izquierda?**
R: Si los 5 agarran su tenedor izquierdo al mismo tiempo, todos quedan esperando el derecho, que lo tiene el vecino — se cierra un círculo de espera (espera circular) y nadie suelta nada.

**P: ¿Cómo se evita el deadlock de los filósofos?**
R: Rompiendo la simetría: no todos piden los tenedores en el mismo orden (por ejemplo, alternar según si el número de filósofo es par o impar).

**P: ¿Qué diferencia hay entre deadlock y starvation (inanición)?**
R: Deadlock: todo se traba, nadie avanza más. Starvation: el programa sigue funcionando y otros procesos avanzan normalmente, pero uno en particular nunca consigue lo que necesita.

**P: En el barbero durmiente, ¿para qué sirve un `select` con `default` anidado?**
R: Para probar varias alternativas sin bloquear: el cliente primero intenta despertar al barbero sin esperar; si no puede (`default`), intenta sentarse en la sala de espera sin esperar; si tampoco puede (otro `default`), se va. Ninguna goroutine queda bloqueada esperando algo que quizás nunca pase.

**P: ¿Qué problema resuelve una función genérica respecto de escribir una función por tipo?**
R: Evita duplicar código casi idéntico para cada tipo (ej: `SumInts` y `SumFloats`). Con genéricos escribís la lógica una sola vez y el tipo se completa al llamarla.

**P: En `func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V`, ¿qué es `comparable` y por qué se lo exige a `K`?**
R: `comparable` es una restricción (constraint) que dice "el tipo se puede comparar con `==`". Se le exige a `K` porque `K` es la clave de un map, y en Go las claves de un map siempre tienen que ser comparables.

**P: En una lista genérica `List[T any]` con nodos `node[T any]`, ¿por qué el nodo también tiene que ser genérico?**
R: Porque el nodo guarda un valor de tipo `T` y un puntero a otro nodo del mismo tipo. Si `node` no fuera genérico, tendría que fijar de antemano un tipo concreto, y entonces `List` no sería realmente genérica por dentro.

**P: ¿Hace falta escribir siempre los parámetros de tipo entre corchetes al llamar una función genérica?**
R: No, casi siempre Go los infiere solo mirando los argumentos (`SumIntsOrFloats(ints)` en vez de `SumIntsOrFloats[string, int64](ints)`).

---

