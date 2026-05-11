package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	
)


func main () {
	fmt.Println("Ingrese una secuencia de caracteres (Finalizar con Enter): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text() // Leemos hasta que encontrems el salto de línea (\n)

	// Contadores de distintos caracteres
	var letras, numeros, especiales int


	// Recorremos el text letra por letra
	for _,caracter := range texto {
		// No tenemos que contar el Enter (CR o LF)
		if caracter == '\n' || caracter == '\r' {
			continue
		}
		// paquete unicode se usa para clasificar 
		if unicode.IsLetter(caracter) {
			letras++
		} else if unicode.IsDigit(caracter){
			numeros++
		} else {especiales++}

	}

	fmt.Printf(" Letras: %d, Números: %d, Caracteres: %d", letras, numeros, especiales)

}
