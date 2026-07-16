package main

import "fmt"

func Sum (s1, s2 []int) []int{
	// Me quedo con la menor longitud 
	n := len(s1)
	if len(s2) < n {
		n = len(s2)
	}
	// Creo nuevo slice 
	res := make([]int, n)
	// sumo posiciones
	for i:=0; i<n; i++ {
		res[i] = s1[i] + s2[i]
	}
	return res
}


func Prom (s []int) float64 {
	if len(s) ==0 {
		return 0
	}
	suma :=0
	for _, v := range s {
		suma += v
	}
	return float64(suma) / float64(len(s))
}


func main () {
	// creo el slice  
	s1 := []int{1,2,3,4,5}
	s2 := []int{10,20,30}

	resultadoSuma := Sum(s1,s2)
	fmt.Println("Suma de Slices:", resultadoSuma)

	promedio := Prom(s1)
	fmt.Printf("Promedio de Slices: %.2f\n", promedio)
}
