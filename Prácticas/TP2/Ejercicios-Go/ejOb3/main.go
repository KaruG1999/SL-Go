package main

import "fmt"


type nodo struct {
	valor int
	cant  int
}

// slice de nodos comprimidos (RLE)
type OptimumSlice []nodo

// arma el OptimumSlice recorriendo el slice de enteros y agrupando los repetidos
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

// vacío si no quedó ningún nodo
func IsEmpty(o OptimumSlice) bool {
	return len(o) == 0
}

// es la suma de todas las cantidades, el tamaño del arreglo expandido
func Len(o OptimumSlice) int {
	total := 0
	for _, n := range o {
		total += n.cant
	}
	return total
}

// primer valor del arreglo expandido
func FrontElement(o OptimumSlice) int {
	if IsEmpty(o) { // corregido: los profes marcaron que faltaba este chequeo (panic index out of range)
		panic("OptimumSlice vacío")
	}
	return o[0].valor
}

// ultimo valor del arreglo expandido
func LastElement(o OptimumSlice) int {
	if IsEmpty(o) { // corregido: mismo problema que FrontElement, faltaba el chequeo de vacío
		panic("OptimumSlice vacío")
	}
	return o[len(o)-1].valor
}

// promedio del arreglo expandido, pesando cada valor por su cantidad
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

// cuenta cuantas veces aparece un valor en total (sumando todos los nodos que lo tengan)
func Occurrences(o OptimumSlice, element int) int {
	total := 0
	for _, r := range o {
		if r.valor == element {
			total += r.cant
		}
	}
	return total
}

// posición de la primera aparición, pero en el arreglo expandido, no el índice del nodo
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

// el valor que mas se repite. corregido: antes
// llamaba a Occurrences(o, n.valor) para cada nodo, y Occurrences recorre
// todo el slice -> O(n^2). ahora se acumula en el map en una sola pasada
func Mode(o OptimumSlice) int {
	conteo := make(map[int]int)
	max := 0
	mode := 0
	for _, n := range o { // única pasada: se cuenta y se compara el máximo en el mismo loop
		conteo[n.valor] += n.cant
		if conteo[n.valor] > max {
			max = conteo[n.valor]
			mode = n.valor
		}
	}
	return mode
}

// copia aparte para no romper el original al armar el resultado de Insert
func copiar(o OptimumSlice) OptimumSlice {
	copia := make(OptimumSlice, len(o))
	copy(copia, o)
	return copia
}

// agrega un nodo al final, pero si tiene el mismo valor que el ultimo nodo ya puesto  
// los junta en vez de dejar dos nodos separados con el mismo valor
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

// se arma todo de nuevo con appendMerge para que si el valor insertado coincide
// con el nodo de al lado, se fusionen en vez de quedar nodos sueltos repetidos.
func Insert(o OptimumSlice, element int, position int) OptimumSlice {
	if position < 0 || position > Len(o) {
		panic("posición fuera de rango")
	}

	posicion := 0

	for i, r := range o {
		if posicion+r.cant > position {
			// la posición cae adentro de este nodo (r), elementosAntes es
			// cuanto de r queda a la izquierda del punto de corte
			elementosAntes := position - posicion

			// si justo el valor a insertar es el mismo que el del nodo, no hay
			// que partir nada, solo crece la cantidad
			if element == r.valor {
				resultado := copiar(o)
				resultado[i].cant++
				return resultado
			}

			// partimos r en "antes" + el nuevo elemento + "despues"
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

	// no encontró en que nodo cortar porque la posición es el final absoluto
	resultado := copiar(o)
	resultado = appendMerge(resultado, nodo{element, 1})
	return resultado
}

// vuelve a expandir todo a un []int comun, cada nodo repetido "cant" veces
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
	// con []nodo le indicamos que imprima en formao crudo del slice [{a b}..]
	fmt.Println(o)
	

	// al inicio
	r1 := Insert(o, 9, 0)
	fmt.Println("Insert(9,0) -> Esperado: [{9 1} {3 3} {1 2} {5 3}]")
	fmt.Println(r1)
	

	// en medio, rompiendo un bloque
	r2 := Insert(o, 9, 1)
	fmt.Println("Insert(9,1) -> Esperado: [{3 1} {9 1} {3 2} {1 2} {5 3}]")
	fmt.Println(r2)
	
	// fusiona con bloque existente (mismo valor)
	r3 := Insert(o, 3, 3)
	fmt.Println("Insert(3,3) -> Esperado: [{3 4} {1 2} {5 3}]")
	fmt.Println(r3)
	

	// fusiona con el bloque de la derecha
	o2 := New([]int{3, 3, 3, 1, 1, 3, 3})
	fmt.Println("Segundo arreglo:")
	fmt.Println(o2)
	

	r4 := Insert(o2, 3, 5)
	fmt.Println("Insert(3,5) -> Esperado: [{3 3} {1 2} {3 3}]")
	fmt.Println(r4)
}
