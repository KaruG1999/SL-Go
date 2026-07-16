# Ejercicio Obligatorio 1 — Lista de Ingresantes

## Enunciado

Usando la estructura de datos del ejercicio 9 (lista enlazada), resolver el siguiente problema. Se dispone de una lista con información de ingresantes a la Facultad. De cada ingresante se conoce: apellido, nombre, ciudad de origen, fecha de nacimiento (día, mes, año), si presentó el título y el código de la carrera (APU, LI, LS).

Recorrer la lista **una sola vez** para:

**a)** Informar nombre y apellido de los ingresantes cuya ciudad de origen es "Bariloche".

**b)** Calcular e informar el año en que más ingresantes nacieron.

**c)** Informar la carrera con mayor cantidad de inscriptos.

**d)** Eliminar de la lista aquellos ingresantes que no presentaron el título.

---

## Lógica de resolución (como está en `main.go`)

### Tipos

```go
type Fecha struct {
    Dia, Mes, Anio int
}

type Ingresante struct {
    Apellido    string
    Nombre      string
    Ciudad      string
    Nacimiento  Fecha
    TieneTitulo bool
    Carrera     string
}

type nodo struct {
    elem Ingresante
    sig  *nodo
}

type List struct {
    pri *nodo
    len int
}
```

### Recorrido único

La función recibe `*List` (puntero a la struct, no a un nodo suelto), así que cualquier cambio sobre `l.pri` o `l.len` persiste en el `main`:

```go
func procesarLista(l *List) {
    anioConteo := make(map[int]int)
    carreraConteo := make(map[string]int)

    var ant *nodo = nil
    act := l.pri

    for act != nil {
        ing := act.elem

        // a) ciudad = Bariloche
        if ing.Ciudad == "Bariloche" {
            fmt.Printf("   %s, %s\n", ing.Apellido, ing.Nombre)
        }

        // b) contar por año de nacimiento
        anioConteo[ing.Nacimiento.Anio]++

        // c) contar por carrera
        carreraConteo[ing.Carrera]++

        // d) eliminar si no tiene título
        if !ing.TieneTitulo {
            if ant == nil {
                l.pri = act.sig // era la cabeza, se reasigna sobre *List
            } else {
                ant.sig = act.sig
            }
            l.len--
            act = act.sig
            continue
        }

        ant = act
        act = act.sig
    }

    // b) año más frecuente
    maxAnio, maxCantAnio := 0, 0
    for anio, cant := range anioConteo {
        if cant > maxCantAnio {
            maxCantAnio = cant
            maxAnio = anio
        }
    }

    // c) carrera más popular
    maxCarrera, maxCant := "", 0
    for carrera, cant := range carreraConteo {
        if cant > maxCant {
            maxCant = cant
            maxCarrera = carrera
        }
    }
}
```

Se llama como `procesarLista(&l)` desde `main`.

## Variantes de la eliminación (parte d)

**Mover a una lista de "rechazados" en vez de descartar:** se recibe un segundo `*List` y, en vez de solo desenganchar el nodo, se lo reinserta al frente de esa otra lista:

```go
if !ing.TieneTitulo {
    eliminado := act

    if ant == nil {
        l.pri = act.sig
    } else {
        ant.sig = act.sig
    }
    l.len--
    act = act.sig // guardar el siguiente ANTES de reusar eliminado.sig

    eliminado.sig = rechazados.pri
    rechazados.pri = eliminado
    rechazados.len++
    continue
}
```

El orden importa: hay que guardar `act = act.sig` antes de pisar `eliminado.sig` con el enganche a `rechazados`, si no se pierde la referencia al verdadero siguiente nodo de la lista original.

**Eliminar sin `ant`, con puntero a puntero:** en vez de arrastrar una variable `ant`, se usa `aux **nodo` que apunta al lugar donde vive el puntero a modificar (puede ser `l.pri` o el `sig` de un nodo, da igual, ambos son `*nodo`):

```go
aux := &l.pri
for *aux != nil {
    if !(*aux).elem.TieneTitulo {
        *aux = (*aux).sig
        l.len--
    } else {
        aux = &(*aux).sig
    }
}
```

Esta versión resuelve solo el borrado; para reusarla habría que meter adentro del mismo `for` la lectura de `ing := (*aux).elem` y las cuentas de a/b/c antes de decidir si se borra o se avanza.

## Observaciones

- **Recorrido único:** las cuatro consignas se resuelven en una sola pasada por la lista, acumulando en `anioConteo` y `carreraConteo` mientras se recorre. Ir calculando el máximo de cada map después es aparte, pero no vuelve a tocar la lista.
- **Puntero `ant` para eliminar:** en una lista simplemente enlazada no hay forma de "volver atrás" desde `act`, por eso hace falta guardar el nodo anterior a mano. Cuando `act` se elimina, `ant.sig = act.sig` salta el nodo sin romper la cadena.
- **Por qué el borrado de la cabeza persiste:** `procesarLista` recibe `*List` (puntero a la struct, no a un nodo suelto), así que reasignar `l.pri` dentro de la función cambia la lista real, no una copia. Si en cambio la función recibiera `List` por valor (sin el `*`), esa misma línea solo cambiaría una copia local y el ingresante sin título seguiría apareciendo como cabeza al volver al `main` — es el error clásico de pasaje de punteros en Go, pero acá está bien evitado.
- **`elem Ingresante` por valor vs `elem *Ingresante`:** guardar el `Ingresante` completo en cada nodo hace que cada lectura (`ing := act.elem`) copie la struct entera (los campos de `Fecha`, el bool, y los headers de los strings). No es que se dupliquen los textos — los strings en Go son inmutables y copiarlos solo copia un header de 16 bytes, no el contenido — pero sí es una copia más grande que si el nodo guardara `*Ingresante` (un puntero de 8 bytes). Para esta lista, con pocos ingresantes, no importa; en una lista grande sería la primera optimización a probar.
- **¿Por qué `ant` es una variable local y no un campo del nodo?** Porque `nodo` solo tiene `sig` (lista simplemente enlazada): ningún nodo sabe quién es su anterior, así que hay que ir guardándolo a mano mientras se recorre. La alternativa sería agregar un campo `ant *nodo` a la struct (lista doblemente enlazada, como `container/list` del ej9), y ahí para borrar alcanzaría con `act.ant.sig = act.sig` sin variable local. No es que la forma actual esté mal: para un enunciado que solo pide recorrer para adelante, agregar `ant` al nodo sería peso de más (un puntero por nodo, y el doble de punteros para mantener consistentes en cada inserción/borrado) sin necesitarlo.
- **Empate en el año con más ingresantes:** el código se queda con el primer año que encuentra recorriendo `anioConteo` con el máximo conteo visto hasta ahí. El orden de iteración de un map en Go es aleatorio a propósito, así que ante un empate el resultado puede variar entre corridas del mismo programa con los mismos datos.
- **Lista vacía:** si `l` no tiene ingresantes, el `for` no entra nunca, los maps quedan vacíos y `maxAnio`/`maxCarrera` terminan en su zero value (`0` y `""`), que es lo que se termina imprimiendo sin ningún aviso de que no había datos.
