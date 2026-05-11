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

## Lógica de resolución

### Tipos

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
    Apellido  string
    Nombre    string
    Ciudad    string
    Nacimiento Fecha
    TieneTitulo bool
    Carrera   Carrera
}
```

### Parte a — String()

```go
func (i Ingresante) String() string {
    titulo := "No"
    if i.TieneTitulo { titulo = "Sí" }
    return fmt.Sprintf("%s, %s | %s | Nac: %02d/%02d/%d | Título: %s | Carrera: %s",
        i.Apellido, i.Nombre, i.Ciudad,
        i.Nacimiento.Dia, i.Nacimiento.Mes, i.Nacimiento.Anio,
        titulo, i.Carrera)
}
```

### Parte b — funciones de comparación

```go
// Por edad: mayor año de nacimiento = más joven
func MasJoven(a, b Ingresante) bool {
    fa, fb := a.Nacimiento, b.Nacimiento
    if fa.Anio != fb.Anio { return fa.Anio > fb.Anio }
    if fa.Mes != fb.Mes   { return fa.Mes > fb.Mes }
    return fa.Dia > fb.Dia
}

// Por orden alfabético (apellido, luego nombre)
func MenorAlfabetico(a, b Ingresante) bool {
    if a.Apellido != b.Apellido { return a.Apellido < b.Apellido }
    return a.Nombre < b.Nombre
}
```

### Parte c — ordenar con el package sort

```go
import "sort"

ingresantes := []Ingresante{ /* ... */ }

// Ordenar por edad (más jóvenes primero)
sort.Slice(ingresantes, func(i, j int) bool {
    return MasJoven(ingresantes[i], ingresantes[j])
})

// Ordenar por apellido y nombre
sort.Slice(ingresantes, func(i, j int) bool {
    return MenorAlfabetico(ingresantes[i], ingresantes[j])
})
```

> `sort.Slice` recibe el slice y una función de comparación `less(i, j int) bool` que retorna true si el elemento i debe ir antes que el j. Es la forma idiomática de ordenar en Go sin implementar ninguna interfaz.
