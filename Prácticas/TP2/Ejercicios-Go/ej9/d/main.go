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
	return List{}
}

// Todos los accesos son métodos: la lista es el receptor, no un argumento.

func (l List) IsEmpty() bool {
	return l.len == 0
}

func (l List) Len() int {
	return l.len
}

func (l List) FrontElement() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("lista vacía")
	}
	return l.pri.elem, nil
}

func (l List) Next() (List, error) {
	if l.IsEmpty() {
		return List{}, errors.New("lista vacía")
	}
	return List{pri: l.pri.sig, len: l.len - 1}, nil
}

// Implementar String() hace que fmt.Printf("%s", l) lo llame automáticamente.
func (l List) String() string {
	res := ""
	for actual := l.pri; actual != nil; actual = actual.sig {
		res += fmt.Sprintf("[%d] -> ", actual.elem)
	}
	return res
}

func (l *List) PushFront(elem int) {
	l.pri = &nodo{elem: elem, sig: l.pri}
	l.len++
}

func (l *List) PushBack(elem int) {
	if l.IsEmpty() {
		l.PushFront(elem)
		return
	}
	actual := l.pri
	for actual.sig != nil {
		actual = actual.sig
	}
	actual.sig = &nodo{elem: elem}
	l.len++
}

func (l *List) Remove() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("lista vacía")
	}
	valor := l.pri.elem
	l.pri = l.pri.sig
	l.len--
	return valor, nil
}

func (l *List) Iterate(f func(int) int) {
	for actual := l.pri; actual != nil; actual = actual.sig {
		actual.elem = f(actual.elem)
	}
}

func main() {
	fmt.Println("=== Creación y carga ===")
	l := New()
	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(1)
	fmt.Printf("Lista: %s\n", l) // llama a l.String() automáticamente
	fmt.Printf("Longitud: %d\n", l.Len())
	fmt.Printf("¿Vacía?: %v\n", l.IsEmpty())

	elem, err := l.FrontElement()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Primer elemento: %d\n", elem)
	}

	fmt.Println("\n=== Recorrido con Next ===")
	for cursor := l; !cursor.IsEmpty(); {
		v, _ := cursor.FrontElement()
		fmt.Printf("  Elemento: %d\n", v)
		cursor, _ = cursor.Next()
	}

	fmt.Println("\n=== Iterate: restar 1 ===")
	l.Iterate(func(n int) int { return n - 1 })
	fmt.Printf("Lista: %s\n", l)

	fmt.Println("\n=== Remove ===")
	for !l.IsEmpty() {
		val, err := l.Remove()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Removido: %d | Lista restante: %s\n", val, l)
		}
	}

	fmt.Println("\n=== Errores en lista vacía ===")
	_, err = l.Remove()
	fmt.Println("Remove:", err)

	_, err = l.FrontElement()
	fmt.Println("FrontElement:", err)

	_, err = l.Next()
	fmt.Println("Next:", err)
}
