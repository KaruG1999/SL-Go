# Ejercicio 13 — Banco

## Enunciado

Un banco tiene un listado de clientes que vienen a pagar impuestos. De cada cliente se conoce: DNI, Nombre, Apellido, código de impuesto (A, B, C o D) y monto a pagar.

**a)** Atender clientes hasta recaudar al menos $10.000 o hasta que no queden más clientes.

**b)** Al finalizar, informar el código de impuesto que más veces se pagó.

**c)** Al finalizar, informar cuántos clientes quedaron sin atender (si los hay).

---

## Lógica de resolución (como está en `main.go`)

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

const META = 10000.0
```

Las tres partes (a, b y c) se resuelven en un solo recorrido del slice de clientes, en vez de tres loops separados:

```go
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
```

Después de ese loop:

```go
// parte b: el código con más ocurrencias en 'conteo'
var masPagado CodigoImpuesto
max := 0
for cod, cant := range conteo {
    if cant > max {
        max = cant
        masPagado = cod
    }
}

// parte c: lo que quedó sin atender
sinAtender := len(clientes) - atendidos
```

## Observaciones

- Se corta la atención con `break` apenas se llega a la meta, así que `atendidos` y `conteo` solo reflejan a los clientes efectivamente atendidos (justo lo que pide la parte b).
- El `map[CodigoImpuesto]int` es lo que permite contar por código sin tener que declarar cuatro variables (una por A, B, C, D) a mano.
- Si `clientes` se termina antes de llegar a la meta, el `for` simplemente se acaba solo (no hace falta chequear el largo aparte), y `sinAtender` da 0.
