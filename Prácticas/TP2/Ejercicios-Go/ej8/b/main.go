package main

import "fmt"

const digitos = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Convert(v int, b int) string {
	// caso si dato es 0
	if v==0 {return "0"}
	// caso si b es inválido
	if b<2 || b>32 {return "Base inválida"}
	// en string vamos acumulando el resultado
	resultado := ""
	negativo := false
	if v<0 { 
		negativo = true
		v = -v}
	for v>0 {
		
		resto := v % b
		digito := digitos[resto]
		// Luego sumo el digito al resultado ADELANTE
		resultado = string(digito) + resultado
		// Sigo con el resto del valor
		v = v/b 
	}
	if negativo == true {resultado = "-" + resultado}
	return resultado
}

func main() {
	fmt.Println(Convert(23, 2))   // "10111"
	fmt.Println(Convert(-10, 2))  // "-1010"
	fmt.Println(Convert(0, 10))   // "0"
}
