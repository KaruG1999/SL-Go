package main

import "fmt"

type CodigoImpuesto string

const (
	ImpA CodigoImpuesto = "A"
	ImpB CodigoImpuesto = "B"
	ImpC CodigoImpuesto = "C"
	ImpD CodigoImpuesto = "D"
)

type Cliente struct {
	DNI      int
	Nombre   string
	Apellido string
	Impuesto CodigoImpuesto
	Monto    float64
}

const META = 10000.0

func main() {
	clientes := []Cliente{
		{111, "Ana", "Perez", ImpA, 3000},
		{222, "Luis", "Gomez", ImpB, 2500},
		{333, "Sol", "Fernandez", ImpA, 4000},
		{444, "Juan", "Diaz", ImpC, 1000},
		{555, "Marta", "Ruiz", ImpD, 5000},
	}

	var recaudado float64
	atendidos := 0
	conteo := make(map[CodigoImpuesto]int)

	for _, c := range clientes {
		if recaudado >= META {
			break
		}
		recaudado += c.Monto
		conteo[c.Impuesto]++
		atendidos++
	}

	fmt.Printf("Recaudado: $%.2f con %d clientes atendidos\n", recaudado, atendidos)

	var masPagado CodigoImpuesto
	max := 0
	for cod, cant := range conteo {
		if cant > max {
			max = cant
			masPagado = cod
		}
	}
	fmt.Printf("Impuesto mas pagado: %s (%d veces)\n", masPagado, max)

	sinAtender := len(clientes) - atendidos
	if sinAtender > 0 {
		fmt.Printf("Clientes sin atender: %d\n", sinAtender)
	} else {
		fmt.Println("Se atendieron todos los clientes")
	}
}
