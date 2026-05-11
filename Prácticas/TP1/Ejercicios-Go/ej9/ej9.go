package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// reemplazarConCasing busca todas las ocurrencias de `original` dentro de
// `frase` de forma case-insensitive y las reemplaza por `reemplazo`,
// copiando el patrón de mayúsculas/minúsculas de la ocurrencia encontrada
// posición a posición.
//
// Se trabaja con []rune (y no con bytes) porque en Go un string es una
// secuencia de bytes UTF-8; caracteres como é, ñ, ó ocupan 2 bytes, por lo
// que iterar con índice de bytes daría posiciones incorrectas.
// Al convertir a []rune cada elemento representa un carácter completo.
//
// En este ejercicio ambas palabras son ASCII puro ("jueves"/"martes"),
// así que rune y byte coinciden; pero se usa []rune desde ya para que la
// función sea reutilizable con palabras acentuadas (ver Obligatorio 1).
func reemplazarConCasing(frase, original, reemplazo string) string {
	fraseRunes := []rune(frase)
	origRunes := []rune(strings.ToLower(original))
	reemplRunes := []rune(strings.ToLower(reemplazo))
	n := len(origRunes)

	var sb strings.Builder // Builder acumula el resultado sin alocar un string nuevo en cada paso

	i := 0
	for i < len(fraseRunes) {
		// ¿Hay suficientes caracteres por delante para que quepa `original`?
		if i+n <= len(fraseRunes) {
			// Comparamos el segmento en minúsculas con la palabra buscada.
			segmento := strings.ToLower(string(fraseRunes[i : i+n]))
			if segmento == string(origRunes) {
				// Ocurrencia encontrada: escribir `reemplazo` con el casing
				// de cada posición de la ocurrencia original.
				for j, r := range reemplRunes {
					if unicode.IsUpper(fraseRunes[i+j]) {
						sb.WriteRune(unicode.ToUpper(r))
					} else {
						sb.WriteRune(r)
					}
				}
				i += n
				continue
			}
		}
		// Carácter normal: copiarlo tal cual.
		sb.WriteRune(fraseRunes[i])
		i++
	}
	return sb.String()
}

func main() {
	// bufio.Scanner lee la línea completa (incluyendo espacios).
	// fmt.Scan solo lee hasta el primer espacio, lo que partiría la frase.
	fmt.Print("Ingrese una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println(reemplazarConCasing(frase, "jueves", "martes"))
}
