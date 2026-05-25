package miarbol

import (
	"fmt"
	"strings"
)

type nodoArbol struct {
	elem int
	HI   *nodoArbol
	HD   *nodoArbol
}

type arbolBin struct {
	raiz *nodoArbol
	len  int
}

type Order int

const (
	PreOrder  Order = iota
	InOrder
	PostOrder
)

func New() arbolBin {
	return arbolBin{}
}

func (this arbolBin) IsEmpty() bool {
	return this.raiz == nil
}

func (this arbolBin) GetElem() int {
	return this.raiz.elem
}

func (this arbolBin) GetRight() arbolBin {
	return arbolBin{raiz: this.raiz.HD}
}

func (this arbolBin) GetLeft() arbolBin {
	return arbolBin{raiz: this.raiz.HI}
}

func (this arbolBin) Len() int {
	return this.len
}

func (this arbolBin) Depth() int {
	return calcularProf(this.raiz)
}

func calcularProf(n *nodoArbol) int {
	if n == nil {
		return 0
	}
	profIzq := calcularProf(n.HI)
	profDer := calcularProf(n.HD)
	if profIzq > profDer {
		return profIzq + 1
	}
	return profDer + 1
}

// Usa doble puntero para modificar el puntero del nodo padre directamente
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
	return false // ya existe
}

func (this arbolBin) Add(elem int) arbolBin {
	if agregarNodo(&this.raiz, elem) {
		this.len++
	}
	return this
}

func (this arbolBin) Traverse(fp func(int), o Order) {
	traverseRec(this.raiz, fp, o)
}

func traverseRec(n *nodoArbol, fp func(int), o Order) {
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

func (this arbolBin) Apply(fp func(int) int) {
	applyRec(this.raiz, fp)
}

func applyRec(n *nodoArbol, fp func(int) int) {
	if n == nil {
		return
	}
	n.elem = fp(n.elem)
	applyRec(n.HI, fp)
	applyRec(n.HD, fp)
}

// Aprovecha el orden del ABB para búsqueda O(log n)
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

func (this arbolBin) Find(fp func(int) bool) bool {
	if this.IsEmpty() {
		return false
	}
	if fp(this.raiz.elem) {
		return true
	}
	return this.GetLeft().Find(fp) || this.GetRight().Find(fp)
}

func (this arbolBin) String() string {
	var elems []string
	this.Traverse(func(v int) {
		elems = append(elems, fmt.Sprintf("%d", v))
	}, InOrder)
	return "[" + strings.Join(elems, " ") + "]"
}
