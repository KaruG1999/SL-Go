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

### b) Worker pool con canal de tareas (como está en `main.go`)

```go
const tamLote = 1000
type lote struct{ inicio, fin int }

func primosParalelo(N int, numGoroutines int) []int {
    tareas := make(chan lote, numGoroutines)
    resultados := make(chan []int, numGoroutines)
    var wg sync.WaitGroup

    for w := 0; w < numGoroutines; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for l := range tareas {
                var parciales []int
                for n := l.inicio; n <= l.fin; n++ {
                    if esPrimo(n) { parciales = append(parciales, n) }
                }
                resultados <- parciales
            }
        }()
    }

    go func() {
        for inicio := 2; inicio <= N; inicio += tamLote {
            fin := inicio + tamLote - 1
            if fin > N { fin = N }
            tareas <- lote{inicio, fin}
        }
        close(tareas)
    }()

    go func() { wg.Wait(); close(resultados) }()

    var todos []int
    for parcial := range resultados {
        todos = append(todos, parcial...)
    }
    sort.Ints(todos)
    return todos
}
```

Cada goroutine saca el siguiente lote de 1000 números apenas termina el anterior, en vez de recibir de antemano un rango fijo. Por qué se llegó a esto — ver la sección **De rangos fijos a worker pool** más abajo.

### c) Medir speed-up

```go
start := time.Now()
// ... ejecutar versión a) o b) ...
elapsed := time.Since(start)
fmt.Printf("T: %v\n", elapsed)
```

Calcular `S(p) = T(1) / T(p)` con los tiempos medidos.

---

## De rangos fijos a worker pool: análisis del desbalance de carga

La primera versión de `primosParalelo` repartía `[2, N]` en `numGoroutines` rangos contiguos de igual **tamaño** (`tamRango := N / numGoroutines`, con la última goroutine absorbiendo el resto de la división entera). Funcionalmente daba bien (nunca faltaba ni sobraba un primo), pero el reparto de **trabajo** no era parejo. Midiendo cuánto tardaba cada goroutine por separado para `N=1.000.000` con 8 goroutines:

```
goroutine 0 [1,       125000]: 135ms
goroutine 1 [125001,  250000]: 210ms
goroutine 2 [250001,  375000]: 265ms
goroutine 3 [375001,  500000]: 309ms
goroutine 4 [500001,  625000]: 309ms
goroutine 5 [625001,  750000]: 310ms
goroutine 6 [750001,  875000]: 310ms
goroutine 7 [875001, 1000000]: 332ms
```

Mismo tamaño de rango (125.000 números cada una), pero la última tarda **más del doble** que la primera. La razón: `esPrimo` prueba divisores hasta `√n`, así que verificar un número cercano a 1.000.000 cuesta bastante más que uno cercano a 1. Como el tiempo total lo marca la goroutine más lenta, esto explica gran parte de por qué el speed-up no se acerca a 8x con 8 núcleos.

### Primer intento de arreglo (que salió peor): intercalar por índice

La idea intuitiva es que cada goroutine revise números salteados (`goroutine id revisa id+2, id+2+p, id+2+2p, ...`) en vez de un bloque contiguo, para que cada una termine con una mezcla de números baratos y caros. Midiendo esa versión con los mismos 8 goroutines:

```
goroutine 0: 7ms      goroutine 1: 598ms
goroutine 2: 6ms      goroutine 3: 611ms
goroutine 4: 6ms      goroutine 5: 593ms
goroutine 6: 6ms      goroutine 7: 592ms
```

Peor todavía. Con `numGoroutines=8` (par), cada residuo `id+2` conserva siempre la misma paridad: las goroutines con `id` par terminan revisando casi solo números pares (que `esPrimo` descarta al instante, `n%2==0` en la primera vuelta), y las de `id` impar revisan casi solo números impares (que sí necesitan recorrer todo el trial division). El intercalado no rompe el patrón de costo — lo alinea con la paridad y lo empeora. Speed-up medido con esta versión: **2.61x**, peor que la de rangos fijos.

### Solución: worker pool con lotes chicos

La idea (a esto se le llama "worker pool"): en vez de decidir de antemano qué número le toca a cada goroutine, se arma una pila de tareas chicas (lotes de 1000 números) en un canal, en el medio de la mesa. Cada goroutine, apenas se desocupa, saca la siguiente tarea de esa pila — nadie les asigna nada, ellas mismas van sacando de a una. Así el reparto se ajusta solo al costo real de cada lote, sin necesidad de conocer de antemano que los números grandes cuestan más ni de evitar correlaciones como la de paridad.

### Comparación de las tres versiones (N=1.000.000, 8 goroutines, esta máquina)

| Versión | Speed-up |
|---|---|
| Rangos contiguos fijos | ~4.5x |
| Intercalado por índice (id, id+p, id+2p, ...) | ~2.6x |
| Worker pool (lotes de 1000) | ~3.7x–4.9x según la corrida |

Los tres dan los primos correctos siempre — la diferencia es solo de **rendimiento**. El worker pool no es mágico (el tiempo total varía de corrida a corrida por otros procesos en la máquina), pero es la única de las tres que no depende de ningún supuesto sobre cómo se distribuye el costo entre los números del rango: si el patrón de costo cambiara (por ejemplo, revisando otra propiedad de los números en vez de primalidad), las otras dos versiones podrían desbalancearse de formas distintas e impredecibles, mientras que el worker pool se adapta solo.
