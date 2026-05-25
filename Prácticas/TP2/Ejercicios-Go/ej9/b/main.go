package main

import (
	"errors"
	"fmt"
)

type nodo struct {
	elem int
	sig  *nodo
}

type List struct {
	pri *nodo
	len int
}

func New() List {
	return List{nil, 0}
}

func IsEmpty(l List) bool {
	return l.pri == nil && l.len == 0
}

func Len(l List) int {
	return l.len
}

func FrontElement(l List) int {
	return l.pri.elem
}

func Next(l List) List {
	if l.pri == nil || l.pri.sig == nil {
		return New()
	}
	return List{
		pri: l.pri.sig,
		len: l.len - 1,
	}
}

func ToString(l List) string {
	res := ""
	actual := l.pri
	for actual != nil {
		res += fmt.Sprintf("[%d] -> ", actual.elem)
		actual = actual.sig
	}
	return res
}

func (l *List) PushFront(elem int) {
	nuevo := &nodo{elem: elem, sig: l.pri}
	l.pri = nuevo
	l.len++
}

func (l *List) PushBack(element int) {
	if IsEmpty(*l) {
		l.PushFront(element)
		return
	}
	nuevo := &nodo{elem: element, sig: nil}
	actual := l.pri
	for actual.sig != nil {
		actual = actual.sig
	}
	actual.sig = nuevo
	l.len++
}

func (l *List) Remove() (int, error) {
	if l.pri == nil {
		return 0, errors.New("lista vacia")
	}
	valor := l.pri.elem
	l.pri = l.pri.sig
	l.len--
	return valor, nil
}

func (l *List) Iterate(f func(int) int) {
	actual := l.pri
	for actual != nil {
		actual.elem = f(actual.elem)
		actual = actual.sig
	}
}

func main() {
	fmt.Println("=== Creación y carga ===")
	l := New()
	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(1)
	fmt.Printf("Lista: %s\n", ToString(l))
	fmt.Printf("Longitud: %d\n", Len(l))
	fmt.Printf("¿Vacía?: %v\n", IsEmpty(l))
	fmt.Printf("Primer elemento: %d\n", FrontElement(l))

	fmt.Println("\n=== Recorrido con Next ===")
	for cursor := l; !IsEmpty(cursor); cursor = Next(cursor) {
		fmt.Printf("  Elemento: %d\n", FrontElement(cursor))
	}

	fmt.Println("\n=== Iterate: multiplicar por 2 ===")
	l.Iterate(func(n int) int { return n * 2 })
	fmt.Printf("Lista: %s\n", ToString(l))

	fmt.Println("\n=== Remove ===")
	for !IsEmpty(l) {
		val, err := l.Remove()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Removido: %d | Lista restante: %s\n", val, ToString(l))
		}
	}

	fmt.Println("\n=== Remove en lista vacía ===")
	_, err := l.Remove()
	fmt.Println("Error esperado:", err)
}
