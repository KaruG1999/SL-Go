# Ejercicio Obligatorio 1 — Números primos paralelos

## Enunciado

Realizar un programa que acepte un número entero positivo N desde la línea de comandos y calcule todos los números primos menores o iguales a N.

**a)** Con una única goroutine que muestre la lista de primos encontrados.

**b)** Usando múltiples goroutines para dividir el trabajo:
- i) Cada goroutine recibe un rango de números a comprobar y devuelve la lista de primos de ese rango.
- ii) El main divide el trabajo en goroutines, asignando un rango a cada una.
- iii) Una vez que todas terminan, el main recopila y muestra todos los primos.

**c)** Ejecutar con N = 1.000, 100.000 y 1.000.000 para a) y b). Calcular el **speed-up**:

```
S(p) = T(1) / T(p)
```

donde T(1) es el tiempo con una goroutine y T(p) con p goroutines.

*Objetivo: goroutines, paralelismo, WaitGroups, channels*

---

## Lógica de resolución

### Función de chequeo de primo

```go
func esPrimo(n int) bool {
    if n < 2 { return false }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 { return false }
    }
    return true
}
```

### a) Una sola goroutine

```go
// leer N desde os.Args[1]
N, _ := strconv.Atoi(os.Args[1])
for i := 2; i <= N; i++ {
    if esPrimo(i) { fmt.Println(i) }
}
```

Medir con `time.Now()` antes y después.

### b) Múltiples goroutines con WaitGroup + canal de resultados

```go
numGoroutines := runtime.NumCPU()
ch := make(chan []int, numGoroutines)
var wg sync.WaitGroup

tamRango := N / numGoroutines
for i := 0; i < numGoroutines; i++ {
    inicio := i*tamRango + 1
    fin := inicio + tamRango - 1
    if i == numGoroutines-1 { fin = N }
    wg.Add(1)
    go func(inicio, fin int) {
        defer wg.Done()
        var primos []int
        for n := inicio; n <= fin; n++ {
            if esPrimo(n) { primos = append(primos, n) }
        }
        ch <- primos
    }(inicio, fin)
}

go func() { wg.Wait(); close(ch) }()

var todos []int
for parcial := range ch {
    todos = append(todos, parcial...)
}
sort.Ints(todos)
```

### c) Medir speed-up

```go
start := time.Now()
// ... ejecutar versión a) o b) ...
elapsed := time.Since(start)
fmt.Printf("T: %v\n", elapsed)
```

Calcular `S(p) = T(1) / T(p)` con los tiempos medidos.
