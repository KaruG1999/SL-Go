package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*Realice un programa que reciba una palabra como argumento y lee de la entrada una frase. Luego, el programa debe
imprimir la frase que leyó con cada una de las ocurrencias de la palabra con las mayúsculas y minúsculas invertidas.
Por ejemplo, si la frase es: "Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO"  y la palabra es "PEQUEÑO"
entonces el programa imprimirá: "Parece PEQueÑO, pero no es tan PEQUEñO el pequeño"  Tenga en cuenta que la palabra
a buscar puede ser ingresada con mayúsculas y minúsculas mezcladas. */

func invertirMayusculasYMinusculas(palabra string) string {
	resultado := ""
	for _, letra := range palabra {
		if unicode.IsUpper(letra) {
			resultado += string(unicode.ToLower(letra))
		} else {
			resultado += string(unicode.ToUpper(letra))
		}
	}
	return resultado
}

func cambiarFrase(frase string, palabra string) string {
	
	// Convertimos a []rune para trabajar carácter a carácter con Unicode, no con bytes (Slice de Runas, no runa)
	runes := []rune(frase)

	// Guardamos la longitud en runes (no en bytes) de la palabra buscada.
	// Por ejemplo "pequeño" tiene 8 runes pero 9 bytes, por eso no usamos len(palabra).
	n := len([]rune(palabra))

	resultado := ""
	i := 0

	for i < len(runes) {
		// Si quedan suficientes caracteres por delante, tomamos un segmento del
		// tamaño de la palabra buscada y comparamos sin importar mayúsculas/minúsculas
		if i+n <= len(runes) {

			sub := string(runes[i : i+n]) // Convertimos el segmento a string para compararlo con la palabra buscada


			if strings.EqualFold(sub, palabra) {
				// Encontramos una ocurrencia: la agregamos con el casing invertido
				resultado += invertirMayusculasYMinusculas(sub)
				i += n
				continue // Saltamos el segmento que acabamos de procesar
			}
		}
		// No hubo coincidencia: copiamos el carácter tal cual y avanzamos
		resultado += string(runes[i])
		i++
	}

	return resultado
}

func main() {
	palabra := os.Args[1] 

	fmt.Println("Ingrese una frase")
	// bufio.Scanner lee la línea completa incluyendo espacios
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println(cambiarFrase(frase, palabra))
}


// Dif os.Args[0] es el nombre del programa, Args[1] es el primer argumento que se le pasa al programa, en este caso la palabra a buscar. Si no se le pasa ningún argumento, el programa fallará con un error de índice fuera de rango. Por eso es importante asegurarse de que se le pase la palabra al ejecutar el programa.
// Scanner lee durante la ejecución del programa, y cuando el usuario ingresa una línea de texto y presiona Enter, esa línea se almacena en scanner.Text() para ser procesada posteriormente. En este caso, se utiliza para obtener la frase que el usuario desea modificar.
