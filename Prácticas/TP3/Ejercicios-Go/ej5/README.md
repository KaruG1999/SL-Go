# Ejercicio 5 — PING PONG con channels

## Enunciado

Garantizar que el siguiente programa imprima exactamente:

```
PING
PONG
PING
PONG
PING
PONG
PING
PONG
```

El código base lanza goroutines que envían "PING" y "PONG" a un canal sin orden garantizado:

```go
func pxng(m chan string, str string) {
    m <- str
}

func main() {
    messages := make(chan string)
    for i := 0; i < 5; i++ {
        go pxng(messages, "PING")
        go pxng(messages, "PONG")
    }
    for i := 0; i < 10; i++ {
        fmt.Println(<-messages)
    }
}
```

*Objetivo: channels, sincronización*

---

## Lógica de resolución

### Problema con el código base

Las goroutines de PING y PONG compiten por el canal. El orden de llegada es no determinista.

### Solución: dos canales alternados

Usar dos canales separados (`ping` y `pong`) y hacer que cada goroutine espere la señal del otro antes de enviar:

```go
func ping(pingCh chan struct{}, pongCh chan struct{}) {
    for i := 0; i < 4; i++ {
        <-pingCh             // espera la ficha para jugar
        fmt.Println("PING")
        pongCh <- struct{}{} // le pasa la ficha a PONG
    }
}

func pong(pingCh chan struct{}, pongCh chan struct{}, done chan bool) {
    for i := 0; i < 4; i++ {
        <-pongCh             // espera la ficha para jugar
        fmt.Println("PONG")
        pingCh <- struct{}{} // le devuelve la ficha a PING
    }
    done <- true // avisa al main que el juego terminó
}

func main() {
    pingCh := make(chan struct{}, 1)
    pongCh := make(chan struct{})
    done := make(chan bool)

    pingCh <- struct{}{} // se mete la primera ficha para que arranque PING

    go ping(pingCh, pongCh)
    go pong(pingCh, pongCh, done)

    <-done // el main espera acá hasta que pong avise que terminó
}
```

> El canal `pingCh` con buffer de 1 permite que PING arranque sin bloquearse. Cada goroutine espera su señal antes de imprimir, garantizando la alternancia.

---

## Conceptos de Teoría

**Canal sin buffer (`make(chan T)`):** la operación de envío bloquea hasta que haya un receptor listo, y viceversa. Sincronización punto a punto. No tiene orden garantizado cuando hay múltiples goroutines compitiendo.

**Canal con buffer (`make(chan T, n)`):** permite hasta `n` envíos sin receptor listo. Usado acá con capacidad 1 para "inyectar" el primer token y que PING arranque sin bloquearse.

**Canal como token/semáforo:** un valor en un canal representa "tu turno". El receptor espera el token, actúa, y lo pasa al siguiente. Es el patrón estándar en Go para garantizar orden entre goroutines sin mutexes.

**No determinismo en canales:** cuando varias goroutines envían al mismo canal, Go no garantiza el orden de recepción. Por eso el código base del enunciado no puede garantizar PING/PONG alternado — se necesitan dos canales separados.
