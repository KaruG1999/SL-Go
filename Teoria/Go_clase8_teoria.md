# Seminario de Lenguajes — Go — Clase 8: Problemas Clásicos de Concurrencia

---

## Contexto de Conexión

En la clase anterior aprendimos las herramientas de concurrencia de Go: goroutines, channels, select y mutex. Esta clase las aplicamos a los dos problemas clásicos de sincronización: los **Filósofos Comensales** (Dining Philosophers) y el **Barbero Durmiente** (Sleeping Barber). Cada uno expone patrones de deadlock, starvation y coordinación que aparecen en sistemas reales.

---

## Conceptos Core

- **Deadlock**: situación donde un conjunto de goroutines se bloquean mutuamente esperando recursos que las otras retienen, y ninguna puede avanzar.
- **Condiciones de Coffman** (necesarias para deadlock): exclusión mutua + retención y espera + no apropiación + espera circular.
- **Starvation (inanición)**: una goroutine puede progresar pero sistemáticamente no consigue los recursos porque otros siempre se adelantan.
- **Dining Philosophers**: 5 filósofos comparten 5 tenedores; para comer necesitan los dos tenedores adyacentes. El riesgo es deadlock si todos toman el tenedor izquierdo al mismo tiempo.
- **Sleeping Barber**: un barbero duerme si no hay clientes. Cada cliente que llega despierta al barbero o se sienta a esperar; si no hay sillas libres, se va.

---

## Desarrollo

### 1. Dining Philosophers — versión básica (con deadlock)

```go
var philos = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}
var forks   = [5]sync.Mutex{}

func philosopher(id int, forkL, forkR *sync.Mutex) {
    for i := 0; i < 50; i++ {
        // pensar ...
        forkL.Lock()
        forkR.Lock()
        // comer ...
        forkR.Unlock()
        forkL.Unlock()
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

**Problema:** si los 5 filósofos toman su tenedor izquierdo al mismo tiempo, ninguno puede tomar el derecho → **deadlock**.

Las 4 condiciones de Coffman se cumplen:
1. **Exclusión mutua** — cada tenedor es un Mutex.
2. **Retención y espera** — cada filósofo retiene el tenedor izquierdo mientras espera el derecho.
3. **No apropiación** — nadie le quita el tenedor a otro.
4. **Espera circular** — filósofo 0 espera al 1, 1 al 2, ..., 4 al 0.

---

### 2. Dining Philosophers — solución sin deadlock (romper espera circular)

Basta con que **un** filósofo tome los tenedores en orden inverso. La solución clásica: el filósofo de índice par toma primero el izquierdo, el de índice impar toma primero el derecho.

```go
func main() {
    dining.Add(5)
    for i := range philos {
        go philosopher(i, &forks[(i+i%2)%5], &forks[(i+1-i%2)%5])
    }
    dining.Wait()
}
```

> Rompe la **espera circular**: ya no todos esperan al siguiente. Al menos un filósofo puede siempre obtener ambos tenedores.

---

### 3. Dining Philosophers — starvation y solución con penalización

Evitar deadlock no garantiza equidad. Un filósofo puede comer 50 veces mientras otro casi nunca come (starvation).

**Solución:** introducir un mecanismo de penalización — si un filósofo come mucho más que el máximo de sus vecinos, se lo hace esperar:

```go
var count   = [5]int{}
var penalty = [5]int{}

func philosopher(id int, forkL, forkR *sync.Mutex) {
    for i := 0; i < 50; i++ {
        time.Sleep(time.Duration(500 * penalty[id]))  // espera proporcional a la penalización
        forkL.Lock(); forkR.Lock()
        forkR.Unlock(); forkL.Unlock()
        report(id)
    }
    dining.Done()
}

func report(id int) {
    var mu sync.RWMutex
    count[id]++
    mu.RLock()
    if float64(count[id]) > float64(maxCount(count, id))*1.1 {
        penalty[id]++   // comió más de un 10% que el máximo vecino → penalizar
    } else if penalty[id] > 0 {
        penalty[id]--
    }
    mu.RUnlock()
}
```

---

### 4. Sleeping Barber

**El problema:**
- Hay una sala de corte (1 silla de barbero) y una sala de espera (N sillas).
- Si el barbero no tiene clientes, duerme.
- Un cliente que llega: despierta al barbero si duerme, se sienta en sala de espera si hay lugar, o se va si no hay sillas.

**Implementación con channels:**

```go
const n = 5   // sillas de sala de espera

func main() {
    sillas   := make(chan string, n)   // buffered: sala de espera
    despertar := make(chan string)     // para despertar al barbero
    listo     := make(chan bool)       // señal de corte terminado

    go barbero(sillas, listo, despertar)
    // ... lanzar goroutines de clientes ...
}
```

**Barbero:**
```go
func barbero(sillas <-chan string, listo chan<- bool, despertar <-chan string) {
    for {
        // Duerme hasta que alguien lo despierte
        nombre := <-despertar
        cortar(nombre, listo, sillas)

        // Atiende a los que están esperando
        for len(sillas) > 0 {
            nombre := <-sillas
            cortar(nombre, listo, sillas)
        }
    }
}
```

**Cliente:**
```go
func cliente(nombre string, sillas chan string, listo <-chan bool, despertar chan<- string) {
    select {
    case despertar <- nombre:
        // El barbero estaba durmiendo, lo despertamos
    default:
        select {
        case sillas <- nombre:
            // Hay silla libre, esperamos
            <-listo
        default:
            // No hay sillas, el cliente se va
            fmt.Println(nombre, "no hay lugar en sala de espera")
        }
    }
}

func cortar(nombre string, listo chan<- bool, sillas <-chan string) {
    time.Sleep(time.Duration(100 + rand.Intn(100)))
    listo <- true
}
```

**Coordinación:**
- `despertar` (unbuffered): si el barbero no está dormido (ya recibió), el send no bloqueará → el `select` con `default` en el cliente no envía por ese canal.
- `sillas` (buffered con cap N): actúa como cola de espera. Si está llena, el cliente se va.
- `listo` (unbuffered): el cliente bloquea hasta que el barbero termina el corte.

---

## Lo que no podés ignorar

> 1. **Deadlock requiere las 4 condiciones simultáneamente**: romper cualquiera de las 4 lo evita. La más fácil suele ser romper la **espera circular** (ordenar los recursos).
> 2. **Evitar deadlock no es suficiente**: podés tener un sistema sin deadlock pero con starvation, donde algunas goroutines progresan y otras no. Son problemas distintos.
> 3. **El orden de adquisición de locks importa**: si todos los filósofos toman los tenedores en el mismo orden (ej. siempre el de menor índice primero), se elimina la espera circular.
> 4. **Channels buffered como salas de espera**: un `make(chan T, N)` modela naturalmente una sala de espera de capacidad N. Si está lleno, el `select default` permite que el cliente "se vaya".
> 5. **`select` con `default` como intento no bloqueante**: es el patrón clave del Barbero Durmiente — el cliente intenta despertar al barbero; si no puede (porque ya está despierto), intenta sentarse en la sala de espera.
> 6. **No acceder a variables compartidas sin protección**: `count` y `penalty` del ejemplo de starvation se acceden desde múltiples goroutines — requieren RWMutex o channels para ser seguros.
