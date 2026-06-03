package main

import (
	"fmt"
	ibt "practicas/TP3/Ejercicios-Go/ej3/tree"
)

func main() {
	// Uso 1: árbol de enteros (igual que TP2)
	a := ibt.New[int]()
	for _, v := range []int{5, 3, 7, 1, 4, 6, 8} {
		a = a.Add(v)
	}
	fmt.Printf("InOrder (int):  %s\n", a.String())
	fmt.Printf("Includes(4): %v | Includes(9): %v\n", a.Includes(4), a.Includes(9))

	// Uso 2: árbol de strings — imposible en TP2 sin reescribir todo
	b := ibt.New[string]()
	for _, v := range []string{"mango", "banana", "pera", "anana"} {
		b = b.Add(v)
	}
	fmt.Printf("InOrder (string): %s\n", b.String())
}
