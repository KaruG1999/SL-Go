package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// reemplazarConCasing busca 'original' en 'frase' y la cambia por 'reemplazo' manteniendo el casing.
func reemplazarConCasing(frase, original, reemplazo string) string {
	// convertimos a slice de runas (evitamos que se rompa al usar letras con acento)
	fraseRunes := []rune(frase)
	origRunes := []rune(strings.ToLower(original))
	reemplRunes := []rune(strings.ToLower(reemplazo))
	
	lenOrig := len(origRunes)   // sin runa el len retornaría un byte extra si la frase tiene tilde 
	var sb strings.Builder      // acumulador de texto ultra eficiente -> strings.Builder escribe directamente sobre un buffer de memoria

	i := 0
	for i < len(fraseRunes) {
		// Validar si entra la palabra original en lo que queda de la frase
		if i+lenOrig <= len(fraseRunes) {
			segmento := strings.ToLower(string(fraseRunes[i : i+lenOrig]))
			
			if segmento == string(origRunes) {

				// procesa el reemplazo letra por letra (j -> indice y r -> letra)
				for j, r := range reemplRunes {
					// Evitamos el index out of range si el reemplazo es más largo
					// Si j supera el tamaño original, evalúa el casing de la última letra del match
					idxEvaluar := i + j
					if j >= lenOrig {
						idxEvaluar = i + lenOrig - 1 // Para en el índice de la última letra de la palabra original
					}

					// Acá compara con frase original conirtiendo mayusc y minusc
					if unicode.IsUpper(fraseRunes[idxEvaluar]) {
						sb.WriteRune(unicode.ToUpper(r))
					} else {
						sb.WriteRune(r)
					} 
				}
				i += lenOrig // Avanzar el tamaño de la palabra encontrada
				continue
			}
		}
		// Avanzar un caracter normal si no hubo coincidencia
		sb.WriteRune(fraseRunes[i])
		i++
	}
	return sb.String()
}

func main() {
	fmt.Print("Ingrese una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	// Mismo código ej 9 -> strings con tildes de 2 bytes 
	fmt.Println("Resultado Obligatorio 1:", reemplazarConCasing(frase, "miércoles", "automóvil"))
}