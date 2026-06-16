# Ejercicio 9 — Select con timeouts

## Enunciado

Realizar un programa que envíe datos a **dos canales desde dos goroutines** y los reciba en el programa principal con `select`. Los datos deben recibirse:

- En uno de los canales por el lapso de **5 segundos**.
- En el otro canal por el lapso de **10 segundos**.

*Objetivo: timeouts*

---

## Lógica de resolución

### Patrón de timeout con `time.After`

`time.After(d)` devuelve un canal que recibe un valor después de la duración `d`. Se usa como caso en `select` para implementar un timeout:

```go
select {
case v := <-ch:
    fmt.Println("recibido:", v)
case <-time.After(5 * time.Second):
    fmt.Println("timeout en ch1")
}
```

### Estructura completa

```go
ch1 := make(chan int)
ch2 := make(chan int)

// goroutine 1: envía cada 1 segundo
go func() {
    for i := 0; ; i++ {
        time.Sleep(time.Second)
        ch1 <- i
    }
}()

// goroutine 2: envía cada 2 segundos
go func() {
    for i := 0; ; i++ {
        time.Sleep(2 * time.Second)
        ch2 <- i
    }
}()

timeout1 := time.After(5 * time.Second)
timeout2 := time.After(10 * time.Second)

for {
    select {
    case v := <-ch1:
        fmt.Println("ch1:", v)
    case v := <-ch2:
        fmt.Println("ch2:", v)
    case <-timeout1:
        fmt.Println("timeout ch1 (5s)")
        ch1 = nil // dejar de escuchar ch1
    case <-timeout2:
        fmt.Println("timeout ch2 (10s)")
        return // terminar el programa
    }
}
```

> `timeout1` y `timeout2` se crean una sola vez fuera del loop. Si se pusieran dentro del loop, se reiniciarían en cada iteración y nunca dispararían.

---

## Conceptos de Teoría

**`time.After(d)`:** devuelve un `<-chan time.Time` que envía un valor único luego de la duración `d`. Se usa como caso en `select` para implementar timeouts sin goroutines adicionales.

**Timeout con `select`:** el caso `<-time.After(d)` compite con los demás casos del select. Si ningún canal tiene datos en `d` tiempo, el timeout gana y se puede reaccionar (cancelar, loggear, salir).

**Creación del timer fuera del loop:** `time.After` crea un canal de un solo disparo. Si se llama dentro del loop, cada iteración crea un timer nuevo que empieza desde cero — el timeout nunca llega. Crearlo antes del loop hace que cuente desde el inicio del programa.

**Canal `nil` para dejar de escuchar:** al asignar `nil` a `ch1` después del primer timeout, ese canal deja de participar en el `select`. El programa sigue procesando `ch2` hasta que dispara `timeout2`. Evita tener que reorganizar la lógica del loop.
