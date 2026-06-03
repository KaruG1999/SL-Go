package main

import (
	"fmt"
	"practicas/TP3/Ejercicios-Go/ej2/list"
)

type Persona struct{
	Nombre string
	Edad int
}

func main() {
	listaEnteros := list.New[int]()
	listaEnteros.PushBack(10)
	listaEnteros.PushBack(23)
	listaEnteros.PushFront(5)

	fmt.Println(" Lista de enteros: ", listaEnteros.String())

	// Busqueda 
	num, encontrado := listaEnteros.Find(func(i int) bool {
		return i>20
	})
	fmt.Printf("¿Mayor a 20 encontrado?: %v, Valor: %d\n\n", encontrado, num)

	// Ejemplo Lista de Estructuras Complejas
	listaPersonas := list.New[Persona]()
	listaPersonas.PushBack(Persona{"Karen", 24})
	listaPersonas.PushBack(Persona{"Alejandro", 30})

	fmt.Println("Lista de personas:", listaPersonas.String())

	// Búsqueda en estructuras: buscamos por el campo Nombre
	persona, encontrada := listaPersonas.Find(func(p Persona) bool {
		return p.Nombre == "Karen"
	})
	fmt.Printf("¿Se encontró a Karen?: %v, Edad: %d\n", encontrada, persona.Edad)
	
}
