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

var opuesto = map[PuntoCardinal]PuntoCardinal{
    N:  S,
    S:  N,
    E:  O,
    O:  E,
    NE: SO,
    SO: NE,
    NO: SE,
    SE: NO,
}

func (p PuntoCardinal) String() string {
    // Arreglo de strings en el mismo orden que el iota
    nombres := []string{"Norte", "Sur", "Este", "Oeste", "Noreste", "Sudoeste", "Noroeste", "Sudeste"}
    
    // Validación básica para evitar que el programa falle si p es inválido
    if int(p) < 0 || int(p) >= len(nombres) {
        return "Desconocido"
    }
    return nombres[p]
}

func contrarioMap(p PuntoCardinal) PuntoCardinal {
	return opuesto[p]
}

func main() {
	fmt.Println(contrarioMap(N))
	fmt.Println(contrarioMap(NE))
}
