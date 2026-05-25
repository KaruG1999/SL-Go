package main

import (
	"errors"
	"fmt"
)

// --- Lista enlazada (de ej9) ---

type nodo struct {
	elem int
	sig  *nodo
}

type List struct {
	pri *nodo
	len int
}

func newList() List {
	return List{}
}

func (l List) IsEmpty() bool {
	return l.len == 0
}

func (l *List) PushFront(elem int) {
	l.pri = &nodo{elem: elem, sig: l.pri}
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

func (l List) FrontElement() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("lista vacía")
	}
	return l.pri.elem, nil
}

func (l *List) Iterate(f func(int) int) {
	for actual := l.pri; actual != nil; actual = actual.sig {
		actual.elem = f(actual.elem)
	}
}

func (l List) String() string {
	res := ""
	for actual := l.pri; actual != nil; actual = actual.sig {
		res += fmt.Sprintf("[%d] -> ", actual.elem)
	}
	return res
}

// --- Stack construido sobre List ---
// Push  = PushFront  →  el tope es siempre el primer nodo
// Pop   = Remove     →  saca el primer nodo
// Top   = FrontElement

type Stack struct {
	data List
}

func New() Stack {
	return Stack{data: newList()}
}

func (s Stack) IsEmpty() bool {
	return s.data.IsEmpty()
}

func (s Stack) Len() int {
	return s.data.len
}

func (s Stack) Top() (int, error) {
	return s.data.FrontElement()
}

func (s Stack) String() string {
	return "tope -> " + s.data.String()
}

func (s *Stack) Push(element int) {
	s.data.PushFront(element)
}

func (s *Stack) Pop() (int, error) {
	return s.data.Remove()
}

func (s *Stack) Iterate(f func(int) int) {
	s.data.Iterate(f)
}

func main() {
	fmt.Println("=== Creación y Push ===")
	s := New()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("Pila: %s\n", s)
	fmt.Printf("Longitud: %d\n", s.Len())
	fmt.Printf("¿Vacía?: %v\n", s.IsEmpty())

	top, err := s.Top()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Tope (sin sacar): %d\n", top)
	}

	fmt.Println("\n=== Iterate: duplicar ===")
	s.Iterate(func(n int) int { return n * 2 })
	fmt.Printf("Pila: %s\n", s)

	fmt.Println("\n=== Pop ===")
	for !s.IsEmpty() {
		val, err := s.Pop()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Sacado: %d | Restante: %s\n", val, s)
		}
	}

	fmt.Println("\n=== Errores en pila vacía ===")
	_, err = s.Pop()
	fmt.Println("Pop:", err)
	_, err = s.Top()
	fmt.Println("Top:", err)
}
