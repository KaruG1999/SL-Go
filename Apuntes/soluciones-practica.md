# Soluciones — Ejercicios de práctica

No mires esto hasta haber intentado resolver `ejercicios-practica.md` primero.

---

## Go-1

**1.1**
```go
n := 42
p := 3.14
s := "hola"
var edad int = 30
fmt.Println(n, p, s, edad)
```

**1.2 — No compila.** `Kilometros` y `Millas` son tipos nombrados distintos (aunque los dos tengan `float64` como tipo subyacente). Go no los suma directamente: `k + m` da error de compilación (`mismatched types`). Habría que convertir uno explícitamente, ej: `float64(k) + float64(m)`.

**1.3** `3 1` (7/2 = 3 en división entera, 7%2 = 1 de resto).

**1.4**
```go
func EsPar(n int) bool {
    return n%2 == 0
}
```

---

## Go-2

**2.1**
```go
func FizzBuzz(n int) {
    for i := 1; i <= n; i++ {
        switch {
        case i%15 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}
```

**2.2** `dos`. Entra al `case 2`, lo imprime, y no hace fallthrough al `case 3` (en Go no es automático).

**2.3**
```go
func SumaEsPar(a, b int) (int, bool) {
    suma := a + b
    return suma, suma%2 == 0
}
```

**2.4 — No compila.** `v` fue declarada en la sentencia de inicialización del `if` (`if v := 10; ...`), así que su alcance es solo ese `if`/`else`. El segundo `fmt.Println(v)` de afuera da error: `undefined: v`.

---

## Go-3

**3.1**
```go
func Pares(nums []int) []int {
    var resultado []int
    for _, n := range nums {
        if n%2 == 0 {
            resultado = append(resultado, n)
        }
    }
    return resultado
}
```

**3.2**
```
[1 99 3 100 5]
[99 3 100]
```
`b := a[1:3]` tiene `len=2, cap=4` (comparte el array de `a` desde el índice 1 hasta el final del array original). `b[0] = 99` modifica `a[1]` directamente (quedan compartiendo memoria). Como `cap(b)` es 4 (hay lugar hasta el final de `a`), `append(b, 100)` **no** crea un array nuevo: escribe en `a[3]` (pisando el `4` que había ahí) y `b` pasa a tener `len=3`. Por eso `a` queda `[1 99 3 100 5]` (se pisó el `4` con el `100`) y `b` queda `[99 3 100]`.

**3.3**
```go
func ContarPalabras(texto string) map[string]int {
    conteo := make(map[string]int)
    for _, palabra := range strings.Fields(texto) {
        conteo[palabra]++
    }
    return conteo
}
```

**3.4** `0 false`. La clave `"c"` no existe en el map, entonces `v` toma el "zero value" del tipo (`0` para `int`), y `ok` es `false`.

**3.5**
```go
func Invertir(s []int) []int {
    resultado := make([]int, len(s))
    for i, v := range s {
        resultado[len(s)-1-i] = v
    }
    return resultado
}
```

---

## Go-4

**4.1**
```go
type Persona struct {
    Nombre string
    Edad   int
}

func (p Persona) Saludar() {
    fmt.Printf("Hola, soy %s y tengo %d años\n", p.Nombre, p.Edad)
}
```

**4.2** `2`. `Incrementar` tiene receiver puntero (`*Contador`), así que sí modifica el `valor` real de `c` (Go llama automáticamente con `&c` aunque `c` no sea un puntero explícito).

**4.3 — No compila.** `Sonido()` está definido sobre `*Perro` (receiver puntero), así que solo `*Perro` implementa la interfaz `Animal`, no `Perro` (valor). La línea `var a Animal = Perro{}` da error de compilación. Para que compile, habría que escribir `var a Animal = &Perro{}`.

**4.4**
```go
func Sumar1(p *int) {
    *p++
}
```

**4.5 — Sí compila, imprime `20`.** `Rectangulo` implementa `Forma` porque `Area()` tiene receiver por valor (`Rectangulo`, no `*Rectangulo`). Si `Area()` estuviera definida con receiver `*Rectangulo`, la línea `f = r` (asignar el valor `r`, sin `&`) **no compilaría**, porque en ese caso solo `*Rectangulo` implementaría `Forma` — habría que escribir `f = &r`.

