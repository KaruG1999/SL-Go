package main

import "fmt"

func main() {
	/*  imprima en la salida estándar la suma
	de los primeros números positivos pares menores o iguales a
	250 */
	var sum int = 0
	for i := 0; i <= 250; i += 2 {
		sum += i
	}
	fmt.Println("La suma de los primeros números positivos pares menores o iguales a 250 es:", sum)

	/* Cambiar el programa para que itere en el sentido contrario
	   pero obtener el mismo resultado */
	sum = 0
	for i := 250; i >= 0; i -= 2 {
		sum += i
	}
	fmt.Println("La suma de los primeros números positivos pares menores o iguales a 250 es:", sum)

	/* Cambiar el programa para que
	   en lugar de usar un literal como tope se use una constante */
	const limit int = 250
	sum = 0
	for i := 0; i <= limit; i += 2 {
		sum += i
	}
	fmt.Println("La suma de los primeros números positivos pares menores o iguales a", limit, "es:", sum)

}
