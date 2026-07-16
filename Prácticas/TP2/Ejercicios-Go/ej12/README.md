# Ejercicio 12 — Tipo Ingresante

## Enunciado

Implementar el tipo de datos **Ingresante** y la funcionalidad solicitada.

**a)** Definir el tipo `Ingresante` con:
- Apellido, Nombre
- Ciudad de origen
- Fecha de nacimiento (día, mes, año)
- Si presentó el título del colegio secundario
- Código de carrera (APU, LI, LS)
- Función `String()` para `fmt.Println`

**b)** Funciones de comparación por edad y por orden alfabético.

**c)** Cargar varios datos en un slice de ingresantes y ordenar por edad y por apellido/nombre. Investigar el package `sort`.

---

## Lógica de resolución (como está en `main.go`)

```go
type Carrera string
const (
    APU Carrera = "APU"
    LI  Carrera = "LI"
    LS  Carrera = "LS"
)

type Fecha struct {
    Dia, Mes, Anio int
}

type Ingresante struct {
    Apellido    string
    Nombre      string
    Ciudad      string
    Nacimiento  Fecha
    TieneTitulo bool
    Carrera     Carrera
}

func (i Ingresante) String() string {
    titulo := "No"
    if i.TieneTitulo { titulo = "Si" }
    return fmt.Sprintf("%s, %s (%s) - Nac: %02d/%02d/%d - Titulo secundario: %s - Carrera: %s",
        i.Apellido, i.Nombre, i.Ciudad,
        i.Nacimiento.Dia, i.Nacimiento.Mes, i.Nacimiento.Anio,
        titulo, i.Carrera)
}
```

### Funciones de comparación (parte b)

```go
// MasJoven: true si a nacio despues que b
func MasJoven(a, b Ingresante) bool {
    if a.Nacimiento.Anio != b.Nacimiento.Anio {
        return a.Nacimiento.Anio > b.Nacimiento.Anio
    }
    if a.Nacimiento.Mes != b.Nacimiento.Mes {
        return a.Nacimiento.Mes > b.Nacimiento.Mes
    }
    return a.Nacimiento.Dia > b.Nacimiento.Dia
}

// MenorAlfabetico: compara apellido y despues nombre
func MenorAlfabetico(a, b Ingresante) bool {
    if a.Apellido != b.Apellido {
        return a.Apellido < b.Apellido
    }
    return a.Nombre < b.Nombre
}
```

### Ordenar con `sort.Slice` (parte c)

```go
sort.Slice(ingresantes, func(i, j int) bool {
    return MasJoven(ingresantes[i], ingresantes[j])
})

sort.Slice(ingresantes, func(i, j int) bool {
    return MenorAlfabetico(ingresantes[i], ingresantes[j])
})
```

## Observaciones

- `sort.Slice` no pide implementar ninguna interfaz: solo necesita el slice y una función `less(i, j int) bool`. Es la forma más simple de ordenar en Go cuando no querés definir `sort.Interface` a mano.
- Cada llamada a `sort.Slice` ordena el mismo slice de nuevo, sobre el resultado del orden anterior — por eso el ejemplo imprime tres veces, para ver cómo cambia el orden.
- La comparación de edad va por año, después mes, después día — hay que comparar de a un campo a la vez porque `Fecha` no tiene una comparación directa entre structs (`>` no existe para structs en Go).
