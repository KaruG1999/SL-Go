package main

import "fmt"

// Vetor de flotantes de tamaño fijo
const N = 3
type Vector [N]float64

func Initialize (v *Vector, f float64){
	for i:=0; i<N; i++ {
		v[i] = f
	}
}

// Retorna un vector nuevos, es decir, los originales no cambian
func Sum (v1, v2 Vector) Vector{
	var v Vector
	for i:=0; i<N; i++ {
		v[i] = v1[i] + v2[i]
	}
	return v
}

// Resultado se almacena en v1 (* -> pasaje por referencia)
func SumInPlace (v1 *Vector, v2 Vector) {
	for i:=0; i<N; i++{
		v1[i] = v1[i] + v2[i]  // Se modifica directamente en memoria
	}
}

func main (){
	var vecA Vector
	var vecB Vector

	// 1. Probar Initialize (usamos & para pasar la dirección de memoria)
	Initialize(&vecA, 10.5)
	Initialize(&vecB, 5.0)
	fmt.Println("Vector A inicializado:", vecA)
	fmt.Println("Vector B inicializado:", vecB)

	// 2. Probar Sum (retorna un vector nuevo, vecA y vecB no cambian)
	vecC := Sum(vecA, vecB)
	fmt.Println("Suma (nuevo vector C):", vecC)
	fmt.Println("Vector A sigue igual:", vecA)

	// 3. Probar SumInPlace (modifica el original vecA) 
	SumInPlace(&vecA, vecB)
	fmt.Println("Vector A después de SumInPlace (A = A + B):", vecA)
}
