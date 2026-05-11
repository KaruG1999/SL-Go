package main

import "fmt"

const N = 10

type Celsius float64

func main() {
	var temps [N]Celsius
	for i := range temps {
		fmt.Scan(&temps[i])
	}

	maxT := temps[0]
	minT := temps[0]
	grupos := map[string]int{
		"alta": 0, 
		"normal": 0, 
		"baja": 0}

	for _, temp := range temps {
		if temp > maxT {
			maxT = temp
		}
		if temp < minT {
			minT = temp
		}
		switch {
		case temp > 37.5:
			grupos["alta"]++
		case temp >= 36:
			grupos["normal"]++
		default:
			grupos["baja"]++
		}
	}

	promedio := int((maxT + minT) / 2)

	fmt.Printf("Alta:   %.1f%%\n", float64(grupos["alta"])/N*100)
	fmt.Printf("Normal: %.1f%%\n", float64(grupos["normal"])/N*100)
	fmt.Printf("Baja:   %.1f%%\n", float64(grupos["baja"])/N*100)
	fmt.Printf("Promedio (max+min)/2: %d°C\n", promedio)


	fmt.Println("Conteos por Map: ", grupos)
}
