package main

import "fmt"

func main() {

	// =========================================================
	// PASO 1: Declarar acumuladores y variables de control
	// Necesitamos contar cuántos pacientes caen en cada categoría
	// y rastrear la temperatura máxima y mínima del conjunto.
	// =========================================================

	var alta, normal, baja int
	var maxTemp, minTemp float64
	primera := true // bandera para inicializar max y min con el primer valor

	// =========================================================
	// PASO 2: Leer 10 temperaturas con un for clásico
	// En cada iteración leemos una temperatura, actualizamos
	// max/min y clasificamos al paciente.
	// =========================================================

	for i := 0; i < 10; i++ {
		var temp float64
		fmt.Printf("Temperatura paciente %d: ", i+1)
		fmt.Scan(&temp)

		// Inicializar max y min con el primer valor leído.
		// Si usáramos 0, el mínimo siempre quedaría en 0.
		if primera {
			maxTemp = temp
			minTemp = temp
			primera = false
		}

		if temp > maxTemp {
			maxTemp = temp
		}
		if temp < minTemp {
			minTemp = temp
		}

		// =====================================================
		// PASO 3: Clasificar la temperatura del paciente
		// >37.5 → fiebre (alta)
		// 36.0 a 37.5 inclusive → normal
		// <36.0 → hipotermia (baja)
		// =====================================================

		if temp > 37.5 {
			alta++
		} else if temp >= 36.0 {
			normal++
		} else {
			baja++
		}
	}

	// =========================================================
	// PASO 4: Calcular e imprimir porcentajes
	// Convertimos los contadores a float64 antes de dividir
	// para no hacer división entera (que truncaría el resultado).
	// =========================================================

	fmt.Println("\n--- Resultados ---")
	fmt.Printf("Fiebre (>37.5):  %.1f%%\n", float64(alta)/10*100)
	fmt.Printf("Normal (36-37.5): %.1f%%\n", float64(normal)/10*100)
	fmt.Printf("Baja (<36):       %.1f%%\n", float64(baja)/10*100)

	// =========================================================
	// PASO 5: Mostrar el promedio entre la máxima y la mínima
	// El enunciado pide (maxTemp + minTemp) / 2, no el promedio
	// de todas las temperaturas.
	// =========================================================

	fmt.Printf("Promedio (max+min)/2: %.2f°C  (max=%.1f, min=%.1f)\n",
		(maxTemp+minTemp)/2, maxTemp, minTemp)
}
