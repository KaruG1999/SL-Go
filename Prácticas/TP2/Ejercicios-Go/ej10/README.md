# Ejercicio 10 — Pila de Enteros

## Enunciado

Implementar usando **slices** una pila (stack) de enteros con las operaciones:

```go
func New() Stack
func IsEmpty(s Stack) bool
func Len(s Stack) int
func ToString(s Stack) string
func FrontElement(s Stack) int
func Push(s Stack, element int)
func Pop(s Stack) int
func Iterate(s Stack, f func(int) int)
```

**b)** Re-implementar usando la lista enlazada del ejercicio 9.

---

## Lógica de resolución

### Tipo

```go
type Stack []int
```

Un slice de enteros es suficiente para implementar una pila. El **tope** de la pila es el **último elemento** del slice.

### Operaciones con slice

```go
func New() Stack { return Stack{} }

func IsEmpty(s Stack) bool { return len(s) == 0 }

func Len(s Stack) int { return len(s) }

func FrontElement(s Stack) int {
    if IsEmpty(s) { panic("pila vacía") }
    return s[len(s)-1]
}

func Push(s *Stack, element int) {
    *s = append(*s, element)
}

func Pop(s *Stack) int {
    if IsEmpty(*s) { panic("pila vacía") }
    top := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return top
}

func ToString(s Stack) string {
    return fmt.Sprintf("%v", []int(s))
}

func Iterate(s Stack, f func(int) int) {
    for i := range s {
        s[i] = f(s[i])
    }
}
```

> `Push` y `Pop` reciben `*Stack` porque modifican el slice subyacente (cambian su longitud). `append` puede reubicar el slice en memoria si la capacidad es insuficiente, por lo que hay que actualizar el puntero en el caller.

### Re-implementación con Lista (parte b)

Usar la `List` del ejercicio 9 como backing store:

```go
type Stack struct {
    data List
}

func (s *Stack) Push(elem int) { PushFront(&s.data, elem) }
func (s *Stack) Pop() int      { return Remove(&s.data) }
func (s *Stack) Top() int      { return FrontElement(s.data) }
```

> Tanto el tope de la pila como el frente de la lista son la misma posición, lo que hace que Push y Pop sean O(1).
