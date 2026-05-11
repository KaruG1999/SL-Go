package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// invertirPalabra devuelve la palabra con sus caracteres en orden inverso.
// Usa []rune para que caracteres Unicode multi-byte (é, ñ, etc.) se
// inviertan como una unidad y no byte a byte (lo que produciría UTF-8 inválido).

func invertirPalabra(s string) string {
	runes := []rune(s)
	// Intercambio desde los extremos hacia el centro (algoritmo in-place).
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Print("Ingrese una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text() // Extrae el contenido del buffer como un string

	// strings.Fields divide la frase por cualquier secuencia de espacios en
	// blanco y devuelve un slice de strings con cada palabra
	palabras := strings.Fields(frase)

	for i := range palabras {
		// Las posiciones se cuentan desde 1, por eso chequeamos (i+1)%2 == 1.
		// Posiciones impares: 1, 3, 5, … → índices 0, 2, 4, …
		if (i+1)%2 == 1 {
			palabras[i] = invertirPalabra(palabras[i])
		}
	}

	// strings.Join vuelve a unir el slice con un espacio entre cada elemento.
	fmt.Println(strings.Join(palabras, " "))
}
