# Multiple choice — práctica para el escrito

30 preguntas tipo opción múltiple, basadas en los "ojo" más comunes de cada tema. Las respuestas están en `multiple-choice-soluciones.md`, aparte. Iban probadas con `go run` donde había código.

---

**1.** ¿Cuál es la diferencia correcta entre concurrencia y paralelismo?
A) Son sinónimos, Go los usa indistintamente.
B) Concurrencia es estructurar tareas que pueden estar "en progreso" a la vez, sin que se ejecuten necesariamente en el mismo instante; paralelismo sí requiere ejecución simultánea real (varios cores).
C) Paralelismo es posible sin concurrencia, pero no al revés.
D) Concurrencia siempre implica usar más de un core.

**2.**
```go
type Kilometros float64
type Millas float64

func main() {
    var k Kilometros = 10
    var m Millas = 5
    total := k + m
    fmt.Println(total)
}
```
A) Imprime `15`.
B) Imprime `15.000000`.
C) No compila: son tipos nombrados distintos, incompatibles aunque compartan tipo subyacente.
D) Compila pero da error en tiempo de ejecución.

**3.** ¿`&&` y `||` en Go son short-circuit?
A) No, siempre evalúan ambos lados.
B) Sí: en `a && b`, si `a` es `false` no se evalúa `b`; en `a || b`, si `a` es `true` no se evalúa `b`.
C) Solo `&&` es short-circuit, `||` no.
D) Depende del tipo de dato.

**4.**
```go
if v := 10; v > 5 {
    fmt.Println(v)
}
fmt.Println(v)
```
A) Imprime `10` dos veces.
B) Imprime `10` y después `0`.
C) No compila: `v` no existe fuera del `if`.
D) Compila pero `v` vale `nil` en el segundo `Println`.

**5.**
```go
x := 2
switch x {
case 1:
    fmt.Println("uno")
case 2:
    fmt.Println("dos")
case 3:
    fmt.Println("tres")
default:
    fmt.Println("otro")
}
```
A) Imprime `dos`, `tres` y `otro` (fallthrough automático).
B) Imprime solo `dos`.
C) Da error de compilación por falta de `break`.
D) Imprime `otro` porque `x` no matchea ningún case literal.

**6.** ¿Cuál es la diferencia entre `fmt.Print` y `fmt.Println`?
A) Son exactamente iguales.
B) `Print` siempre separa con espacio y agrega salto de línea; `Println` no.
C) `Println` siempre separa con espacio entre todos los argumentos y agrega salto de línea; `Print` solo separa con espacio si ninguno de los dos argumentos vecinos es un string.
D) `Print` no acepta más de un argumento.

**7.**
```go
var x, y string
n, e := fmt.Sscanln("700\n800", &x, &y)
```
A) `n=2`, `e=nil`, `x="700"`, `y="800"`.
B) `n=1`, error `"unexpected newline"`, y `y` queda con el valor que tenía antes (no se pisa).
C) Panic en tiempo de ejecución.
D) No compila.

**8.** ¿Cuál es la diferencia central entre un array y un slice en Go?
A) Son lo mismo, solo cambia la sintaxis de declaración.
B) El array tiene tamaño fijo (parte de su tipo); el slice referencia un array de atrás y puede crecer con `append`.
C) El slice siempre es más lento que el array.
D) El array puede tener índices negativos, el slice no.

**9.**
```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]
b[0] = 99
b = append(b, 100)
fmt.Println(a)
fmt.Println(b)
```
A) `[1 2 3 4 5]` y `[99 3 100]` (independientes).
B) `[1 99 3 100 5]` y `[99 3 100]`.
C) `[1 99 3 4 5]` y `[99 3]`.
D) Panic: index out of range.

**10.** ¿Qué pasa si intentás escribir en un `map` que vale `nil` (nunca se le hizo `make`)?
A) Se crea automáticamente vacío y funciona.
B) No compila.
C) Panic en tiempo de ejecución.
D) Se ignora la escritura silenciosamente.

**11.** ¿Cuáles de estos tipos NO pueden ser clave (`key`) de un `map`?
A) `string` e `int`.
B) `bool` y arrays.
C) Slices, maps y funciones (no son comparables con `==`).
D) Cualquier `struct`.

**12.**
```go
type Contador struct{ valor int }
func (c *Contador) Incrementar() { c.valor++ }
func main() {
    c := Contador{}
    c.Incrementar()
    c.Incrementar()
    fmt.Println(c.valor)
}
```
A) `0`, porque el receiver puntero no modifica el original.
B) `2`.
C) No compila: `c` no es un puntero.
D) `1`.

**13.**
```go
type Animal interface{ Sonido() string }
type Perro struct{}
func (p *Perro) Sonido() string { return "Guau" }
func main() {
    var a Animal = Perro{}
    fmt.Println(a.Sonido())
}
```
A) Imprime `Guau`.
B) No compila: `Sonido()` tiene receiver `*Perro`, así que solo `*Perro` implementa `Animal`, no `Perro`.
C) Panic en tiempo de ejecución.
D) Imprime un string vacío.

**14.**
```go
type T struct{ S string }
func (t *T) M() {
    if t == nil {
        fmt.Println("nil!")
        return
    }
    fmt.Println(t.S)
}
type I interface{ M() }
func main() {
    var i I
    var t *T
    i = t
    fmt.Println(i == nil)
    i.M()
}
```
A) `true` y después `nil!`.
B) `false` y después `nil!` — la interfaz tiene tipo `*T` asignado (aunque el valor sea nil), por eso `i == nil` da `false`.
C) `false` y después panic, porque `t` es nil.
D) No compila.

