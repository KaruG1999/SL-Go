package main

import "fmt"

func main() {
	/* integers */
	var zz int = 10  // 0A no es un número válido en Go
	x := 10     // formas correctas -> z := x o var z int = x o ( y sacamos ; )
	var z int = x    // Se debe asignar un int primero
	var y int8 = int8(x + 1)  // Conversión explícita (Casting) con int8
	const n = 5001 // := genera variables, nunca constantes
	const c int = 5001

	/* float */
	var e float32 = 6   // otra form -> e := float32(6)
	var f float32 = e

	fmt.Println(zz, x, z, y, n, c, e, f)
}
