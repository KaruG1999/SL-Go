package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// reemplazarConCasing es la misma función del Ejercicio 9, reutilizable
// para cualquier par de palabras.
//
// Impacto de pasar de "jueves"/"martes" a "miércoles"/"automóvil":
//   - Ambas palabras tienen 9 runes ("miércoles": m-i-é-r-c-o-l-e-s,
//     "automóvil": a-u-t-o-m-ó-v-i-l), así que el reemplazo posición
//     a posición sigue siendo válido.
//   - La diferencia es que ahora los strings tienen bytes extra por las
//     tildes (é = 2 bytes, ó = 2 bytes). Si usáramos índice de bytes,
//     len("miércoles") == 11 (bytes) pero tiene 9 caracteres reales.
//     Por eso es imprescindible trabajar con []rune.
func reemplazarConCasing(frase, original, reemplazo string) string {
	fraseRunes := []rune(frase)
	origRunes := []rune(strings.ToLower(original))
	reemplRunes := []rune(strings.ToLower(reemplazo))
	n := len(origRunes)

	var sb strings.Builder

	i := 0
	for i < len(fraseRunes) {
		if i+n <= len(fraseRunes) {
			segmento := strings.ToLower(string(fraseRunes[i : i+n]))
			if segmento == string(origRunes) {
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

	fmt.Println(reemplazarConCasing(frase, "miércoles", "automóvil"))
}
