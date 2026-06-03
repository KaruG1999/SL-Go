# Ejercicio 3 — Árbol binario genérico

## Enunciado

Re-implementar el árbol binario de la práctica anterior para que funcione de forma genérica, usando:

```go
type NodeCmpBinTree[T cmp.Ordered] struct {
    elem  T
    left  *NodeCmpBinTree[T]
    right *NodeCmpBinTree[T]
}
type CmpBinTree[T cmp.Ordered] struct {
    tree *NodeCmpBinTree[T]
}
```

*Objetivo: tipos genéricos*

---

## Lógica de resolución

### Constraint `cmp.Ordered`

`cmp.Ordered` (del paquete `cmp`, disponible desde Go 1.21) incluye todos los tipos ordenables: enteros, flotantes y strings. Permite usar `<`, `>`, `==` directamente sobre `T`.

```go
import "cmp"
```

### Insert (BST)

```go
func insert[T cmp.Ordered](n *NodeCmpBinTree[T], val T) *NodeCmpBinTree[T] {
    if n == nil {
        return &NodeCmpBinTree[T]{elem: val}
    }
    if val < n.elem {
        n.left = insert(n.left, val)
    } else if val > n.elem {
        n.right = insert(n.right, val)
    }
    return n
}

func (t *CmpBinTree[T]) Insert(val T) {
    t.tree = insert(t.tree, val)
}
```

### Search

```go
func (t *CmpBinTree[T]) Search(val T) bool {
    curr := t.tree
    for curr != nil {
        switch {
        case val == curr.elem: return true
        case val < curr.elem:  curr = curr.left
        default:               curr = curr.right
        }
    }
    return false
}
```

### Recorrido inorder (imprime ordenado)

```go
func inorder[T cmp.Ordered](n *NodeCmpBinTree[T]) {
    if n == nil { return }
    inorder(n.left)
    fmt.Println(n.elem)
    inorder(n.right)
}
```

> La misma implementación sirve para `CmpBinTree[int]`, `CmpBinTree[float64]` o `CmpBinTree[string]` sin cambios.

---

## Nota: `nodoArbol[T]` vs `*nodoArbol[T]` vs `**nodoArbol[T]`

```go
nodoArbol[T]       // el vagón: la estructura con el dato y los hijos
*nodoArbol[T]      // la cadena: sabe dónde está el vagón en memoria
**nodoArbol[T]     // variable que guarda la ubicación de la cadena
```

**`nodoArbol[T]`** — es el vagón en sí. Tiene el dato (`elem`) y dos referencias a sus hijos (`HI`, `HD`).

**`*nodoArbol[T]`** — es una cadena que apunta al vagón. Si la cadena vale `nil`, no hay vagón ahí. Es lo que usan `HI` y `HD` dentro del nodo: cada hijo es una cadena que puede apuntar a otro vagón o estar cortada (`nil`).

**`**nodoArbol[T]`** — es una variable que guarda la ubicación de esa cadena. Sirve para poder **reemplazar** la cadena por otra. En `agregarNodo` se usa así:

```go
func agregarNodo[T cmp.Ordered](ptr **nodoArbol[T], elem T) bool {
    if *ptr == nil {
        *ptr = &nodoArbol[T]{elem: elem} // reemplazamos la cadena cortada por una nueva
        return true
    }
    ...
}
```

Si solo recibiéramos `*nodoArbol[T]`, podríamos ver el vagón al que apunta, pero no podríamos cambiar qué vagón está enganchado en ese lugar. Con `**nodoArbol[T]` tenemos en la mano la cadena misma, y podemos engancharla a un vagón nuevo.

Cuando llamamos `agregarNodo(&this.raiz, elem)`, le pasamos la dirección de `raiz` — es decir, la cadena que une el árbol con su primer vagón — para poder reemplazarla si el árbol estaba vacío.
