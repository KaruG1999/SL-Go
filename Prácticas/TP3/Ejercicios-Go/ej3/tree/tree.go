package tree

import (
	"cmp" // herramientas básicas para comparar valores
	"fmt"
	"strings"
)

// CAMBIO: los structs ahora tienen parámetro de tipo [T cmp.Ordered]
// En TP2 el tipo del elemento estaba hardcodeado como int
type nodoArbol[T cmp.Ordered] struct {
	elem T
	HI   *nodoArbol[T]
	HD   *nodoArbol[T]
}

type ArbolBin[T cmp.Ordered] struct {
	raiz *nodoArbol[T]
	len  int
}

type Order int

const (
	PreOrder  Order = iota
	InOrder
	PostOrder
)

// CAMBIO: New ahora es una función genérica, el caller elige el tipo: New[int](), New[string]()
func New[T cmp.Ordered]() ArbolBin[T] {
	return ArbolBin[T]{}
}

func (this ArbolBin[T]) IsEmpty() bool {
	return this.raiz == nil
}

func (this ArbolBin[T]) GetElem() T {
	return this.raiz.elem
}

func (this ArbolBin[T]) GetRight() ArbolBin[T] {
	return ArbolBin[T]{raiz: this.raiz.HD}
}

func (this ArbolBin[T]) GetLeft() ArbolBin[T] {
	return ArbolBin[T]{raiz: this.raiz.HI}
}

func (this ArbolBin[T]) Len() int {
	return this.len
}

func (this ArbolBin[T]) Depth() int {
	return calcularProf(this.raiz)
}

func calcularProf[T cmp.Ordered](n *nodoArbol[T]) int {
	if n == nil {
		return 0
	}
	// Almacena prof de los hijos y los almacena en hd e hi
	profIzq := calcularProf(n.HI)
	profDer := calcularProf(n.HD)
	if profIzq > profDer {
		return profIzq + 1
	}
	return profDer + 1
}

// ptr **nodoArbol[T] -> es un puntero que apunta a otro puntero
func agregarNodo[T cmp.Ordered](ptr **nodoArbol[T], elem T) bool {
	if *ptr == nil {
		*ptr = &nodoArbol[T]{elem: elem}
		return true
	}
	// 
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

// CAMBIO: fp recibe T en lugar de int
func (this ArbolBin[T]) Traverse(fp func(T), o Order) {
	traverseRec(this.raiz, fp, o)
}

func traverseRec[T cmp.Ordered](n *nodoArbol[T], fp func(T), o Order) {
	if n == nil {
		return
	}
	switch o {
	case PreOrder:
		fp(n.elem)
		traverseRec(n.HI, fp, o)
		traverseRec(n.HD, fp, o)
	case InOrder:
		traverseRec(n.HI, fp, o)
		fp(n.elem)
		traverseRec(n.HD, fp, o)
	case PostOrder:
		traverseRec(n.HI, fp, o)
		traverseRec(n.HD, fp, o)
		fp(n.elem)
	}
}

func (this ArbolBin[T]) Apply(fp func(T) T) {
	applyRec(this.raiz, fp)
}

// Recorre recrusivamente en PreOrden
func applyRec[T cmp.Ordered](n *nodoArbol[T], fp func(T) T) {
	if n == nil {
		return
	}
	n.elem = fp(n.elem)
	applyRec(n.HI, fp)
	applyRec(n.HD, fp)
}

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

func (this ArbolBin[T]) Find(fp func(T) bool) bool {
	if this.IsEmpty() {
		return false
	}
	if fp(this.raiz.elem) {
		return true
	}
	return this.GetLeft().Find(fp) || this.GetRight().Find(fp)
}

// CAMBIO: %v en lugar de %d para que funcione con cualquier tipo ordenable
func (this ArbolBin[T]) String() string {
	var elems []string
	this.Traverse(func(v T) {
		elems = append(elems, fmt.Sprintf("%v", v))
	}, InOrder)
	return "[" + strings.Join(elems, " ") + "]"
}
