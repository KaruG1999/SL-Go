# Ejercicios de práctica — Seminario de Lenguajes opción Go

Ejercicios cortos organizados por tema, mezclando dos formatos típicos de examen: **escribir código** y **predecir qué imprime / si compila** (este segundo formato es el que más se presta a multiple choice). Intentá resolverlos primero; las respuestas están en `soluciones-practica.md`, aparte.

---

## Go-1: Básicos

**1.1 — Escribir código.** Declará tres variables (un `int`, un `float64` y un `string`) usando `:=`, y otra variable `int` usando `var` con tipo explícito. Imprimilas todas con `fmt.Println`.

**1.2 — ¿Compila?**
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

**1.3 — ¿Qué imprime?**
```go
x := 7
y := 2
fmt.Println(x/y, x%y)
```

**1.4 — Escribir código.** Función `EsPar(n int) bool` que devuelva si `n` es par.

---

## Go-2: Control de flujo, funciones, fmt

**2.1 — Escribir código.** Función `FizzBuzz(n int)` que para cada número de 1 a `n` imprima "Fizz" si es múltiplo de 3, "Buzz" si es múltiplo de 5, "FizzBuzz" si es múltiplo de ambos, y el número si no es ninguno.

**2.2 — ¿Qué imprime?**
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

**2.3 — Escribir código.** Función que reciba dos enteros y devuelva dos valores: la suma, y un booleano que indique si esa suma es par.

**2.4 — ¿Compila?**
```go
if v := 10; v > 5 {
    fmt.Println(v)
}
fmt.Println(v)
```

---

## Go-3: Arrays, Slices, Maps

**3.1 — Escribir código.** Función `Pares(nums []int) []int` que devuelva un nuevo slice solo con los números pares de `nums`, sin modificar el original.

**3.2 — ¿Qué imprime?**
```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]
b[0] = 99
b = append(b, 100)
fmt.Println(a)
fmt.Println(b)
```

**3.3 — Escribir código.** Función `ContarPalabras(texto string) map[string]int` que cuente cuántas veces aparece cada palabra en `texto` (separadas por espacios — se puede usar `strings.Fields`).

**3.4 — ¿Qué imprime?**
```go
m := map[string]int{"a": 1, "b": 2}
v, ok := m["c"]
fmt.Println(v, ok)
```

**3.5 — Escribir código.** Función `Invertir(s []int) []int` que devuelva un nuevo slice con los elementos de `s` en orden inverso, sin modificar `s`.

---

## Go-4: Punteros, Structs, Interfaces

**4.1 — Escribir código.** Struct `Persona` con campos `Nombre string` y `Edad int`, y un método `Saludar()` que imprima `"Hola, soy <Nombre> y tengo <Edad> años"`.

**4.2 — ¿Qué imprime?**
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

**4.3 — ¿Compila? ¿Por qué?**
```go
type Animal interface{ Sonido() string }
type Perro struct{}

func (p *Perro) Sonido() string { return "Guau" }

func main() {
    var a Animal = Perro{}
    fmt.Println(a.Sonido())
}
```

**4.4 — Escribir código.** Función `Sumar1(p *int)` que reciba un puntero a `int` y le sume 1 al valor apuntado (que se vea reflejado afuera de la función).

**4.5 — ¿Compila? ¿Qué imprime?**
```go
type Forma interface {
    Area() float64
}
type Rectangulo struct {
    Base, Altura float64
}
func (r Rectangulo) Area() float64 {
    return r.Base * r.Altura
}
func main() {
    var f Forma
    r := Rectangulo{Base: 4, Altura: 5}
    f = r
    fmt.Println(f.Area())
}
```
¿Cambiaría algo si `Area()` estuviera definido con receiver `*Rectangulo` en vez de `Rectangulo`?

---

## Go-5: Errores, funciones, panic/recover

**5.1 — Escribir código.** Función `Dividir(a, b float64) (float64, error)` que devuelva un error si `b` es 0 (en vez de dividir).

**5.2 — ¿Qué imprime?**
```go
func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("fin")
}
```

**5.3 — Escribir código.** Función `DivisionSegura(a, b int) (resultado int, err error)` que divida dos enteros, pero en vez de dejar que el programa haga panic si `b` es 0, use `recover` (dentro de un `defer`) para capturarlo y devolver un `error` en cambio.

**5.4 — Escribir código.** Función `Contador()` que devuelva una función `func() int`; cada vez que se llama a esa función devuelta, debe incrementar un contador interno y devolver el nuevo valor (1, 2, 3, ...).

---

## Go-6: Genéricos

**6.1 — Escribir código.** Función genérica `Maximo[T cmp.Ordered](a, b T) T` que devuelva el mayor de los dos valores.

**6.2 — Escribir código.** Tipo genérico `Pila[T any]` (stack) con un slice interno y dos métodos: `Push(v T)` (agrega arriba) y `Pop() (T, bool)` (saca y devuelve el de arriba; el booleano indica si había algo para sacar).

**6.3 — Para pensar (sin código).** Si quisieras escribir una función genérica que sume todos los elementos de un slice (`func Sumar[T ???](s []T) T`), ¿qué constraint le pondrías a `T` y por qué no alcanza con `any`?

---

## Go-7: Concurrencia

**7.1 — ¿Qué imprime?**
```go
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 3; i++ {
            ch <- i
        }
        close(ch)
    }()
    for v := range ch {
        fmt.Println(v)
    }
}
```
¿Qué pasaría si sacamos el `close(ch)`?

**7.2 — Escribir código.** Programa con 3 goroutines, cada una imprime su propio id (0, 1, 2), usando `sync.WaitGroup` para que `main` espere a que las 3 terminen antes de salir.

**7.3 — ¿Qué problema tiene este código?**
```go
func main() {
    contador := 0
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            contador++
        }()
    }
    wg.Wait()
    fmt.Println(contador)
}
```
¿Cómo lo arreglarías?

**7.4 — Escribir código.** Un `select` que reciba de un canal `ch`, pero si no llega nada en 2 segundos, imprima `"timeout"` (usando `time.After`).

---

## Go-8: Problemas clásicos de concurrencia

**8.1 — Para pensar.** De las 4 condiciones necesarias para un deadlock (exclusión mutua, retención y espera, no apropiación, espera circular), ¿cuál es la que rompe la solución de "que los filósofos no todos agarren los tenedores en el mismo orden"?

**8.2 — Escribir código.** Escribí un ejemplo mínimo de deadlock con 2 goroutines y 2 mutex (`muA`, `muB`), donde una goroutine toma `muA` y después `muB`, y la otra toma `muB` y después `muA`.

---

## Go-Package: Packages, Modules, Dependencies

**P.1 — Para pensar.** ¿Qué comando usás para inicializar un módulo llamado `miproyecto`? ¿Qué archivo se crea y qué contiene?

**P.2 — Para pensar.** Si tu código importa un paquete externo (de terceros) que todavía no descargaste, ¿qué comando usás para agregarlo como dependencia?
