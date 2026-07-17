# Ejercicio Obligatorio 2 — Cajas de supermercado

## Enunciado

Simular la atención de clientes en las cajas de un supermercado. La atención de cada cliente se simula con una espera de entre 0 y 1 segundo.

**a)** Una **cola global**: los clientes esperan en una única cola y van a la caja que quede libre.

**b)** **Colas individuales por caja** con asignación **round-robin**: cada cliente nuevo va a la siguiente caja, dando la vuelta cuando llega a la última.

**c)** **Colas individuales por caja** con asignación a la **caja que tenga la cola más corta** en ese momento.

**d)** Imprimir los tiempos de ejecución de las tres versiones y compararlos.

*Objetivo: channels, goroutines, patrones de distribución de trabajo*

---

## Lógica de resolución (como está en `main.go`)

### Duraciones fijadas de antemano

Antes de correr cualquiera de las tres versiones, se generan las duraciones de atención de los 12 clientes una sola vez, en el goroutine principal (sin nada corriendo todavía en paralelo):

```go
func generarDuraciones() []time.Duration {
    duraciones := make([]time.Duration, numClientes)
    for i := range duraciones {
        duraciones[i] = time.Duration(rand.Intn(1000)) * time.Millisecond
    }
    return duraciones
}
```

Esa misma lista (`duraciones`) se le pasa a las tres versiones, así el cliente `i` tarda siempre exactamente lo mismo en a), b) y c) — necesario para que la comparación de tiempos del punto d) tenga sentido (si cada versión generara sus propios tiempos al azar, estaría comparando peras con manzanas).

### a) Cola global (un solo canal compartido)

```go
func colaGlobal(duraciones []time.Duration) time.Duration {
    inicio := time.Now()
    cola := make(chan int, numClientes)
    var wg sync.WaitGroup

    for id := 0; id < numCajas; id++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for cliente := range cola {
                atender(id, cliente, duraciones[cliente])
            }
        }(id)
    }

    for i := 0; i < numClientes; i++ {
        cola <- i
    }
    close(cola)
    wg.Wait()
    return time.Since(inicio)
}
```

Las 3 cajas comparten el mismo canal `cola`. Ninguna caja tiene "su" cliente asignado de antemano: la que esté libre primero saca el siguiente número. Por eso a esto se le llama cola global — nadie decide de antemano quién atiende a quién.

### b) Round-robin (una cola por caja, reparto por turno)

```go
func roundRobin(duraciones []time.Duration) time.Duration {
    ...
    for i := 0; i < numClientes; i++ {
        cajas[i%numCajas] <- i
    }
    ...
}
```

`i % numCajas` reparte los clientes en orden fijo: 0, 1, 2, 0, 1, 2, ... (con `numCajas=3`). No importa si una caja está más ocupada que otra, igual le sigue tocando cliente por turno.

### c) Cola más corta (una cola por caja, reparto dinámico)

```go
func colaMasCorta(duraciones []time.Duration) time.Duration {
    ...
    for i := 0; i < numClientes; i++ {
        minCola := 0
        for j := 1; j < numCajas; j++ {
            if len(cajas[j]) < len(cajas[minCola]) {
                minCola = j
            }
        }
        cajas[minCola] <- i
        time.Sleep(50 * time.Millisecond) // escalona la llegada
    }
    ...
}
```

`len(cajas[j])` en un canal con buffer no es la capacidad del canal, es **cuántos elementos hay esperando ahí adentro en este momento** — o sea, cuántos clientes hacen fila en esa caja todavía sin atender. Se usa eso para elegir siempre la cola más corta en el momento de repartir. El `time.Sleep(50ms)` entre cliente y cliente simula que no llegan todos de golpe, sino de a poco — si llegaran todos a la vez, la "cola más corta" se decidiría siempre con las colas vacías (0 == 0) y la comparación no tendría sentido.

## Sobre `rand.Seed` y la reproducibilidad (un problema que no es solo de concurrencia)

La primera versión de este código llamaba a `rand.Intn` **dentro** de cada goroutine cajero, y reseteaba `rand.Seed(1)` antes de cada una de las tres versiones, con la idea de que así los tres experimentos usaran los mismos tiempos de atención. Se probó corriendo el programa varias veces y **no daba lo mismo**: el cliente 5, por ejemplo, tardaba distinto en cada corrida.

Hay dos razones, una más obvia que la otra:

1. **Concurrencia:** con 3 cajas leyendo y llamando a `rand.Intn` al mismo tiempo, aunque la secuencia de números esté fijada por el seed, no está fijado en qué orden cada goroutine agarra el próximo número de esa secuencia — depende de cómo el sistema operativo decida turnarlas, y eso cambia en cada corrida.
2. **`rand.Seed` ya no asegura nada por sí solo:** se probó llamar `rand.Seed(1)` dos veces seguidas en el mismo programa (sin ninguna goroutine de por medio) y dio dos secuencias de números **distintas**. Desde Go 1.20, las funciones sueltas de `math/rand` (`rand.Intn`, `rand.Seed`, etc.) están pensadas para no ser reproducibles a propósito — `Seed` está deprecado y ya no controla la secuencia de estas funciones como antes.

La solución que se aplicó esquiva ambos problemas de una: en vez de confiar en que el generador aleatorio dé lo mismo dos veces, se generan las 12 duraciones **una sola vez**, en una lista, y esa misma lista se reutiliza en las tres versiones. No importa si `rand.Seed` reproduce algo o no — el cliente `i` siempre tarda lo mismo porque es literalmente el mismo valor guardado, no un número que se vuelve a sortear.

## Observaciones

- `atender(caja, cliente int, dur time.Duration)` ahora recibe la duración como parámetro en vez de decidirla ella misma — así cualquiera de las tres versiones puede usar exactamente la misma lista de tiempos.
- El resultado de a), b) y c) no tiene por qué ser siempre "cola global gana" o "cola más corta gana": con solo 3 cajas y 12 clientes y duraciones al azar, puede haber corridas donde el reparto por round-robin quede mejor repartido de casualidad. Para una conclusión más sólida habría que correr cada versión varias veces y promediar.