---

## Go-5

**5.1**
```go
func Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("no se puede dividir por cero")
    }
    return a / b, nil
}
```

**5.2**
```
fin
2
1
0
```
Los `defer` se apilan y se ejecutan en orden **inverso** (LIFO) al final de la función, después de que termina el resto del código de `main` (por eso "fin" se imprime primero, y después 2, 1, 0 en ese orden).

**5.3**
```go
func DivisionSegura(a, b int) (resultado int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("error: %v", r)
        }
    }()
    resultado = a / b
    return
}
```
(en Go, dividir un `int` por cero **sí** genera un panic en tiempo de ejecución, a diferencia de los `float64`, donde da `+Inf`).

**5.4**
```go
func Contador() func() int {
    contador := 0
    return func() int {
        contador++
        return contador
    }
}
```

---

## Go-6

**6.1**
```go
func Maximo[T cmp.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

**6.2**
```go
type Pila[T any] struct {
    elementos []T
}

func (p *Pila[T]) Push(v T) {
    p.elementos = append(p.elementos, v)
}

func (p *Pila[T]) Pop() (T, bool) {
    var zero T
    if len(p.elementos) == 0 {
        return zero, false
    }
    ultimo := p.elementos[len(p.elementos)-1]
    p.elementos = p.elementos[:len(p.elementos)-1]
    return ultimo, true
}
```

**6.3** Necesitás un constraint que garantice que `T` soporta el operador `+` (por ejemplo `int64 | float64`, o algo como `constraints.Ordered`/`constraints.Numeric` según la librería). `any` no alcanza porque `any` no garantiza ninguna operación — con `any` ni siquiera podrías escribir `total += v` dentro de la función, porque el compilador no sabe si ese tipo soporta la suma.

---

## Go-7

**7.1**
```
0
1
2
```
Sin `close(ch)`, el `for v := range ch` nunca se entera de que ya no van a llegar más valores — se queda esperando para siempre después del `2` (deadlock: la goroutine que mandaba datos ya terminó su `for`, pero nadie cierra el canal, así que `range` sigue bloqueado esperando el próximo valor que nunca llega).

**7.2**
```go
func main() {
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Println(id)
        }(i)
    }
    wg.Wait()
}
```

**7.3 — Race condition.** `contador++` no es atómico (es leer, sumar 1, escribir), así que si dos goroutines lo hacen "al mismo tiempo" se pueden pisar y perder incrementos — el resultado final puede dar menos de 100, y puede variar entre corridas. Se arregla protegiendo `contador++` con un `sync.Mutex` (`mu.Lock()` / `mu.Unlock()` alrededor del incremento) o usando `sync/atomic`.

**7.4**
```go
select {
case v := <-ch:
    fmt.Println("recibido:", v)
case <-time.After(2 * time.Second):
    fmt.Println("timeout")
}
```

---

## Go-8

**8.1** Rompe la **espera circular**. Al hacer que no todos pidan los tenedores en el mismo orden, ya no se puede formar el círculo cerrado de "cada uno espera al siguiente" que hacía falta para el deadlock.

**8.2**
```go
var muA, muB sync.Mutex

func goroutine1() {
    muA.Lock()
    defer muA.Unlock()
    time.Sleep(time.Millisecond) // da tiempo a que la otra goroutine tome muB
    muB.Lock()
    defer muB.Unlock()
}

func goroutine2() {
    muB.Lock()
    defer muB.Unlock()
    time.Sleep(time.Millisecond)
    muA.Lock()
    defer muA.Unlock()
}

func main() {
    go goroutine1()
    go goroutine2()
    time.Sleep(time.Second)
}
```
Si las dos llegan a tomar su primer mutex antes de que la otra intente tomar el segundo, quedan esperando una a la otra para siempre.

---

## Go-Package

**P.1** `go mod init miproyecto`. Crea el archivo `go.mod`, que contiene el nombre del módulo (`module miproyecto`) y la versión de Go usada (ej: `go 1.20`).

**P.2** `go get <ruta-del-paquete>` (ej: `go get rsc.io/quote`). Descarga el paquete y actualiza `go.mod` (y `go.sum`) con la versión usada.
