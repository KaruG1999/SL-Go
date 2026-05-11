# Seminario de Lenguajes — Go — Clase 6: Genéricos

---

## Contexto de Conexión

Hasta ahora, cada función o tipo estaba atado a un tipo concreto. Si queríamos la misma lógica para `int64` y `float64`, teníamos que duplicar código. Los **genéricos** (introducidos en Go 1.18) permiten escribir funciones y tipos parametrizados por tipo, eliminando esa duplicación sin perder el chequeo estático de tipos.

---

## Conceptos Core

- **Type parameter**: parámetro de tipo declarado entre corchetes `[T constraint]`. Actúa como un "placeholder" para un tipo concreto que se especificará al usar la función o tipo.
- **Constraint**: restricción que indica qué tipos puede tomar el type parameter. Puede ser `any`, `comparable`, una interfaz, o una unión de tipos `int64 | float64`.
- **`any`**: alias de `interface{}`. Acepta cualquier tipo.
- **`comparable`**: tipos que soportan `==` y `!=` (necesario para usar como clave de map).
- **Type inference**: Go puede inferir los type parameters a partir de los argumentos pasados, sin necesidad de especificarlos explícitamente.
- **Tipo genérico**: struct o tipo definido con type parameters, por ejemplo `List[T any]`.

---

## Desarrollo

### 1. Motivación: el problema sin genéricos

Supongamos que queremos sumar los valores de un map. Necesitamos una función por tipo:

```go
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m { s += v }
    return s
}

func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m { s += v }
    return s
}
```

Misma lógica, código duplicado. Con genéricos se unifica en una sola función.

---

### 2. Función genérica

```go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

**Uso — especificando los tipos explícitamente:**
```go
fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))
```

**Uso — con inferencia de tipos:**
```go
fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))
```

**Anatomía de los type parameters:**
```
[K comparable, V int64 | float64]
 ↑                ↑
 parámetro formal  constraint (unión de tipos)
```

- `K comparable` → la clave del map puede ser cualquier tipo comparable.
- `V int64 | float64` → el valor debe ser uno de esos dos tipos.

---

### 3. Formas de constraints

```go
[T any]                                          // cualquier tipo
[T comparable]                                   // tipos con == y !=
[T Stringer]                                     // tipos que implementen la interfaz Stringer
[T int | int8 | int16 | int32 | int64 | float32 | float64]  // unión explícita
[T Stringer | error]                             // unión de interfaces
```

---

### 4. Tipo genérico — Lista enlazada

```go
type List[T any] struct {
    first, last *node[T]
}

type node[T any] struct {
    val  T
    next *node[T]
}

func (l *List[T]) PutOnFront(v T) {
    l.first = &node[T]{v, l.first}
    if l.last == nil {
        l.last = l.first
    }
}

func (l *List[T]) PutOnTail(v T) {
    n := &node[T]{val: v}
    if l.last == nil {
        l.first = n
    } else {
        l.last.next = n
    }
    l.last = n
}

func (l *List[T]) GetAll() []T {
    var elems []T
    for e := l.first; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}
```

**Uso:**
```go
list := List[int]{}
list.PutOnFront(10)
list.PutOnTail(20)
list.PutOnFront(30)
list.PutOnTail(40)
list.PutOnFront(50)
list.PutOnTail(60)
fmt.Println("list:", list.GetAll())
// list: [50 30 10 20 40 60]
```

---

### 5. Tipo genérico — Árbol binario

```go
type Tree[T any] struct {
    val         T
    left, right *Tree[T]
}

func (t *Tree[T]) insert(v T, f func(T, T) bool) *Tree[T] {
    if t == nil {
        return &Tree[T]{val: v}
    }
    if f(v, t.val) {
        t.left = t.left.insert(v, f)
    } else {
        t.right = t.right.insert(v, f)
    }
    return t
}

func (t *Tree[T]) GetAll() []T {
    var elems []T
    if t != nil {
        elems = append(elems, t.left.GetAll()...)
        elems = append(elems, t.val)
        elems = append(elems, t.right.GetAll()...)
    }
    return elems
}
```

**Uso:**
```go
func lt(x, y int) bool { return x <= y }

var tree *Tree[int]
tree = tree.insert(50, lt)
tree = tree.insert(10, lt)
tree = tree.insert(90, lt)
// ...
fmt.Println("Tree:", tree.GetAll())
// Tree: [10 30 40 50 60 80 90]
```

> El árbol recibe una función de comparación como parámetro, lo que permite ordenar cualquier tipo `T` sin que `T` deba implementar ninguna interfaz.

---

## Lo que no podés ignorar

> 1. **Los type parameters van entre corchetes `[]`, no entre `<>`**: es la sintaxis de Go, diferente a Java o C++.
> 2. **La constraint define qué operaciones son válidas**: si declarás `[T any]` no podés sumar `T + T` porque `any` no garantiza que `+` exista. Necesitás `int | float64` o una interfaz con ese método.
> 3. **`comparable` es obligatorio para claves de map**: al declarar `map[K]V`, `K` debe ser `comparable`.
> 4. **La inferencia de tipos funciona en la mayoría de los casos**: si los argumentos dejan clara la instanciación, no hace falta escribir `[string, int64]` explícitamente.
> 5. **Los métodos de un tipo genérico usan `[T]` en el receiver**: `func (l *List[T]) PutOnFront(v T)`, no `func (l *List) PutOnFront[T](v T)`.
> 6. **Genéricos ≠ interfaces**: las interfaces resuelven polimorfismo en tiempo de ejecución (dynamic dispatch); los genéricos resuelven polimorfismo en tiempo de compilación (el compilador genera código específico para cada instanciación).
