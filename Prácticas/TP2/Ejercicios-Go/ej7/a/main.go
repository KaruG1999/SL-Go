package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	
	fmt.Print("Ingrese una secuencia de caracteres: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	var mayusculas, minusculas, numeros, especiales int

	// Acá range procesa el texto en runas 
	for _, c := range texto {
		switch {
		case unicode.IsUpper(c):
			mayusculas++
		case unicode.IsLower(c):
			minusculas++
		case unicode.IsDigit(c):
			numeros++
		default:
			especiales++
		}
	}

	fmt.Printf("Mayúsculas: %d, Minúsculas: %d, Números: %d, Especiales: %d\n",
		mayusculas, minusculas, numeros, especiales)
}
