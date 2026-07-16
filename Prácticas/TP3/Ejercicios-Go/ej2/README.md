# Ejercicio 2 — Lista enlazada genérica

## Enunciado

Definir e implementar las operaciones de lista enlazada de la práctica anterior usando el siguiente tipo genérico:

```go
type List[T any] struct {
    head, tail *element[T]
}
type element[T any] struct {
    next *element[T]
    val  T
}
```

Comparar con la versión sin genéricos (usando `any`):

```go
type element struct {
    next *element
    val  any
}
type List struct {
    head, tail *element
}
```

Responder: ¿es necesario cambiar la definición del tipo para poder tener una función `Find`?

*Objetivo: tipos genéricos*

---

## Lógica de resolución (paquete `list`, como está en `list.go`)

```go
type element[T any] struct {
    next *element[T]
    val  T
}

type List[T any] struct {
    ini, fin *element[T]
    len      int
}

func New[T any]() *List[T] {
    return &List[T]{ini: nil, fin: nil, len: 0}
}
```

### PushBack / PushFront

```go
func (l *List[T]) PushBack(val T) {
    nuevo := &element[T]{val: val, next: nil}
    if l.IsEmpty() {
        l.ini = nuevo
        l.fin = nuevo
    } else {
        l.fin.next = nuevo
        l.fin = nuevo
    }
    l.len++
}

func (l *List[T]) PushFront(val T) {
    nuevo := &element[T]{val: val, next: l.ini}
    if l.IsEmpty() {
        l.fin = nuevo
    }
    l.ini = nuevo
    l.len++
}
```

### Find

```go
func (l *List[T]) Find(criterio func(T) bool) (T, bool) {
    actual := l.ini
    for actual != nil {
        if criterio(actual.val) {
            return actual.val, true
        }
        actual = actual.next
    }
    var valorVacio T
    return valorVacio, false
}
```

## ¿Es necesario cambiar la definición del tipo para tener `Find`?

**No.** `Find` recibe una función `criterio func(T) bool` en vez de comparar `val == algo` con `==`. Como nunca se usa `==` sobre `T`, el constraint se queda en `any` — no hace falta pasar a `comparable`.

Esto se ve en `main.go`: se llama `Find` con una lista de `int` buscando "mayor a 20" (`func(i int) bool { return i > 20 }`) y con una lista de `Persona` buscando por nombre (`func(p Persona) bool { return p.Nombre == "Karen" }`). En ningún caso se compara el valor completo con `==`, se compara lo que haga falta *adentro* de la función criterio — así `Find` funciona para cualquier tipo, comparable o no.

Si en cambio `Find` recibiera directamente un valor a comparar (`Find(val T) bool { return curr.val == val }`), ahí sí haría falta cambiar el constraint a `[T comparable]`, porque `==` no está definido para cualquier `T any` (por ejemplo, no compila para slices o para structs con campos no comparables).

## Observaciones

- La versión sin genéricos (`val any`) obliga a hacer type assertion para leer el valor (`val.(int)`), y perdés el chequeo de tipos en compilación: nada impide meter un `int` y un `string` en la misma lista por error.
- Con genéricos, el tipo queda fijado al instanciar (`list.New[int]()`, `list.New[Persona]()`), y el compilador verifica que todo lo que se agrega sea de ese tipo.
