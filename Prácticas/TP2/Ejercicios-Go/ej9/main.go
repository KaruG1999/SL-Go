package main

import (
	"fmt"
	"errors"
)

type nodo struct {
	elem int
	sig *nodo
}

type List struct{
	pri *nodo
	len int 
}

func New() List{
	return List{nil, 0}
}

func IsEmpty(l List) bool{
	return l.pri==nil && l.len ==0
}

func Len (l List)int {
	return l.len
}

func FrontElement (l List) int {
	return l.pri.elem
}

func Next (l List) List {
	if (l.pri == nil || l.pri.sig == nil){
		return New()
	}

	return List{
		pri : l.pri.sig,
		len : l.len - 1,
	}
}

func ToString (l List) string {
	res := ""
	// Para no modificar lista original
	actual := l.pri
	for actual != nil {
		res += fmt.Sprintf("[%d] -> ", actual.elem)
		actual = actual.sig
	}
	return res
}

// Pasado a método para unificar interfaz (creo el nodo y lo asigno al primer elemento de la lista)
func (l *List) PushFront(elem int) {
	nuevo := &nodo{elem: elem, sig: l.pri}
	l.pri = nuevo
	l.len++
}

// Otra forma de escribirlo
func (l *List) PushBack( element int) {
    if IsEmpty(*l) {
        // 'l' ya es un puntero (*List). Se lo pasás tal cual.
        l.PushFront(element)
        return 
    }

    nuevo := &nodo{elem: element, sig: nil} // -> nuevo es el último
    actual := l.pri // -> actual es el primero y recorro hasta el último (de esta forma no pierdo el puntero)
    for actual.sig != nil {
        actual = actual.sig
    }
    actual.sig = nuevo // -> llegue al final y le "engancho" el nuevo
    l.len++
}

func (l *List) Remove() (int, error) {
	// si está vacía, salimos con error 
	if l.pri == nil {
		return 0, errors.New("Lista vacia")
	}

	// Guardamos el valor del nodo que vamos a sacar
	valor := l.pri.elem

	// Desplazamos el puntero inicial al siguiente nodo (Elimina el primero)
	l.pri = l.pri.sig

	// Decrementamos el contador de tamaño
	l.len--

	return valor, nil // nil significa "sin error"
}

func (l *List) Iterate(f func(int) int) {
	actual := l.pri
	
	// Recorremos la memoria dinámica hasta el final
	for actual != nil {
		actual.elem = f(actual.elem)
		// Avanzamos al siguiente nodo
		actual = actual.sig
	}

}

func main() {

	l:= New()
	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(1)

	fmt.Printf("Lista: %s\n", ToString(l))

	l.Iterate(func(n int) int{
		return n-5
	})

	fmt.Printf("Lista iterada: %s\n", ToString(l))
}





