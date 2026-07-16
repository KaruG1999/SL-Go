package main

import "fmt"

func main() {

	// fmt.Scan lee una palabra (hasta el primer espacio o newline)
	var direccion string
	fmt.Print("Ingrese la dirección del viento (N / S / E / O): ")
	fmt.Scan(&direccion)

	// Aceptamos mayúscula y minúscula en cada case usando
	// múltiples valores separados por coma — no hace falta convertir el string antes de comparar.

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