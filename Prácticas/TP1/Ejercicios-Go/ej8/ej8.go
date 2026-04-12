package main

import "fmt"

func main() {

	// =========================================================
	// PASO 1: Leer la dirección ingresada por el usuario
	// fmt.Scan lee una palabra (hasta el primer espacio o newline).
	// =========================================================

	var direccion string
	fmt.Print("Ingrese la dirección del viento (N / S / E / O): ")
	fmt.Scan(&direccion)

	// =========================================================
	// PASO 2: Determinar hacia dónde sopla el viento con switch
	// El viento "viene de" la dirección indicada, así que
	// "se dirige hacia" la dirección opuesta.
	//
	// Aceptamos mayúscula y minúscula en cada case usando
	// múltiples valores separados por coma — no hace falta
	// convertir el string antes de comparar.
	//
	// default captura cualquier entrada no válida, equivale
	// al else final de una cadena if/else if.
	// =========================================================

	switch direccion {
	case "N", "n":
		fmt.Println("El viento se dirige hacia el Sur")
	case "S", "s":
		fmt.Println("El viento se dirige hacia el Norte")
	case "E", "e":
		fmt.Println("El viento se dirige hacia el Oeste")
	case "O", "o":
		fmt.Println("El viento se dirige hacia el Este")
	default:
		fmt.Println("Dirección no válida. Ingrese N, S, E u O.")
	}
}