**15.** ¿Diferencia entre type assertion (`i.(T)`) y type switch (`switch v := i.(type)`)?
A) Son lo mismo, distinta sintaxis nomás.
B) La type assertion solo funciona con `interface{}`, el type switch con cualquier interfaz.
C) La assertion chequea/extrae un único tipo concreto (y puede hacer panic si no usás la forma con `, ok`); el type switch compara contra varios tipos posibles a la vez.
D) El type switch no puede tener `default`.

**16.**
```go
func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("fin")
}
```
A) `0 1 2 fin`
B) `fin 0 1 2`
C) `fin 2 1 0`
D) `2 1 0 fin`

**17.** ¿Dónde funciona `recover()`?
A) En cualquier parte del código.
B) Solo dentro de una función invocada con `defer`; en cualquier otro lado no hace nada y devuelve `nil`.
C) Solo en la función `main`.
D) Solo si se llama antes de que ocurra el `panic`.

**18.**
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
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
}
```
A) `1 1 1` (cada llamada resetea `x`).
B) `1 4 9` (la función devuelta comparte la misma `x` capturada entre llamadas).
C) No compila.
D) `1 2 3`.

**19.** Tenés `values := []int{1,2,3,4}` y una función variádica `func sum(vals ...int) int`. ¿Cómo la llamás pasándole el slice completo?
A) `sum(values)`
B) `sum(values...)`
C) `sum(*values)`
D) No se puede, hay que iterar a mano.

**20.** En `func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V`, ¿por qué se exige `comparable` para `K`?
A) Por convención, no tiene un motivo técnico.
B) Porque `K` es la clave de un `map`, y las claves de un map siempre tienen que poder compararse con `==`.
C) Porque todos los parámetros de tipo genéricos son `comparable` por defecto.
D) Para que `V` pueda sumarse.

**21.** ¿Hace falta escribir siempre `SumIntsOrFloats[string, int64](ints)` con los tipos explícitos entre corchetes?
A) Sí, siempre.
B) No, Go casi siempre puede inferirlos mirando los argumentos (`SumIntsOrFloats(ints)`).
C) Solo si `K` es `string`.
D) Solo en la primera llamada del programa.

**22.**
```go
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println(i)
    }()
}
wg.Wait()
```
(Go 1.22+, que es lo que usa la cátedra). ¿Qué imprime?
A) `3 3 3` (las tres goroutines comparten la misma `i`).
B) Los valores `0`, `1` y `2`, cada uno exactamente una vez, en algún orden (desde Go 1.22 cada iteración del `for` tiene su propia variable `i`).
C) No compila porque `i` no se pasa como parámetro.
D) Da panic por acceso concurrente a `i`.

**23.**
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
A) Funciona bien, es el patrón correcto.
B) Deadlock: `Withdraw` ya tiene el lock, y `Deposit` intenta tomarlo de nuevo — un `sync.Mutex` no es reentrante.
C) Da un error de compilación.
D) Funciona pero es ineficiente.

**24.** ¿Cuáles son las 4 condiciones necesarias para que exista un deadlock?
A) Exclusión mutua, retención y espera, no apropiación, espera circular.
B) Starvation, livelock, race condition, y exclusión mutua.
C) Solo hace falta la espera circular.
D) Buffer lleno, canal cerrado, mutex bloqueado, y goroutine huérfana.

**25.** En el barbero durmiente, el cliente usa un `select` con `default` anidado dos veces. ¿Para qué sirve el `default` ahí?
A) Para manejar errores de compilación.
B) Para que el `select` sea no bloqueante: si esa alternativa no está lista en el instante, prueba la siguiente en vez de quedarse esperando.
C) Para cerrar el canal automáticamente.
D) No cambia nada, es opcional por estilo.

**26.** ¿Qué archivo crea `go mod init miproyecto`, y para qué sirve?
A) `main.go`, para tener un punto de entrada.
B) `go.sum`, para verificar checksums de dependencias.
C) `go.mod`, que define el nombre del módulo (ruta base) para resolver imports propios y manejar dependencias externas.
D) `Makefile`, para automatizar el build.

**27.**
```go
s := "café"
fmt.Println(len(s))
fmt.Println(len([]rune(s)))
```
A) `4` y `4` (son iguales siempre).
B) `5` y `4` — `len(s)` cuenta bytes UTF-8 (la `é` ocupa 2), `len([]rune(s))` cuenta caracteres reales.
C) `4` y `5`.
D) No compila.

**28.**
```go
s := "café"
for i := 0; i < len(s); i++ {
    fmt.Printf("%c", s[i])
}
```
A) Imprime `café` correctamente.
B) Imprime `cafÃ©` (o similar) — está indexando bytes, no caracteres, y la `é` ocupa 2 bytes que se muestran partidos.
C) Panic: index out of range.
D) No compila: no se puede indexar un string con `[]`.

**29.**
```go
var r rune = 'a'
fmt.Println(r)
```
A) Imprime `a`.
B) Imprime `97` — un rune literal es de tipo `rune` (`int32`), y `Println` sin `%c` muestra el código numérico, no el carácter.
C) No compila: `'a'` es un string, no un rune.
D) Imprime `'a'` con las comillas incluidas.

**30.**
```go
for i := 1; i <= 3; i++ {
    fmt.Print(i)
}
```
A) `1 2 3` (con espacios, porque son argumentos numéricos).
B) `123` — son tres llamados a `Print` separados, cada uno con un solo argumento; no hay nada "entre" que espaciar dentro de cada llamado, y `Print` no agrega nada por su cuenta entre llamados distintos.
C) `1\n2\n3\n` (cada uno en su línea).
D) No compila.
