package main

import "fmt"

type Celsius float64
type Fahrenheit float64


// Funcion de conversion
func CelsiusAFahrenheit (c Celsius) Fahrenheit{
	return Fahrenheit((c*9/5)+32)
}


func main() {
	var tempsCelsius [5]Celsius = [5]Celsius{36.5, 38.2, 35.0, 37.0, 39.5}
	var tempsFahr [5]Fahrenheit 

	for i, t:= range tempsCelsius {
		tempsFahr[i] = CelsiusAFahrenheit(t)
		fmt.Printf(" %.1f°C es %.1f°F\n", tempsCelsius[i], tempsFahr[i])
	}
	
}

