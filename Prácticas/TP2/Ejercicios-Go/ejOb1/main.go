package main

import "fmt"

type Fecha struct { // Opcional? 
	Dia  int
	Mes  int
	Anio int
}

type Ingresante struct {
	Apellido    string
	Nombre      string
	Ciudad      string
	Nacimiento  Fecha
	TieneTitulo bool
	Carrera     string // EJ: "APU", "LI" o "LS"
}

//  Lista enlazada de Ingresante

type nodo struct {
	elem Ingresante
	sig  *nodo
}

type List struct {
	pri *nodo
	len int
}

func New() List { return List{} }

func (l List) IsEmpty() bool { return l.len == 0 }

func (l *List) PushBack(ing Ingresante) {
	nuevo := &nodo{elem: ing}
	if l.pri == nil {
		l.pri = nuevo
	} else {
		actual := l.pri
		for actual.sig != nil {
			actual = actual.sig
		}
		actual.sig = nuevo
	}
	l.len++
}

// ── Recorrido único ─────

func procesarLista(l *List) {
	anioConteo := make(map[int]int)
	carreraConteo := make(map[string]int)

	fmt.Println("a) Ingresantes de Bariloche:")

	var prev *nodo = nil
	curr := l.pri

	for curr != nil {
		ing := curr.elem

		// a) ciudad Bariloche
		if ing.Ciudad == "Bariloche" {
			fmt.Printf("   %s, %s\n", ing.Apellido, ing.Nombre)
		}

		// b) acumular por año de nacimiento
		anioConteo[ing.Nacimiento.Anio]++

		// c) acumular por carrera
		carreraConteo[ing.Carrera]++

		// d) eliminar si no tiene título
		if !ing.TieneTitulo {
			if prev == nil {
				l.pri = curr.sig
			} else {
				prev.sig = curr.sig
			}
			l.len--
			curr = curr.sig
			continue
		}

		prev = curr
		curr = curr.sig
	}

	// b) año con más ingresantes
	maxAnio, maxCantAnio := 0, 0
	for anio, cant := range anioConteo {
		if cant > maxCantAnio {
			maxCantAnio = cant
			maxAnio = anio
		}
	}
	fmt.Printf("\nb) Año con más ingresantes: %d (%d ingresantes)\n", maxAnio, maxCantAnio)

	// c) carrera con más inscriptos
	maxCarrera, maxCant := "", 0
	for carrera, cant := range carreraConteo {
		if cant > maxCant {
			maxCant = cant
			maxCarrera = carrera
		}
	}
	fmt.Printf("\nc) Carrera con más inscriptos: %s (%d inscriptos)\n", maxCarrera, maxCant)
}

func imprimirLista(l List) {
	fmt.Print("Lista resultante: ")
	for n := l.pri; n != nil; n = n.sig {
		fmt.Printf("[%s %s] -> ", n.elem.Nombre, n.elem.Apellido)
	}
	fmt.Println()
}

// ── Main ────

func main() {
	l := New()

	l.PushBack(Ingresante{"García", "Ana", "Bariloche", Fecha{1, 3, 2005}, true, "LI"})
	l.PushBack(Ingresante{"López", "Juan", "La Plata", Fecha{5, 7, 2004}, false, "LS"})
	l.PushBack(Ingresante{"Pérez", "Lucía", "Bariloche", Fecha{12, 11, 2005}, true, "APU"})
	l.PushBack(Ingresante{"Martínez", "Carlos", "Buenos Aires", Fecha{20, 2, 2004}, true, "LI"})
	l.PushBack(Ingresante{"Ruiz", "Sofía", "Bariloche", Fecha{8, 6, 2005}, false, "LS"})
	l.PushBack(Ingresante{"Díaz", "Pedro", "Rosario", Fecha{3, 9, 2004}, true, "APU"})
	l.PushBack(Ingresante{"Fernández", "Marta", "La Plata", Fecha{17, 1, 2005}, true, "LI"})

	procesarLista(&l)

	fmt.Printf("\nd) Ingresantes sin título eliminados.\n")
	imprimirLista(l)
}
