# Ejercicio 9 — Lista Enlazada de Enteros

## Enunciado

Usando memoria dinámica con punteros, implementar una lista enlazada de enteros con las operaciones:

```go
func New() List
func IsEmpty(self List) bool
func Len(self List) int
func FrontElement(self List) int
func Next(self List) List
func ToString(self List) string
func PushFront(self List, elem int)
func PushBack(self List, elem int)
func Remove(self List) int
func Iterate(self List, f func(int) int)
```

**b)** Programa que use las operaciones.

**c)** Investigar y usar el paquete `container/list`. Diferencias con la implementación propia.

**d)** Mejorar la interfaz usando métodos y códigos de error.

---

## Lógica de resolución

### Tipos

```go
type node struct {
    val  int
    next *node
}

// List es un puntero al nodo raíz — nil representa la lista vacía
type List *node
```

> En Go, el puntero `nil` representa naturalmente la lista vacía. `List` es un alias de `*node`.

### Operaciones básicas

```go
func New() List { return nil }

func IsEmpty(self List) bool { return self == nil }

func Len(self List) int {
    count := 0
    for n := self; n != nil; n = n.next {
        count++
    }
    return count
}

func FrontElement(self List) int { return self.val }

func Next(self List) List { return self.next }
```

### PushFront y PushBack

```go
// PushFront: nuevo nodo al inicio
func PushFront(self *List, elem int) {
    *self = &node{val: elem, next: *self}
}

// PushBack: recorrer hasta el final y agregar
func PushBack(self *List, elem int) {
    nuevo := &node{val: elem}
    if *self == nil {
        *self = nuevo
        return
    }
    n := *self
    for n.next != nil {
        n = n.next
    }
    n.next = nuevo
}
```

> Para modificar la lista (cambiar qué nodo es el primero), la función necesita recibir `*List` (puntero al puntero).

### Remove y ToString

```go
func Remove(self *List) int {
    val := (*self).val
    *self = (*self).next
    return val
}

func ToString(self List) string {
    s := "["
    for n := self; n != nil; n = n.next {
        if n != self { s += " -> " }
        s += fmt.Sprintf("%d", n.val)
    }
    return s + "]"
}
```

### Iterate — aplicar función a cada elemento

```go
func Iterate(self List, f func(int) int) {
    for n := self; n != nil; n = n.next {
        n.val = f(n.val)
    }
}
```

### Versión con métodos (parte d)

Definir `List` como struct con puntero al primer nodo y puntero al último (para PushBack en O(1)):

```go
type List struct {
    head, tail *node
    size       int
}

func (l *List) PushFront(elem int) { ... }
func (l *List) PushBack(elem int)  { ... }
func (l *List) IsEmpty() bool       { return l.head == nil }
```

> Con métodos, la sintaxis `lista.PushFront(5)` es más natural y el compilador puede hacer type-checking más fino.
