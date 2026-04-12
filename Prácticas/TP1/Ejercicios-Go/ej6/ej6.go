package main

import (
	"errors"
	"fmt"
)

// 1. Definimos la función. Retorna el resultado y un posible error.
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// Retornamos un error claro si el divisor es cero
		return 0, errors.New("no se puede dividir por cero")
	}
	// Si b no es cero, retornamos el resultado y "nil" (que significa "sin error")
	return a / b, nil
}

func main() {
	var n1, n2 float64

	fmt.Print("Número 1: ")
	fmt.Scan(&n1)
	fmt.Print("Número 2: ")
	fmt.Scan(&n2)

	// 2. Recibimos los dos valores que devuelve la función
	resultado, err := dividir(n1, n2)

	// 3. Primero preguntamos: ¿Hubo un error?
	if err != nil {
		fmt.Println("Error detectado:", err)
	} else {
		// Si err es nil, todo salió bien
		fmt.Println("Resultado:", resultado)
	}
}
