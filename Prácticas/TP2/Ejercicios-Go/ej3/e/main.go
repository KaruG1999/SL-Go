package main

import "fmt"

type PuntoCardinal int

const (
	N  PuntoCardinal = iota
	S
	E
	O
	NE
	SO
	NO
	SE
)

func (p PuntoCardinal) String() string {
	// TODO: retornar el nombre textual del punto cardinal (ej: "N", "NE")
	// Tip: podés usar un slice de strings indexado por p
	panic("no implementado")
}

func main() {
	// Si String() está bien implementada, fmt.Println usará tu método
	// automáticamente en lugar de imprimir el número entero.
	fmt.Println(N)
	fmt.Println(NE)
	fmt.Println(SE)
}
