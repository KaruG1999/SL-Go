# Ejercicio 10 — Agenda de contactos concurrente

## Enunciado

Crear un programa que maneje una lista de contactos de manera concurrente. La agenda debe permitir agregar, eliminar y buscar contactos de forma segura desde múltiples goroutines.

**a)** Definir la estructura `Contact` con campos: `Nombre`, `Apellido`, `CorreoElectronico`, `Telefono`.

**b)** Crear la estructura `Agenda` con un mapa de `Contact` usando el correo como clave.

**c)** Implementar los métodos:
- `AgregarContacto(contacto Contact)`
- `EliminarContacto(correo string)`
- `BuscarContacto(correo string) Contact`

**d)** Asegurar que las operaciones sean seguras para acceso concurrente desde múltiples goroutines.

**e)** Función `main` que cree una agenda, lance varias goroutines para operar simultáneamente y muestre el contenido final.

*Objetivo: concurrencia, sync*

---

## Lógica de resolución

### Estructuras

```go
type Contact struct {
    Nombre           string
    Apellido         string
    CorreoElectronico string
    Telefono         string
}

type Agenda struct {
    mu       sync.RWMutex
    contactos map[string]Contact
}

func NewAgenda() *Agenda {
    return &Agenda{contactos: make(map[string]Contact)}
}
```

### Métodos con protección de concurrencia

```go
func (a *Agenda) AgregarContacto(c Contact) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.contactos[c.CorreoElectronico] = c
}

func (a *Agenda) EliminarContacto(correo string) {
    a.mu.Lock()
    defer a.mu.Unlock()
    delete(a.contactos, correo)
}

func (a *Agenda) BuscarContacto(correo string) Contact {
    a.mu.RLock()
    defer a.mu.RUnlock()
    return a.contactos[correo]
}
```

`BuscarContacto` devuelve el zero value de `Contact` (todos los campos vacíos) si el correo no está — no hay un segundo valor `bool` porque el enunciado pide la firma tal cual `BuscarContacto(correo string) Contact`.

### Main con goroutines simultáneas

```go
func main() {
    agenda := NewAgenda()
    var wg sync.WaitGroup

    // 5 goroutines agregando en simultáneo
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            agenda.AgregarContacto(Contact{
                Nombre:            fmt.Sprintf("Nombre%d", i),
                CorreoElectronico: fmt.Sprintf("correo%d@mail.com", i),
            })
        }(i)
    }
    wg.Wait()

    // eliminar y buscar en simultáneo
    wg.Add(2)
    go func() {
        defer wg.Done()
        agenda.EliminarContacto("correo2@mail.com")
    }()
    go func() {
        defer wg.Done()
        c := agenda.BuscarContacto("correo3@mail.com")
        fmt.Println("Encontrado:", c)
    }()
    wg.Wait()

    // recién acá, con todas las goroutines terminadas, se recorre el mapa directo
    for correo, c := range agenda.contactos {
        fmt.Printf("  %s -> %+v\n", correo, c)
    }
}
```

> Usar `RWMutex` en lugar de `Mutex` permite que múltiples `BuscarContacto` corran en paralelo, bloqueando solo para escrituras.
>
> El `for correo, c := range agenda.contactos` del final accede al mapa directo, sin pasar por el mutex — está bien porque ya se hizo `wg.Wait()` y no queda ninguna goroutine escribiendo. Si se hiciera ese recorrido mientras otras goroutines todavía pueden escribir, sería una race condition real.

## Observación

El enunciado pide lanzar goroutines de agregar, eliminar y buscar "de manera simultánea". El código real las corre en dos tandas: primero las 5 de `AgregarContacto` (con su propio `wg.Wait()`), y recién después `EliminarContacto` y `BuscarContacto` juntas. Cumple igual con "seguro para acceso concurrente" (nada se rompe si se solapan), pero si preguntan puntualmente por las tres operaciones ejecutándose *todas* al mismo tiempo, no es exactamente lo que hace este `main` — se podría lanzar todo en un solo bloque de goroutines sin el `wg.Wait()` intermedio.
