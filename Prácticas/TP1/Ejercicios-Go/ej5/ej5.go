package main

import (
	"fmt"
	"math"
)

/* f(x) = |x| si xe (-inf, -18), xmod4 si xe [-18,-1], x² si xe [1,20) -x si xe [20, inf) */
func evaluarFuncion(x float64) float64 {
	switch {
	case x < -18:
		return math.Abs(x)
	case x <= -1: // El límite inferior ya fue filtrado por el caso anterior
		return math.Mod(x, 4)
	case x < 1:
		// Definir comportamiento para (-1, 1).
		// Actualmente cae en el 'else' final (-x).
		return -x
	case x <= 20:
		return x * x // Más rápido que math.Pow(x, 2)
	default:
		return -x
	}
}

func main() {
	/*lea un número y muestre el valor
	correspondiente aplicando la siguiente función sobre el mismo */
	var num float64
	fmt.Print("Ingrese un número: ")
	fmt.Scanln(&num)

	result := evaluarFuncion(num)
	fmt.Printf("El resultado de evaluar la función para el número %.2f es: %.4f\n", num, result)

}

/* a. ¿Qué tiene de particular la función con el 0 (cero), se
puede escribir sin opción default/else?. Re-escribir con otra
estructura de control selectiva.
b. Re-escribir la función usando punto flotante.
Sub-objetivo: Uso de E/S de enteros y punto flotante.

Respuesta:
a. La función no tiene un caso específico para el valor 0, lo que significa que caerá en el caso 'else' final (-x), resultando en 0. Esto es correcto según la definición de la función, pero podría ser más claro si se manejara explícitamente el caso de 0. Se puede re-escribir usando una estructura if-else para mayor claridad:

func evaluarFuncion(x float64) float64 {
	if x < -18 {
		return math.Abs(x)
	} else if x <= -1 {
		return math.Mod(x, 4)
	} else if x < 1 {
		return -x
	} else if x <= 20 {
		return x * x
	} else {
		return -x
	}
}

b. La función ya está utilizando punto flotante (float64) para la entrada y salida, por lo que no es necesario hacer cambios adicionales para usar punto flotante. Sin embargo, si se quisiera enfatizar el uso de punto flotante, se podría asegurar que todas las constantes numéricas también sean de tipo float64:

func evaluarFuncion(x float64) float64 {
	if x < -18.0 {
		return math.Abs(x)
	} else if x <= -1.0 {
		return math.Mod(x, 4.0)
	} else if x < 1.0 {
		return -x
	} else if x <= 20.0 {
		return x * x
	} else {
		return -x
	}
}
*/
