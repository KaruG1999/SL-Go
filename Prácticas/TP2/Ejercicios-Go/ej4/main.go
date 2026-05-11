package main

import ( 
	"fmt"
)

const N = 5

func maxmin (serie [N]int) (int, int){
	max := serie[0]
	min := serie[0]

	for _, v := range serie {
		if v > max {max=v}
		if v < min {min=v}
	}
	return max, min
}

func productoria (serie [N]int) int {
	p := 1
	for _, v := range serie {
		 p *= (v*v*v)
	}
	return p
}

func sumatoria (serie [N]int) float64 {
	var s float64
	for _, v := range serie {
		s += 1.0 / float64(v) // Obligatorio para no perder decimal
	}
	return s
}

func main (){
	var x [N]int
	var y [N]int
	var z [N]int

	fmt.Println("Ingrese los calores para X, Y, Z: ")
	for i:= 0; i<N; i++ {fmt.Scan(&x[i])}
	for i:= 0; i<N; i++ {fmt.Scan(&y[i])}
	for i:= 0; i<N; i++ {fmt.Scan(&z[i])}

	// Calculos con distintas funciones 
	s := sumatoria(x) // Retorna float64
	p := productoria(z)
	max, min := maxmin(y)

	resto := s - float64(p)
	resultado := resto * float64(max) * float64 (min)

	fmt.Printf("El resultado es: %.4f\n", resultado)
}

