# Ejercicio 3 — Puntos Cardinales

## Enunciado

Declarar el tipo de datos **punto cardinal** como enumerativo. Leer un punto cardinal (N, S, E, O, NO, SE, NE, SO) e imprimir hacia cuál se dirige el viento. Encapsular en una función.

**a)** Resolver con `switch/case`.

**b)** Resolver aprovechando el orden de definición (el contrario se calcula por posición par/impar).

**c)** Resolver con un `Map` que tiene como índice el punto cardinal y cada elemento es su opuesto.

**d)** ¿Cómo se declaran enumerativos en otros lenguajes?

**e)** Definir la función `String()` que implementa la interfaz `Stringer` para usar con `fmt.Println`.

**f)** ¿Qué sucede cuando las funciones reciben un valor fuera de rango?

---

## Lógica de resolución

### Tipo enumerativo con iota

```go
type PuntoCardinal int

const (
    N PuntoCardinal = iota  // 0
    S                        // 1
    E                        // 2
    O                        // 3
    NE                       // 4
    SO                       // 5
    NO                       // 6
    SE                       // 7
)
```

### Parte a — switch/case

```go
func contrarioSwitch(p PuntoCardinal) PuntoCardinal {
    switch p {
    case N:  return S
    case S:  return N
    case E:  return O
    case O:  return E
    case NE: return SO
    case SO: return NE
    case NO: return SE
    case SE: return NO
    default: panic("punto cardinal inválido")
    }
}
```

### Parte b — usando el orden de definición

Los puntos están definidos en pares opuestos: (N,S), (E,O), (NE,SO), (NO,SE). El índice par y el siguiente impar son opuestos:
- Si p es par → opuesto = p+1
- Si p es impar → opuesto = p-1

```go
func contrarioOrden(p PuntoCardinal) PuntoCardinal {
    if p%2 == 0 {
        return p + 1
    }
    return p - 1
}
```

### Parte c — con Map

```go
var opuesto = map[PuntoCardinal]PuntoCardinal{
    N: S, S: N, E: O, O: E,
    NE: SO, SO: NE, NO: SE, SE: NO,
}

func contrarioMap(p PuntoCardinal) PuntoCardinal {
    return opuesto[p]
}
```

### Parte e — interfaz Stringer

```go
func (p PuntoCardinal) String() string {
    names := []string{"N", "S", "E", "O", "NE", "SO", "NO", "SE"}
    if int(p) < 0 || int(p) >= len(names) {
        return "inválido"
    }
    return names[p]
}
```

Con esto `fmt.Println(N)` imprime `"N"` en vez de `"0"`.

### Lectura desde entrada

```go
func leer() PuntoCardinal {
    var s string
    fmt.Scan(&s)
    names := map[string]PuntoCardinal{
        "N": N, "S": S, "E": E, "O": O,
        "NE": NE, "SO": SO, "NO": NO, "SE": SE,
    }
    p, ok := names[s]
    if !ok { panic("punto cardinal inválido: " + s) }
    return p
}
```

> **Parte f:** si la función recibe un valor fuera de rango (e.g., `PuntoCardinal(99)`), el switch no lo cubre y el default hace panic; el map devuelve el zero value (0 = N); la función de orden produce un resultado incorrecto. Conviene siempre validar la entrada.
