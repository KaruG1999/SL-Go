# Seminario de Lenguajes — Go — Clase 7: Concurrencia

---

## Contexto de Conexión

Hasta ahora todos los programas ejecutaban una sola línea de instrucciones. En esta clase aprendemos a ejecutar múltiples tareas **al mismo tiempo** usando goroutines y a coordinarlas con WaitGroups, Channels y Select. También vemos cómo proteger datos compartidos con monitores y semáforos.

---

## Conceptos Core

- **Goroutine**: función que se ejecuta concurrentemente con otras. Se lanza con `go f(args)`. Es mucho más liviana que un thread del sistema operativo.
- **`sync.WaitGroup`**: contador que permite esperar a que un conjunto de goroutines termine. Métodos: `Add(n)`, `Done()`, `Wait()`.
- **Channel**: conducto tipado para comunicar goroutines. Se crea con `make(chan T)`. Envío bloqueante (`ch <- v`) y recepción bloqueante (`v := <-ch`).
- **Buffered channel**: channel con capacidad interna (`make(chan T, n)`). El envío solo bloquea si el buffer está lleno; la recepción solo bloquea si está vacío.
- **Channel unidireccional**: `chan<- T` (solo envío) y `<-chan T` (solo recepción). Se usan en firmas de funciones para documentar la dirección de uso.
- **`close(ch)`**: cierra un channel. El receptor puede detectarlo con `v, ok := <-ch` (ok=false) o con `range`.
- **`select`**: permite que una goroutine espere en múltiples operaciones de channel a la vez. Elige al azar si varias están listas. Con `default` no bloquea.
- **Monitor**: goroutine que centraliza el acceso a un recurso compartido usando `select`.
- **`sync.Mutex`**: semáforo binario. `Lock()` y `Unlock()` protegen una sección crítica.
- **`sync.RWMutex`**: permite múltiples lectores simultáneos pero un solo escritor. Métodos: `Lock/Unlock` (escritura) y `RLock/RUnlock` (lectura).

---

## Desarrollo

### 1. Goroutines

Una goroutine se lanza con la keyword `go`:

```go
func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
    }
}

func main() {
    go f(0)       // se ejecuta concurrentemente
    fmt.Scanln()  // espera una tecla para no salir antes de que termine
}
```

Con múltiples goroutines el orden de salida es **no determinístico** — el scheduler decide cuándo ejecuta cada una:

```go
for i := 0; i < 10; i++ {
    go f(i)
}
fmt.Scanln()
```

---

### 2. WaitGroup

`fmt.Scanln()` es un parche; la forma correcta de esperar goroutines es con `sync.WaitGroup`:

```go
var wg sync.WaitGroup

for _, url := range urls {
    wg.Add(1)           // incrementa el contador antes de lanzar
    go responseSize(url)
}
wg.Wait()               // bloquea hasta que el contador llegue a 0
```

La goroutine llama a `wg.Done()` cuando termina (usualmente con `defer`):

```go
func responseSize(url string) {
    defer wg.Done()
    // ... lógica ...
}
```

**Patrón con función anónima** (encierra la llamada a `Done` junto a la lógica):

```go
for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        responseSize(u)
    }(url)    // ← pasar url como argumento para evitar el bug de closure
}
wg.Wait()
```

---

### 3. Channels

**Declaración y operaciones básicas:**

```go
naturals := make(chan int)
squares  := make(chan int)
```

```go
naturals <- x    // envío (bloquea hasta que alguien reciba)
x := <-naturals  // recepción (bloquea hasta que alguien envíe)
```

**Pipeline clásico — tres goroutines en cadena:**

```go
func counter(out chan<- int) {
    for x := 0; x < 10; x++ {
        out <- x
    }
    close(out)
}

func squarer(in <-chan int, out chan<- int) {
    for x := range in {   // range termina cuando el channel se cierra
        out <- x * x
    }
    close(out)
}

func printer(in <-chan int) {
    for x := range in {
        fmt.Println(x)
    }
}

func main() {
    naturals := make(chan int)
    squares  := make(chan int)
    go counter(naturals)
    go squarer(naturals, squares)
    printer(squares)
}
```

**Detectar cierre del channel:**

```go
x, ok := <-nums   // ok=false cuando el channel está cerrado y vacío
```

---

### 4. Buffered channels

```go
ch := make(chan string, 3)  // capacidad 3

ch <- "A"
ch <- "B"
ch <- "C"
// ch <- "D"  ← bloquearía porque el buffer está lleno

fmt.Println(<-ch)    // "A"
fmt.Println(cap(ch)) // 3
fmt.Println(len(ch)) // 2
```

