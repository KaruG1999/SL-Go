package main

import "fmt"

type nodo struct {
	valor int
	cant  int
}

type OptimumSlice []nodo

func New(s []int) OptimumSlice {
	if len(s) == 0 {
		return OptimumSlice{}
	}

	var result OptimumSlice
	elem := s[0]
	cant := 1

	for i := 1; i < len(s); i++ {
		if s[i] == elem {
			cant++
		} else {
			result = append(result, nodo{elem, cant})
			elem = s[i]
			cant = 1
		}
	}

	result = append(result, nodo{elem, cant})
	return result
}

func IsEmpty(o OptimumSlice) bool {
	return len(o) == 0
}

func Len(o OptimumSlice) int {
	total := 0
	for _, n := range o {
		total += n.cant
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
	total := 0
	cant := 0
	for _, r := range o {
		total += r.valor * r.cant
		cant += r.cant
	}
	if cant == 0 {
		return 0
	}
	return float64(total) / float64(cant)
}

func Occurrences(o OptimumSlice, element int) int {
	total := 0
	for _, r := range o {
		if r.valor == element {
			total += r.cant
		}
	}
	return total
}

func IndexOf(o OptimumSlice, element int) int {
	pos := 0
	for _, n := range o {
		if n.valor == element {
			return pos
		}
		pos += n.cant
	}
	return -1
}

func Mode(o OptimumSlice) int {
	max := 0
	mode := 0
	for _, n := range o {
		cant := Occurrences(o, n.valor)
		if cant > max {
			max = cant
			mode = n.valor
		}
	}
	return mode
}

func copiar(o OptimumSlice) OptimumSlice {
	copia := make(OptimumSlice, len(o))
	copy(copia, o)
	return copia
}

func appendMerge(o OptimumSlice, r nodo) OptimumSlice {
	if r.cant <= 0 {
		return o
	}
	if len(o) > 0 && o[len(o)-1].valor == r.valor {
		o[len(o)-1].cant += r.cant
		return o
	}
	return append(o, r)
}

func Insert(o OptimumSlice, element int, position int) OptimumSlice {
	if position < 0 || position > Len(o) {
		panic("posición fuera de rango")
	}

	posicion := 0

	for i, r := range o {
		if posicion+r.cant > position {
			elementosAntes := position - posicion

			if element == r.valor {
				resultado := copiar(o)
				resultado[i].cant++
				return resultado
			}

			antes := nodo{r.valor, elementosAntes}
			nuevo := nodo{element, 1}
			despues := nodo{r.valor, r.cant - elementosAntes}

			resultado := make(OptimumSlice, 0, len(o)+2)

			for _, x := range o[:i] {
				resultado = appendMerge(resultado, x)
			}
			resultado = appendMerge(resultado, antes)
			resultado = appendMerge(resultado, nuevo)
			resultado = appendMerge(resultado, despues)
			for _, x := range o[i+1:] {
				resultado = appendMerge(resultado, x)
			}
			return resultado
		}
		posicion += r.cant
	}

	// Insertar al final absoluto
	resultado := copiar(o)
	resultado = appendMerge(resultado, nodo{element, 1})
	return resultado
}

func SliceArray(o OptimumSlice) []int {
	var result []int
	for _, n := range o {
		for i := 0; i < n.cant; i++ {
			result = append(result, n.valor)
		}
	}
	return result
}

func main() {
	base := []int{3, 3, 3, 1, 1, 5, 5, 5}
	o := New(base)

	fmt.Println("Slice original:")
	fmt.Println([]nodo(o))
	fmt.Println()

	// al inicio
	r1 := Insert(o, 9, 0)
	fmt.Println("Insert(9,0) -> Esperado: [{9 1} {3 3} {1 2} {5 3}]")
	fmt.Println([]nodo(r1))
	fmt.Println()

	// en medio, rompiendo un bloque
	r2 := Insert(o, 9, 1)
	fmt.Println("Insert(9,1) -> Esperado: [{3 1} {9 1} {3 2} {1 2} {5 3}]")
	fmt.Println([]nodo(r2))
	fmt.Println()

	// fusiona con bloque existente (mismo valor)
	r3 := Insert(o, 3, 3)
	fmt.Println("Insert(3,3) -> Esperado: [{3 4} {1 2} {5 3}]")
	fmt.Println([]nodo(r3))
	fmt.Println()

	// fusiona con el bloque de la derecha
	o2 := New([]int{3, 3, 3, 1, 1, 3, 3})
	fmt.Println("Segundo arreglo:")
	fmt.Println([]nodo(o2))
	fmt.Println()

	r4 := Insert(o2, 3, 5)
	fmt.Println("Insert(3,5) -> Esperado: [{3 3} {1 2} {3 3}]")
	fmt.Println([]nodo(r4))
}
