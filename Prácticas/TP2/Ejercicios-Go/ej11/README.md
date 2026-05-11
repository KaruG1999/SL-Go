# Ejercicio 11 — Árbol Binario de Enteros

## Enunciado

Definir e implementar el package **Árbol Binario de enteros** con la siguiente interfaz:

```go
func New() IntBinTree
func (this IntBinTree) IsEmpty() bool
func (this IntBinTree) GetElem() int
func (this IntBinTree) GetRight() IntBinTree
func (this IntBinTree) GetLeft() IntBinTree
func (this IntBinTree) Len() int
func (this IntBinTree) Depth() int
func (this IntBinTree) Add(elem int) IntBinTree
func (this IntBinTree) Traverse(fp func(int), o Order)
func (this IntBinTree) Apply(fp func(int) int)
func (this IntBinTree) Includes(elem int) bool
func (this IntBinTree) Find(fp func(int) bool) bool
func (this IntBinTree) String() string
```

---

## Lógica de resolución

### Tipos

```go
type Order int
const (
    InOrder Order = iota   // izquierda → raíz → derecha
    PreOrder               // raíz → izquierda → derecha
    PostOrder              // izquierda → derecha → raíz
)

type treeNode struct {
    val         int
    left, right *treeNode
}

// IntBinTree es un puntero al nodo raíz; nil = árbol vacío
type IntBinTree *treeNode
```

### Operaciones básicas

```go
func New() IntBinTree { return nil }

func (t IntBinTree) IsEmpty() bool  { return t == nil }
func (t IntBinTree) GetElem() int   { return t.val }
func (t IntBinTree) GetLeft() IntBinTree  { return IntBinTree(t.left) }
func (t IntBinTree) GetRight() IntBinTree { return IntBinTree(t.right) }
```

### Add — árbol binario de búsqueda

```go
func (t IntBinTree) Add(elem int) IntBinTree {
    if t == nil {
        return &treeNode{val: elem}
    }
    if elem <= t.val {
        t.left = t.Add(elem)   // no, espera: ver abajo
    } else {
        t.right = t.Add(elem)
    }
    return t
}
```

> Ojo: `t` es un puntero, así que `t.left = ...` modifica el nodo. La función retorna el árbol (necesario para el caso raíz nil).

### Len y Depth

```go
func (t IntBinTree) Len() int {
    if t == nil { return 0 }
    return 1 + t.GetLeft().Len() + t.GetRight().Len()
}

func (t IntBinTree) Depth() int {
    if t == nil { return 0 }
    l, r := t.GetLeft().Depth(), t.GetRight().Depth()
    if l > r { return 1 + l }
    return 1 + r
}
```

### Traverse — recorrido según orden

```go
func (t IntBinTree) Traverse(fp func(int), o Order) {
    if t == nil { return }
    switch o {
    case PreOrder:
        fp(t.val); t.GetLeft().Traverse(fp, o); t.GetRight().Traverse(fp, o)
    case InOrder:
        t.GetLeft().Traverse(fp, o); fp(t.val); t.GetRight().Traverse(fp, o)
    case PostOrder:
        t.GetLeft().Traverse(fp, o); t.GetRight().Traverse(fp, o); fp(t.val)
    }
}
```

### Apply, Includes, Find

```go
func (t IntBinTree) Apply(fp func(int) int) {
    if t == nil { return }
    t.val = fp(t.val)
    t.GetLeft().Apply(fp); t.GetRight().Apply(fp)
}

func (t IntBinTree) Includes(elem int) bool {
    if t == nil { return false }
    if t.val == elem { return true }
    return t.GetLeft().Includes(elem) || t.GetRight().Includes(elem)
}

func (t IntBinTree) Find(fp func(int) bool) bool {
    if t == nil { return false }
    if fp(t.val) { return true }
    return t.GetLeft().Find(fp) || t.GetRight().Find(fp)
}
```

### String

```go
func (t IntBinTree) String() string {
    if t == nil { return "[]" }
    var elems []string
    t.Traverse(func(v int) {
        elems = append(elems, fmt.Sprintf("%d", v))
    }, InOrder)
    return "[" + strings.Join(elems, " ") + "]"
}
```
