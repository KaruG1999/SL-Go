package main

import (
	"fmt"
)


type Celsius float64
type Fahrenheit float64

func main() {
	var temps [10]Celsius
	var alta, media, baja int 
	// Leo los 10 datos
	for i:=0; i<10; i++ {
		fmt.Scan(&temps[i]) // Lee y guarda en pos i del array temps
	}

	maxT := temps[0] // Inicializo maxT con el primer valor del array
	minT := temps[0] // Inicializo minT con el primer valor del array

	for _, temp := range temps {
		// Evaluo dato 
		if temp > maxT { maxT = temp}
		if temp < minT { minT = temp}
		// clasifico por rangos 
		switch {
		case temp > 37.5:
			alta++
		case temp >= 37.5:
			media++
		default:
			baja++
		}
	}

	promedio := int ((maxT + minT) /2)

	fmt.Println("Porcentajes:")
	fmt.Println("Alta: %.1f%%", float64(alta)/10*100)
	fmt.Println("Normal: %.1f%%", float64(media)/10*100)
	fmt.Println("Baja: %.1f%%", float64(baja)/10*100)
	fmt.Println("Promedio: %d", promedio)



}