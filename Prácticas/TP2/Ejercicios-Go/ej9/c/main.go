package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(" container/list: operaciones equivalentes ")
	l := list.New()

	// Equivalente a PushBack / PushFront
	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(1)

	fmt.Printf("Longitud: %d\n", l.Len())

	// Equivalente a IsEmpty
	fmt.Printf("¿Vacía?: %v\n", l.Len() == 0)

	// Equivalente a FrontElement
	fmt.Printf("Primer elemento: %v\n", l.Front().Value)

	// Equivalente a ToString
	fmt.Print("Lista: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("[%v] -> ", e.Value)
	}
	fmt.Println()

	// Recorrido hacia atrás: posible porque es doblemente enlazada
	// (nuestra implementación no puede hacer esto)
	fmt.Print("Lista inversa: ")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("[%v] -> ", e.Value)
	}
	fmt.Println()

	// Remove: saca un *Element específico, no necesariamente el primero
	// Nuestra Remove() solo sacaba el primero
	front := l.Front()
	removed := l.Remove(front)
	fmt.Printf("\nRemovido: %v | Longitud restante: %d\n", removed, l.Len())

	// container/list almacena 'any', así que acepta tipos mixtos
	// Nuestra lista solo aceptaba int
	fmt.Println("\n=== Lista de tipos mixtos (solo posible con container/list) ===")
	mixed := list.New()
	mixed.PushBack(42)
	mixed.PushBack("hola")
	mixed.PushBack(3.14)
	for e := mixed.Front(); e != nil; e = e.Next() {
		fmt.Printf("  valor: %v  tipo: %T\n", e.Value, e.Value)
	}

	// Para hacerlo genérico con nuestra propia implementación,
	// a partir de Go 1.18 se puede usar parámetros de tipo:
	//
	//   type List[T any] struct { pri *nodo[T]; len int }
	//   type nodo[T any] struct { elem T; sig *nodo[T] }
	//
	// Así List[int], List[string], List[float64] etc. son tipos distintos
	// y el compilador verifica los tipos en lugar de usar 'any'.
}
