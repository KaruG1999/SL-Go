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

func (a *Agenda) BuscarContacto(correo string) (Contact, bool) {
    a.mu.RLock()
    defer a.mu.RUnlock()
    c, ok := a.contactos[correo]
    return c, ok
}
```

### Main con goroutines simultáneas

```go
func main() {
    agenda := NewAgenda()
    var wg sync.WaitGroup

    // lanzar goroutines de agregar
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            agenda.AgregarContacto(Contact{
                Nombre: fmt.Sprintf("Nombre%d", i),
                CorreoElectronico: fmt.Sprintf("correo%d@mail.com", i),
            })
        }(i)
    }

    wg.Wait()
    // mostrar contenido final...
}
```

> Usar `RWMutex` en lugar de `Mutex` permite que múltiples `BuscarContacto` corran en paralelo, bloqueando solo para escrituras.