**Caso de uso — Productor/Consumidor:**

```go
func Producer(out chan<- int) {
    for i := 0; i < 10; i++ {
        time.Sleep(...)
        out <- rand.Intn(1000)
    }
}

func Consumer(in <-chan int) {
    for i := range in {
        time.Sleep(...)
        fmt.Println("consumed:", i)
    }
}

func main() {
    ch := make(chan int, 5)
    var wg sync.WaitGroup
    wg.Add(1)
    go func() { Producer(ch); close(ch) }()
    go func() { Consumer(ch); wg.Done() }()
    wg.Wait()
}
```

**Mirrored request** — quedarse con la respuesta más rápida:

```go
func mirroredQuery() string {
    responses := make(chan string, 3)
    go func() { responses <- request("asia.google.com") }()
    go func() { responses <- request("europe.google.com") }()
    go func() { responses <- request("americas.google.com") }()
    return <-responses   // primera respuesta que llegue; las otras se descartan
}
```

> El buffered channel con capacidad 3 evita que las dos goroutines "lentas" queden bloqueadas para siempre al intentar enviar cuando ya nadie escucha.

---

### 5. Select

`select` espera en varias operaciones de channel a la vez. Cuando una está lista, la ejecuta. Si varias están listas, elige una al azar:

```go
for ok1 && ok2 {
    select {
    case val, ok1 = <-ch1:
        if ok1 { fmt.Println("ch1:", val) }
    case val, ok2 = <-ch2:
        if ok2 { fmt.Println("ch2:", val) }
    }
}
```

**Select con `default` (no bloqueante):**

```go
select {
case val := <-ch1:
    // hay dato en ch1
case val := <-ch2:
    // hay dato en ch2
default:
    // ningún channel listo, continuar sin bloquear
}
```

---

### 6. Exclusión mutua — Monitor con channels

Un **monitor** es una goroutine que serializa el acceso a un recurso mediante `select`:

```go
var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
    var balance int
    for {
        select {
        case amount := <-deposits:
            balance += amount
        case balances <- balance:
        }
    }
}

func init() { go teller() }
```

> Solo `teller` accede a `balance`. No hace falta ningún mutex porque solo hay una goroutine que lee/escribe la variable.

---

### 7. Exclusión mutua — sync.Mutex

`sync.Mutex` es un semáforo binario: garantiza que solo una goroutine a la vez ejecute la sección crítica.

```go
var (
    mu      sync.Mutex
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    balance += amount
}

func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}
```

**Evitar deadlock con funciones anidadas:** si `Deposit` llama a `Balance` y ambas intentan adquirir el mismo mutex, el programa se cuelga. Solución: separar la lógica interna en funciones sin lock:

```go
func deposit(amount int) { balance += amount }  // función interna, sin lock

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

---

### 8. Exclusión mutua — sync.RWMutex

Cuando hay muchas lecturas y pocas escrituras, `RWMutex` permite que múltiples lectores accedan simultáneamente:

```go
var mu sync.RWMutex

func Balance() int {
    mu.RLock()          // múltiples goroutines pueden RLock al mismo tiempo
    defer mu.RUnlock()
    return balance
}

func Deposit(amount int) {
    mu.Lock()           // excluye a todos (lectores y escritores)
    defer mu.Unlock()
    balance += amount
}
```

---

## Lo que no podés ignorar

> 1. **`wg.Add(1)` va ANTES de lanzar la goroutine**: si lo ponés dentro de la goroutine, `Wait()` puede terminar antes de que el contador se incremente.
> 2. **No cerrar un channel desde el receptor**: solo el emisor sabe cuándo terminó de enviar. Cerrar un channel ya cerrado produce panic.
> 3. **`range` sobre un channel bloquea hasta que se cierre**: si nadie cierra el channel, el `for range` nunca termina.
> 4. **Buffered channels no eliminan la necesidad de sincronización**: solo dan un margen temporal. Si el buffer se llena, el emisor igual bloquea.
> 5. **`select` con `default` nunca bloquea**: úsalo cuando querés intentar una operación de channel sin esperar.
> 6. **Un mutex no puede adquirirse dos veces por la misma goroutine** (Go no tiene mutexes reentrantes): llamar a `Lock()` cuando ya tenés el lock causa deadlock. Separar funciones internas (sin lock) de funciones públicas (con lock) es el patrón estándar.
> 7. **`sync.RWMutex` solo conviene cuando las lecturas son mucho más frecuentes que las escrituras**: tiene más overhead que `Mutex` cuando hay pocas lecturas.
