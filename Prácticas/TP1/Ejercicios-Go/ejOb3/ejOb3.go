package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// invertirMayusculasYMinusculas invierte el casing de una palabra usando un Builder de forma eficiente.
func invertirMayusculasYMinusculas(palabra string) string {
	var sb strings.Builder
	for _, letra := range palabra {
		if unicode.IsUpper(letra) {
			sb.WriteRune(unicode.ToLower(letra))
		} else {
			sb.WriteRune(unicode.ToUpper(letra))
		}
	}
	return sb.String()
}

// cambiarFrase busca coincidencias case-insensitive y las reemplaza con el casing invertido.
func cambiarFrase(frase string, palabra string) string {
	fraseRunes := []rune(frase)
	palabraRunes := []rune(palabra)
	
	lenPalabra := len(palabraRunes)
	var sb strings.Builder
	i := 0

	for i < len(fraseRunes) {
		// Validamos si la palabra buscada entra en lo que queda de la frase
		if i+lenPalabra <= len(fraseRunes) {
			subSegmento := string(fraseRunes[i : i+lenPalabra])

			// strings.EqualFold realiza comparación de strings insensible a mayúsculas/minúsculas (case-insensitive)
			if strings.EqualFold(subSegmento, palabra) {
				sb.WriteString(invertirMayusculasYMinusculas(subSegmento))
				i += lenPalabra
				continue
			}
		}
		// No hubo coincidencia: copiamos el carácter de la frase tal cual
		sb.WriteRune(fraseRunes[i])
		i++
	}

	return sb.String()
}

func main() {
	// Correción profes ->  Validación obligatoria de argumentos de entrada
	if len(os.Args) < 2 {
		fmt.Println("Error: Debe ingresar la palabra a buscar como argumento al ejecutar el programa.")
		fmt.Println("Uso: go run ejOb3.go <palabra>")
		return
	}
	
	palabra := os.Args[1]

	fmt.Print("Ingrese una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println("Resultado:", cambiarFrase(frase, palabra))
}