# Ejercicio 2 — Lista enlazada genérica

## Enunciado

Reimplementar las operaciones de lista enlazada de la práctica anterior usando el siguiente tipo genérico:

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

## Lógica de resolución

### Diferencia clave entre versiones

- Con `any`: el compilador no sabe el tipo de `val`, se necesita type assertion al leer (`val.(int)`, etc.) y `Find` no puede comparar valores directamente.
- Con genéricos `[T any]`: el tipo está fijado en la instanciación (`List[int]`, `List[string]`), pero `T any` **no es `comparable`**, así que `Find` tampoco puede usar `==`.
- Para poder comparar en `Find`, el constraint debe ser `[T comparable]`.

### Método Insert

```go
func (l *List[T]) Insert(val T) {
    node := &element[T]{val: val}
    if l.head == nil {
        l.head = node
        l.tail = node
        return
    }
    l.tail.next = node
    l.tail = node
}
```

### Método Find (requiere `[T comparable]`)

```go
func (l *List[T]) Find(val T) bool {
    curr := l.head
    for curr != nil {
        if curr.val == val {
            return true
        }
        curr = curr.next
    }
    return false
}
```

> Conclusión: sí es necesario cambiar el constraint de `any` a `comparable` para poder implementar `Find` con `==`.
