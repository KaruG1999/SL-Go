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

## Lógica de resolución (paquete `tree`, como está en `tree.go`)

Es el árbol binario del TP2 (`ej11`), con el mismo diseño (contador `len` aparte, `agregarNodo` con doble puntero), pero con `[T cmp.Ordered]` en vez de `int` hardcodeado.

### Constraint `cmp.Ordered`

`cmp.Ordered` (paquete `cmp`, desde Go 1.21) incluye todos los tipos ordenables: enteros, flotantes y strings. Permite usar `<`, `>`, `==` directamente sobre `T`.

```go
type nodoArbol[T cmp.Ordered] struct {
    elem T
    HI   *nodoArbol[T]
    HD   *nodoArbol[T]
}

type ArbolBin[T cmp.Ordered] struct {
    raiz *nodoArbol[T]
    len  int
}

func New[T cmp.Ordered]() ArbolBin[T] {
    return ArbolBin[T]{}
}
```

### Add — inserción de ABB, misma idea que en TP2 pero genérica

```go
func agregarNodo[T cmp.Ordered](ptr **nodoArbol[T], elem T) bool {
    if *ptr == nil {
        *ptr = &nodoArbol[T]{elem: elem}
        return true
    }
    if elem < (*ptr).elem {
        return agregarNodo(&(*ptr).HI, elem)
    } else if elem > (*ptr).elem {
        return agregarNodo(&(*ptr).HD, elem)
    }
    return false
}

func (this ArbolBin[T]) Add(elem T) ArbolBin[T] {
    if agregarNodo(&this.raiz, elem) {
        this.len++
    }
    return this
}
```

### Includes — aprovecha el orden del ABB, igual que en TP2

```go
func (this ArbolBin[T]) Includes(elem T) bool {
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

> `New[int]()`, `New[string]()` — el mismo código sirve para cualquier tipo que soporte `<` y `>`, sin reescribir nada. En `main.go` se prueba con un árbol de `int` y uno de `string`.

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
