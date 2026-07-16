package main

import (
	"fmt"
	"sort"
)

type Carrera string

const (
	APU Carrera = "APU"
	LI  Carrera = "LI"
	LS  Carrera = "LS"
)

type Fecha struct {
	Dia, Mes, Anio int
}

type Ingresante struct {
	Apellido    string
	Nombre      string
	Ciudad      string
	Nacimiento  Fecha
	TieneTitulo bool
	Carrera     Carrera
}

func (i Ingresante) String() string {
	titulo := "No"
	if i.TieneTitulo {
		titulo = "Si"
	}
	return fmt.Sprintf("%s, %s (%s) - Nac: %02d/%02d/%d - Titulo secundario: %s - Carrera: %s",
		i.Apellido, i.Nombre, i.Ciudad,
		i.Nacimiento.Dia, i.Nacimiento.Mes, i.Nacimiento.Anio,
		titulo, i.Carrera)
}

// MasJoven: true si a nacio despues que b
func MasJoven(a, b Ingresante) bool {
	if a.Nacimiento.Anio != b.Nacimiento.Anio {
		return a.Nacimiento.Anio > b.Nacimiento.Anio
	}
	if a.Nacimiento.Mes != b.Nacimiento.Mes {
		return a.Nacimiento.Mes > b.Nacimiento.Mes
	}
	return a.Nacimiento.Dia > b.Nacimiento.Dia
}

// MenorAlfabetico: compara apellido y despues nombre
func MenorAlfabetico(a, b Ingresante) bool {
	if a.Apellido != b.Apellido {
		return a.Apellido < b.Apellido
	}
	return a.Nombre < b.Nombre
}

func main() {
	ingresantes := []Ingresante{
		{"Perez", "Ana", "La Plata", Fecha{12, 3, 2003}, true, LI},
		{"Gomez", "Luis", "Bariloche", Fecha{5, 11, 2001}, false, APU},
		{"Fernandez", "Sol", "La Plata", Fecha{20, 7, 2002}, true, LS},
		{"Perez", "Juan", "Cordoba", Fecha{1, 1, 2003}, true, APU},
	}

	fmt.Println("=== Sin ordenar ===")
	for _, i := range ingresantes {
		fmt.Println(i)
	}

	fmt.Println("\n=== Ordenados por edad (mas joven primero) ===")
	sort.Slice(ingresantes, func(i, j int) bool {
		return MasJoven(ingresantes[i], ingresantes[j])
	})
	for _, i := range ingresantes {
		fmt.Println(i)
	}

	fmt.Println("\n=== Ordenados por apellido y nombre ===")
	sort.Slice(ingresantes, func(i, j int) bool {
		return MenorAlfabetico(ingresantes[i], ingresantes[j])
	})
	for _, i := range ingresantes {
		fmt.Println(i)
	}
}
