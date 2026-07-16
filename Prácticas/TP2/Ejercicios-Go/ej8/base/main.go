package main

import "fmt"

const digitos = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Convert(v int, b int) string {
	// caso si dato es 0
	if v==0 {return "0"}
	// caso si b es inválido
	if b<2 || b>36 {return "Base inválida"}
	// en string vamos acumulando el resultado
	resultado := ""
	for v>0 {
		resto := v % b
		digito := digitos[resto]
		// Luego sumo el digito al resultado ADELANTE
		resultado = string(digito) + resultado
		// Sigo con el resto del valor
		v = v/b
	}
	return resultado

	// Con strings.Builder no alcanza con cambiar el tipo: Builder es
	// eficiente para appendear, pero acá estamos prependeando en cada
	// vuelta. Para que sirva hay que appendear los dígitos en el orden
	// que salen (unidades primero) y después invertir el string una vez:
	//
	// var sb strings.Builder
	// for v > 0 {
	//     sb.WriteByte(digitos[v%b])
	//     v = v / b
	// }
	// runes := []rune(sb.String())
	// for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	//     runes[i], runes[j] = runes[j], runes[i]
	// }
	// return string(runes)
}

func main() {
	fmt.Println(Convert(23, 2))   // "10111"
	fmt.Println(Convert(255, 16)) // "FF"
	fmt.Println(Convert(0, 10))   // "0"
}
