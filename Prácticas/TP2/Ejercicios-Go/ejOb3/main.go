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
	for _, elem := range s[1:] { // s[1:] = desde el segundo elemento hasta el final
		ultima := &result[len(result)-1] // puntero al último run para poder modificarlo
		if elem == ultima.valor {
			ultima.ocurrencias++
		} else {
			result = append(result, run{elem, 1})
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
	posicion := 0
	for _, r := range o {
		if r.valor == element {
			return posicion
		}
		posicion += r.ocurrencias
	}
	return -1
}

func Mode(o OptimumSlice) int {
	// make(map[K]V): crea un mapa vacío
	conteo := make(map[int]int)
	for _, r := range o {
		conteo[r.valor] += r.ocurrencias
	}
	valorModa, maxOcurrencias := 0, 0
	for valor, cant := range conteo {
		if cant > maxOcurrencias {
			maxOcurrencias = cant
			valorModa = valor
		}
	}
	return valorModa
}

// copia el slice sin mutar el original
func copiar(o OptimumSlice) OptimumSlice {
	copia := make(OptimumSlice, len(o))
	copy(copia, o)
	return copia
}

func Insert(o OptimumSlice, element int, position int) OptimumSlice {
	posicion := 0

	for i, r := range o {
		if posicion+r.ocurrencias > position {
			elementosAntes := position - posicion // cuántos elementos de esta racha van antes del nuevo

			// mismo valor que la racha actual → solo incrementar
			if element == r.valor {
				resultado := copiar(o)
				resultado[i].ocurrencias++
				return resultado
			}

			// inserción en el límite y coincide con la racha anterior → fusionar
			if elementosAntes == 0 && i > 0 && element == o[i-1].valor {
				resultado := copiar(o)
				resultado[i-1].ocurrencias++
				return resultado
			}

			// caso general: partir la racha en tres partes
			antes   := run{r.valor, elementosAntes}
			nuevo   := run{element, 1}
			despues := run{r.valor, r.ocurrencias - elementosAntes}

			// make(tipo, len, cap): slice vacío con capacidad pre-asignada
			resultado := make(OptimumSlice, 0, len(o)+2)
			resultado = append(resultado, o[:i]...)   // o[:i]...: expande el sub-slice como argumentos
			if antes.ocurrencias > 0 {
				resultado = append(resultado, antes)
			}
			resultado = append(resultado, nuevo)
			if despues.ocurrencias > 0 {
				resultado = append(resultado, despues)
			}
			resultado = append(resultado, o[i+1:]...) // o[i+1:]: desde el run siguiente al final
			return resultado
		}
		posicion += r.ocurrencias
	}

	// insertar al final
	resultado := copiar(o)
	if !IsEmpty(resultado) && resultado[len(resultado)-1].valor == element {
		resultado[len(resultado)-1].ocurrencias++
		return resultado
	}
	return append(resultado, run{element, 1})
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
	base := []int{3, 3, 3, 1, 1, 5, 5, 5}
	o := New(base)
	fmt.Printf("Base: %v\n\n", []run(o))

	// caso 1: insert al inicio (posición 0, valor diferente)
	r1 := Insert(o, 9, 0)
	fmt.Printf("Insert(9, 0) — al inicio:\n  %v\n", []run(r1))

	// caso 2: insert al final
	r2 := Insert(o, 9, 9999)
	fmt.Printf("Insert(9, final):\n  %v\n", []run(r2))

	// caso 3: insert en medio rompiendo un bloque
	// base: [{3,3},{1,2},{5,3}] → insertar 9 en posición 1 (dentro del run {3,3})
	r3 := Insert(o, 9, 1)
	fmt.Printf("Insert(9, 1) — rompe bloque {3,3}:\n  %v\n", []run(r3))

	// caso 4: fusión con bloque vecino
	// base: [{3,3},{1,2},{5,3}] → insertar 3 en posición 3 (límite entre {3,3} y {1,2})
	// sin fusión quedaría [{3,3},{3,1},{1,2},{5,3}] ← inválido
	// con fusión debe quedar [{3,4},{1,2},{5,3}]
	r4 := Insert(o, 3, 3)
	fmt.Printf("Insert(3, 3) — fusión con bloque anterior:\n  %v\n", []run(r4))

	// verificar que o no fue mutado
	fmt.Printf("\nOriginal sin mutar: %v\n", []run(o))
}
