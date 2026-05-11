package main

import "fmt"

const digitos = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Convert(v int, b int) string {
	// TODO: convertir v (positivo) a su representación en base b
	// Tip 1: obtené los dígitos de derecha a izquierda con v%b, luego v/b
	// Tip 2: string(digitos[v%b]) convierte el índice a su carácter correspondiente
	// Tip 3: preponer el dígito al resultado: result = char + result
	panic("no implementado")
}

func main() {
	fmt.Println(Convert(23, 2))   // "10111"
	fmt.Println(Convert(255, 16)) // "FF"
	fmt.Println(Convert(0, 10))   // "0"
}
