package list

import "fmt"

// Tipos génerico de lista
type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	ini, fin *element[T]
	len      int
}

func New[T any]() *List[T] {
	return &List[T]{ini: nil, fin: nil, len: 0}
}

func (l *List[T]) IsEmpty() bool {
	return l.ini == nil
}

func (l *List[T]) Len() int {
	return l.len
}

// Agrego al final con Append
func (l *List[T]) PushBack(val T) {
	nuevo := &element[T]{val: val, next: nil} // creo el nodo a insertar
	if l.IsEmpty() {
		l.ini = nuevo
		l.fin = nuevo
	} else {
		// ?
		l.fin.next = nuevo
		l.fin = nuevo
	}
	l.len++
}

// Agrego al inicio con Prepend
func (l *List[T]) PushFront(val T) {
	nuevo := &element[T]{val: val, next: l.ini}
	if l.IsEmpty() {
		l.fin = nuevo
	}
	l.ini = nuevo
	l.len++
}

func (l *List[T]) Find(criterio func(T) bool) (T, bool) {
	actual := l.ini
	for actual != nil {
		// Si existe retorna valor y boolean
		if criterio(actual.val) {
			return actual.val, true
		}
		actual = actual.next
	}
	var valorVacio T // retorna valor por defecto
	return valorVacio, false
}

func (l *List[T]) String() string {
	if l.IsEmpty() {
		return "[]"
	}
	var resultado string
	actual := l.ini
	for actual != nil {
		resultado += fmt.Sprintf("%v -> ", actual.val)
		actual = actual.next
	}
	return resultado + "nil"
}
