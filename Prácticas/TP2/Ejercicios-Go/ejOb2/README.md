# Ejercicio Obligatorio 2 — Blockchain

## Enunciado

Implementar una blockchain para una criptomoneda que incluya billeteras para clientes. Una blockchain es un libro de contabilidad público e inmutable donde cada transacción se agrupa en un bloque enlazado con el anterior (similar a una lista enlazada).

**Structs requeridos:**
- `Transaccion`: monto, ID emisor, ID receptor, timestamp
- `Bloque`: hash, hash previo, transacción (data), timestamp de creación
- `Blockchain`: la cadena
- `Billetera`: ID, nombre, apellido del cliente

**b)** Funciones para:
- i) Crear una billetera
- ii) Enviar una transacción
- iii) Insertar un bloque en la cadena
- iv) Obtener el saldo de un usuario (recorriendo toda la cadena)
- v) Validar que la cadena sea consistente (verificar hashes)
- vi) Re-implementar con lista enlazada (ejercicio 9) en vez de slice
- vii) Validar que el emisor tenga saldo suficiente

**Tip:** usar `crypto/sha256` para generar el hash del bloque.

---

## Lógica de resolución

### Tipos

```go
import (
    "crypto/sha256"
    "fmt"
    "time"
)

type Billetera struct {
    ID       string
    Nombre   string
    Apellido string
}

type Transaccion struct {
    Monto    float64
    EmisorID string
    ReceptorID string
    Timestamp time.Time
}

type Bloque struct {
    Hash       string
    HashPrevio string
    Data       Transaccion
    Timestamp  time.Time
}

// Versión con slice
type Blockchain struct {
    bloques []Bloque
}
```

### Generar hash de un bloque

```go
func calcularHash(b Bloque) string {
    datos := fmt.Sprintf("%s%s%v%v", b.HashPrevio, b.Data.EmisorID, b.Data.Monto, b.Timestamp)
    hash := sha256.Sum256([]byte(datos))
    return fmt.Sprintf("%x", hash)
}
```

### i) Crear billetera

```go
func NuevaBilletera(id, nombre, apellido string) Billetera {
    return Billetera{ID: id, Nombre: nombre, Apellido: apellido}
}
```

### ii y iii) Enviar transacción / Insertar bloque

```go
func (bc *Blockchain) EnviarTransaccion(tx Transaccion) error {
    // vii) Validar saldo
    if bc.ObtenerSaldo(tx.EmisorID) < tx.Monto {
        return fmt.Errorf("saldo insuficiente")
    }

    hashPrevio := ""
    if len(bc.bloques) > 0 {
        hashPrevio = bc.bloques[len(bc.bloques)-1].Hash
    }

    nuevo := Bloque{
        HashPrevio: hashPrevio,
        Data:       tx,
        Timestamp:  time.Now(),
    }
    nuevo.Hash = calcularHash(nuevo)
    bc.bloques = append(bc.bloques, nuevo)
    return nil
}
```

### iv) Obtener saldo

Recorrer toda la cadena sumando entradas y restando salidas:

```go
func (bc *Blockchain) ObtenerSaldo(id string) float64 {
    var saldo float64
    for _, b := range bc.bloques {
        if b.Data.ReceptorID == id {
            saldo += b.Data.Monto
        }
        if b.Data.EmisorID == id {
            saldo -= b.Data.Monto
        }
    }
    return saldo
}
```

### v) Validar consistencia

```go
func (bc *Blockchain) EsValida() bool {
    for i := 1; i < len(bc.bloques); i++ {
        actual  := bc.bloques[i]
        anterior := bc.bloques[i-1]

        if actual.Hash != calcularHash(actual) {
            return false   // hash del bloque corrupto
        }
        if actual.HashPrevio != anterior.Hash {
            return false   // enlace roto
        }
    }
    return true
}
```

### vi) Re-implementar con lista enlazada

Reemplazar `[]Bloque` por la `List` del ejercicio 9 (adaptada para `Bloque`). El comportamiento es el mismo; el impacto es que `PushBack` es O(1) si la lista guarda puntero al último nodo, pero acceder al último bloque (para obtener `HashPrevio`) sigue siendo O(1) con ese puntero. La iteración para calcular saldo sigue siendo O(n).

> La blockchain funciona como una lista enlazada donde cada nodo conoce el hash del nodo anterior — cualquier modificación en un bloque intermedio invalida todos los posteriores, garantizando la inmutabilidad.
