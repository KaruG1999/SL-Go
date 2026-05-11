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

// Para que al imprimir muestre nombre y no el numero asignado con iota
func (p PuntoCardinal) String() string {
    // Arreglo de strings en el mismo orden que el iota
    nombres := []string{"Norte", "Sur", "Este", "Oeste", "Noreste", "Sudoeste", "Noroeste", "Sudeste"}
    
    // Validación básica para evitar que el programa falle si p es inválido
    if int(p) < 0 || int(p) >= len(nombres) {
        return "Desconocido"
    }
    return nombres[p]
}

func contrarioOrden(p PuntoCardinal) PuntoCardinal {
	if p %2 == 0 {
		return p+1
	}
	return p-1
}

func main() {
	// detecta Stringer automaticamente y lo ejecuta
	fmt.Println(contrarioOrden(N))
	fmt.Println(contrarioOrden(NE))
}
