package main

import "fmt"

// [3,3,3,1,1] se guarda como [{3,3},{1,2}]
type run struct {
	valor       int
	ocurrencias int
}

// tipo basado en []run
type OptimumSlice []run

func New(s []int) OptimumSlice {
	if len(s) == 0 {
		return OptimumSlice{}
	}
	// {{s[0], 1}}: slice literal con un run inicial
	result := OptimumSlice{{s[0], 1}}
	for _, v := range s[1:] { // s[1:] = desde el segundo elemento hasta el final
		last := &result[len(result)-1] // puntero al último run para poder modificarlo
		if v == last.valor {
			last.ocurrencias++
		} else {
			result = append(result, run{v, 1})
		}
	}
	return result
}

func IsEmpty(o OptimumSlice) bool {
	return len(o) == 0
}

func Len(o OptimumSlice) int {
	total := 0
	for _, r := range o {
		total += r.ocurrencias
	}
	return total
}

func FrontElement(o OptimumSlice) int {
	return o[0].valor
}

func LastElement(o OptimumSlice) int {
	return o[len(o)-1].valor
}

func Average(o OptimumSlice) float64 {
	var suma float64
	for _, r := range o {
		suma += float64(r.valor) * float64(r.ocurrencias)
	}
	return suma / float64(Len(o))
}

func Occurrences(o OptimumSlice, element int) int {
	total := 0
	for _, r := range o {
		if r.valor == element {
			total += r.ocurrencias
		}
	}
	return total
}

func IndexOf(o OptimumSlice, element int) int {
	pos := 0
	for _, r := range o {
		if r.valor == element {
			return pos
		}
		pos += r.ocurrencias
	}
	return -1
}

func Mode(o OptimumSlice) int {
	// make(map[K]V): crea un mapa vacío
	conteo := make(map[int]int)
	for _, r := range o {
		conteo[r.valor] += r.ocurrencias
	}
	bestVal, bestCant := 0, 0
	for val, cant := range conteo {
		if cant > bestCant {
			bestCant = cant
			bestVal = val
		}
	}
	return bestVal
}

func Insert(o OptimumSlice, element int, position int) OptimumSlice {
	posAcum := 0

	for i, r := range o {
		if posAcum+r.ocurrencias > position {
			offset := position - posAcum

			if element == r.valor {
				o[i].ocurrencias++
				return o
			}

			antes   := run{r.valor, offset}
			nuevo   := run{element, 1}
			despues := run{r.valor, r.ocurrencias - offset}

			// make(tipo, len, cap): slice vacío con capacidad pre-asignada
			result := make(OptimumSlice, 0, len(o)+2)
			result = append(result, o[:i]...)   // o[:i]...: expande el sub-slice como argumentos
			if antes.ocurrencias > 0 {
				result = append(result, antes)
			}
			result = append(result, nuevo)
			if despues.ocurrencias > 0 {
				result = append(result, despues)
			}
			result = append(result, o[i+1:]...) // o[i+1:]: desde el run siguiente al final
			return result
		}
		posAcum += r.ocurrencias
	}

	if !IsEmpty(o) && o[len(o)-1].valor == element {
		o[len(o)-1].ocurrencias++
		return o
	}
	return append(o, run{element, 1})
}

func SliceArray(o OptimumSlice) []int {
	// make([]int, 0, cap): evita re-alocar memoria durante los append
	result := make([]int, 0, Len(o))
	for _, r := range o {
		for i := 0; i < r.ocurrencias; i++ {
			result = append(result, r.valor)
		}
	}
	return result
}

func main() {
	s := []int{3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 1, 1, 23, 23, 23, 23, 23, 23, 3, 3, 3, 3, 3, 3, 3, 3, 7, 5, 5, 5}
	o := New(s)

	fmt.Printf("Runs:           %v\n", []run(o)) // []run(o): convierte OptimumSlice a []run para imprimir
	fmt.Printf("Len:            %d\n", Len(o))
	fmt.Printf("FrontElement:   %d\n", FrontElement(o))
	fmt.Printf("LastElement:    %d\n", LastElement(o))
	fmt.Printf("Average:        %.2f\n", Average(o))
	fmt.Printf("Occurrences(3): %d\n", Occurrences(o, 3))
	fmt.Printf("IndexOf(23):    %d\n", IndexOf(o, 23))
	fmt.Printf("Mode:           %d\n", Mode(o))

	fmt.Println("\n--- Insert(o, 9, 6) ---")
	o2 := Insert(o, 9, 6)
	fmt.Printf("Runs:  %v\n", []run(o2))
	fmt.Printf("Array: %v\n", SliceArray(o2))

	fmt.Println("\n--- Insert mismo valor ---")
	o3 := Insert(o, 3, 2)
	fmt.Printf("Runs: %v\n", []run(o3))

	fmt.Println("\n--- Insert al final ---")
	o4 := Insert(o, 99, 9999)
	fmt.Printf("Último run: %v\n", o4[len(o4)-1])

	fmt.Printf("\n¿o vacío?:       %v\n", IsEmpty(o))
	fmt.Printf("¿New([]) vacío?: %v\n", IsEmpty(New([]int{})))
}
