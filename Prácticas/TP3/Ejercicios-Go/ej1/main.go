package main

import "fmt"

// Tipo generico mapa ( K debe ser comparable; V cualquier tipo de dato) retorna un dato map
type Map[K comparable, V any] map[K]V

func main() {
	// --- USO 1: Clave string, Valor int ---
	inventario := make(Map[string, int])
	
	inventario["Arroz"] = 15
	inventario["Fideos"] = 4
	
	fmt.Println("--- Uso 1 (Map[string, int]) ---")
	for k, v := range inventario {
		fmt.Printf("Clave: %s | Valor: %d\n", k, v)
	}

	fmt.Println()

	// --- USO 2: Clave int, Valor string ---
	usuarios := make(Map[int, string])
	
	usuarios[1024] = "Karen"
	usuarios[2048] = "Alejandro"
	
	fmt.Println("--- Uso 2 (Map[int, string]) ---")
	for k, v := range usuarios {
		fmt.Printf("ID: %d | Usuario: %s\n", k, v)
	}
}
