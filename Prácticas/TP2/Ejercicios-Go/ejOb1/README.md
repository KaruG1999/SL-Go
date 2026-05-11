# Ejercicio Obligatorio 1 — Lista de Ingresantes

## Enunciado

Usando la estructura de datos del ejercicio 9 (lista enlazada), resolver el siguiente problema. Se dispone de una lista con información de ingresantes a la Facultad. De cada ingresante se conoce: apellido, nombre, ciudad de origen, fecha de nacimiento (día, mes, año), si presentó el título y el código de la carrera (APU, LI, LS).

Recorrer la lista **una sola vez** para:

**a)** Informar nombre y apellido de los ingresantes cuya ciudad de origen es "Bariloche".

**b)** Calcular e informar el año en que más ingresantes nacieron.

**c)** Informar la carrera con mayor cantidad de inscriptos.

**d)** Eliminar de la lista aquellos ingresantes que no presentaron el título.

---

## Lógica de resolución

### Tipo Ingresante (del ejercicio 12)

```go
type Ingresante struct {
    Apellido    string
    Nombre      string
    Ciudad      string
    Nacimiento  Fecha
    TieneTitulo bool
    Carrera     string
}
```

### Lista enlazada de Ingresantes (del ejercicio 9, adaptada)

```go
type node struct {
    val  Ingresante
    next *node
}
type List *node
```

### Recorrido único

La clave es resolver **a, b, c y d en un solo recorrido**. Para eso, se acumulan todos los datos necesarios mientras se itera:

```go
anioConteo  := make(map[int]int)
carreraConteo := make(map[string]int)

var prev *node = nil
curr := lista

for curr != nil {
    ing := curr.val

    // a) ciudad = Bariloche
    if ing.Ciudad == "Bariloche" {
        fmt.Printf("%s, %s\n", ing.Apellido, ing.Nombre)
    }

    // b) contar por año de nacimiento
    anioConteo[ing.Nacimiento.Anio]++

    // c) contar por carrera
    carreraConteo[ing.Carrera]++

    // d) eliminar si no tiene título
    if !ing.TieneTitulo {
        if prev == nil {
            lista = curr.next   // eliminar cabeza
        } else {
            prev.next = curr.next
        }
        curr = curr.next
        continue
    }

    prev = curr
    curr = curr.next
}
```

### Post-recorrido: calcular máximos

```go
// b) año más frecuente
maxAnio, maxCantAnio := 0, 0
for anio, cant := range anioConteo {
    if cant > maxCantAnio {
        maxCantAnio = cant
        maxAnio = anio
    }
}
fmt.Printf("Año con más ingresantes: %d (%d ingresantes)\n", maxAnio, maxCantAnio)

// c) carrera más popular
maxCarrera, maxCantCarrera := "", 0
for carrera, cant := range carreraConteo {
    if cant > maxCantCarrera {
        maxCantCarrera = cant
        maxCarrera = carrera
    }
}
fmt.Printf("Carrera con más inscriptos: %s (%d)\n", maxCarrera, maxCantCarrera)
```

> La eliminación durante el recorrido (`d`) requiere mantener un puntero `prev` al nodo anterior para poder "saltar" el nodo a eliminar redirigiendo `prev.next`. Es el patrón clásico de eliminación en lista enlazada simple sin centinela.
