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

## Lógica de resolución (`main.go`)

```go
type Stack struct {
    data []int
}

func New() Stack { return Stack{data: []int{}} }

func (s Stack) IsEmpty() bool { return len(s.data) == 0 }
func (s Stack) Len() int      { return len(s.data) }

func (s Stack) Top() (int, error) {
    if s.IsEmpty() {
        return 0, errors.New("pila vacia")
    }
    return s.data[len(s.data)-1], nil
}

func (s *Stack) Push(element int) {
    s.data = append(s.data, element)
}

func (s *Stack) Pop() (int, error) {
    if s.IsEmpty() {
        return 0, errors.New("Pila vacia")
    }
    top := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return top, nil
}
```

El tope de la pila es el último elemento del slice. `String()` implementa Stringer imprimiendo desde el tope hacia el fondo.

## Parte b — con la lista del ejercicio 9

`b/main.go` reutiliza la `List` (versión con métodos y errores, la del ej9 parte d) como backing store en vez de un slice:

```go
type Stack struct {
    data List
}

func (s *Stack) Push(element int) { s.data.PushFront(element) }
func (s *Stack) Pop() (int, error) { return s.data.Remove() }
func (s Stack) Top() (int, error)  { return s.data.FrontElement() }
```

Push equivale a PushFront y Pop a Remove: el tope de la pila y el frente de la lista son la misma posición, así que ambas operaciones quedan en O(1).

## Observaciones

- `Push` y `Pop` necesitan receiver puntero (`*Stack`) porque `append` puede reasignar el slice interno si no entra en la capacidad actual; sin el puntero ese cambio se perdería al salir de la función.
- A diferencia de la versión con slice, la versión con lista (parte b) no tiene ese problema de reubicación en memoria, porque cada nodo se reserva por separado.
- Tanto `Top` como `Pop` devuelven `error` en vez de hacer panic en pila vacía — ya viene con el mismo criterio que se usó en el ej9 parte d.
