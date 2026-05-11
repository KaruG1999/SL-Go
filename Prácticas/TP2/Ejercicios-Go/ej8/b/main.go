package main

import "fmt"

const digitos = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Convert(v int, b int) string {
	// TODO: igual que base pero contemplando v negativo
	// Tip: si v < 0, trabajá con -v y luego agregá "-" al resultado final
	panic("no implementado")
}

func main() {
	fmt.Println(Convert(23, 2))   // "10111"
	fmt.Println(Convert(-10, 2))  // "-1010"
	fmt.Println(Convert(0, 10))   // "0"
}
