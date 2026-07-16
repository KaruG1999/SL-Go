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

## Lógica de resolución (package `mi_arbol`, `arbol.go`)

```go
type nodoArbol struct {
    elem   int
    HI, HD *nodoArbol
}

type arbolBin struct {
    raiz *nodoArbol
    len  int // cuenta de nodos, se mantiene aparte para que Len() sea O(1)
}

type Order int
const (
    PreOrder Order = iota
    InOrder
    PostOrder
)

func New() arbolBin { return arbolBin{} }

func (this arbolBin) IsEmpty() bool { return this.raiz == nil }
func (this arbolBin) GetElem() int  { return this.raiz.elem }
func (this arbolBin) Len() int      { return this.len }
```

### Add — inserción de ABB con doble puntero

```go
func agregarNodo(ptr **nodoArbol, elem int) bool {
    if *ptr == nil {
        *ptr = &nodoArbol{elem: elem}
        return true
    }
    if elem < (*ptr).elem {
        return agregarNodo(&(*ptr).HI, elem)
    } else if elem > (*ptr).elem {
        return agregarNodo(&(*ptr).HD, elem)
    }
    return false // ya existe, no se inserta
}

func (this arbolBin) Add(elem int) arbolBin {
    if agregarNodo(&this.raiz, elem) {
        this.len++
    }
    return this
}
```

El `**nodoArbol` permite reasignar directamente el puntero del nodo padre (`HI` o `HD`) sin tener que devolver el subárbol modificado y reengancharlo a mano. Si el elemento ya está en el árbol, no se agrega y `len` no se incrementa.

### Depth, Traverse, Apply, Find — recursión sobre `*nodoArbol`

```go
func calcularProf(n *nodoArbol) int {
    if n == nil { return 0 }
    profIzq, profDer := calcularProf(n.HI), calcularProf(n.HD)
    if profIzq > profDer { return profIzq + 1 }
    return profDer + 1
}

func traverseRec(n *nodoArbol, fp func(int), o Order) {
    if n == nil { return }
    switch o {
    case PreOrder:
        fp(n.elem); traverseRec(n.HI, fp, o); traverseRec(n.HD, fp, o)
    case InOrder:
        traverseRec(n.HI, fp, o); fp(n.elem); traverseRec(n.HD, fp, o)
    case PostOrder:
        traverseRec(n.HI, fp, o); traverseRec(n.HD, fp, o); fp(n.elem)
    }
}
```

`Traverse`, `Apply` y `Find` son wrappers finitos que llaman a su versión recursiva pasando `this.raiz`.

### Includes — aprovecha que es un ABB

```go
func (this arbolBin) Includes(elem int) bool {
    n := this.raiz
    for n != nil {
        if elem == n.elem {
            return true
        } else if elem < n.elem {
            n = n.HI
        } else {
            n = n.HD
        }
    }
    return false
}
```

No recorre todo el árbol como haría un `Find` genérico: al ser árbol de búsqueda, en cada nodo descarta la mitad, así que es O(log n) en vez de O(n) (si el árbol está razonablemente balanceado).

## Observaciones

- `Len()` no cuenta nodos recorriendo el árbol cada vez, se mantiene como contador (`len`) que se actualiza en `Add`. Cuidado si preguntan por esto: cualquier otra forma de insertar nodos que no pase por `Add` rompería el contador.
- `Add` sobre un elemento repetido no lo agrega y devuelve el árbol sin cambios — se ve probado en `main.go` con `a.Add(5)` después de ya haberlo insertado.
- `Includes` es más rápido que recorrer todo con `Find` porque usa el orden del ABB para descartar la mitad del árbol en cada paso; `Find` en cambio sirve para cualquier condición arbitraria y sí tiene que mirar todos los nodos en el peor caso.
