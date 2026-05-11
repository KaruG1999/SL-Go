# Ejercicio 13 — Banco

## Enunciado

Un banco tiene un listado de clientes que vienen a pagar impuestos. De cada cliente se conoce: DNI, Nombre, Apellido, código de impuesto (A, B, C o D) y monto a pagar.

**a)** Atender clientes hasta recaudar al menos $10.000 o hasta que no queden más clientes.

**b)** Al finalizar, informar el código de impuesto que más veces se pagó.

**c)** Al finalizar, informar cuántos clientes quedaron sin atender (si los hay).

---

## Lógica de resolución

### Tipos

```go
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
```

### Parte a — atención de clientes

```go
clientes := []Cliente{ /* cargar datos */ }

const META = 10000.0
var recaudado float64
var atendidos int

for i, c := range clientes {
    if recaudado >= META {
        break
    }
    recaudado += c.Monto
    atendidos = i + 1
}

fmt.Printf("Recaudado: $%.2f con %d clientes atendidos\n", recaudado, atendidos)
```

### Parte b — impuesto más pagado

Usar un `map` para contar cuántas veces se pagó cada código:

```go
conteo := make(map[CodigoImpuesto]int)

for _, c := range clientes[:atendidos] {
    conteo[c.Impuesto]++
}

// Encontrar el máximo
var masVeces CodigoImpuesto
var maxConteo int
for cod, cant := range conteo {
    if cant > maxConteo {
        maxConteo = cant
        masVeces = cod
    }
}
fmt.Printf("Impuesto más pagado: %s (%d veces)\n", masVeces, maxConteo)
```

### Parte c — clientes sin atender

```go
sinAtender := len(clientes) - atendidos
if sinAtender > 0 {
    fmt.Printf("Clientes sin atender: %d\n", sinAtender)
} else {
    fmt.Println("Todos los clientes fueron atendidos")
}
```

> El map es la estructura natural para acumular conteos por categoría en Go. Al recorrer el slice de clientes ya atendidos (`clientes[:atendidos]`), se usa la sintaxis de slice parcial que aprendimos en clase.
