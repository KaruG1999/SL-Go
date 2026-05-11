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

func contrarioSwitch(p PuntoCardinal) PuntoCardinal {
	 switch p{
		case N:
			return S
		case S:
			return N
		case E:
			return O
		case O:
			return E
		case NE:
        	return SO
		case SO:
			return NE
		case NO:
			return SE
		case SE:
			return NO
		default:
			// Es obligatorio manejar un caso por defecto o tener un return al final
			return -1 
		}
	 }
	

func main() {
	fmt.Println(contrarioSwitch(N))
	fmt.Println(contrarioSwitch(NE))
}
