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

## Lógica de resolución (como está en `main.go`)

### Tipos

La cadena está resuelta directamente con lista enlazada (parte vi), guardando además un puntero al último nodo:

```go
type Billetera struct {
    ID       string
    Nombre   string
    Apellido string
}

type Transaccion struct {
    Monto      float64
    EmisorID   string
    ReceptorID string
    Timestamp  time.Time
}

type Bloque struct {
    Hash       string
    HashPrevio string
    Data       Transaccion
    Timestamp  time.Time
}

type nodo struct {
    elem Bloque
    sig  *nodo
}

type Blockchain struct {
    pri *nodo
    ult *nodo // puntero al último bloque, para no recorrer toda la cadena
    len int
}

func (bc *Blockchain) pushBack(b Bloque) {
    nuevo := &nodo{elem: b}
    if bc.pri == nil {
        bc.pri = nuevo
    } else {
        bc.ult.sig = nuevo
    }
    bc.ult = nuevo
    bc.len++
}
```

### Generar hash de un bloque

```go
func calcularHash(b Bloque) string {
    datos := fmt.Sprintf("%s|%s|%s|%.2f|%v|%v",
        b.HashPrevio, b.Data.EmisorID, b.Data.ReceptorID, b.Data.Monto, b.Data.Timestamp, b.Timestamp)
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

### iii) Insertar bloque (helper interno)

```go
func (bc *Blockchain) insertarBloque(tx Transaccion) {
    hashPrevio := ""
    if bc.ult != nil {
        hashPrevio = bc.ult.elem.Hash
    }
    nuevo := Bloque{
        HashPrevio: hashPrevio,
        Data:       tx,
        Timestamp:  time.Now(),
    }
    nuevo.Hash = calcularHash(nuevo)
    bc.pushBack(nuevo)
}
```

### ii + vii) Enviar transacción, con validación de saldo

```go
func (bc *Blockchain) EnviarTransaccion(tx Transaccion) error {
    if bc.ObtenerSaldo(tx.EmisorID) < tx.Monto {
        return errors.New("saldo insuficiente")
    }
    bc.insertarBloque(tx)
    return nil
}
```

`AcuñarFondos` es una función aparte, no pedida explícitamente por el enunciado, que le permite al sistema emitir moneda sin chequear saldo — así se cargan los fondos iniciales de cada billetera (equivalente a un bloque génesis):

```go
func (bc *Blockchain) AcuñarFondos(receptorID string, monto float64) {
    bc.insertarBloque(Transaccion{
        Monto: monto, EmisorID: "SISTEMA", ReceptorID: receptorID, Timestamp: time.Now(),
    })
}
```

### iv) Obtener saldo

```go
func (bc *Blockchain) ObtenerSaldo(id string) float64 {
    var saldo float64
    for n := bc.pri; n != nil; n = n.sig {
        if n.elem.Data.ReceptorID == id {
            saldo += n.elem.Data.Monto
        }
        if n.elem.Data.EmisorID == id {
            saldo -= n.elem.Data.Monto
        }
    }
    return saldo
}
```

### v) Validar consistencia

```go
func (bc *Blockchain) EsValida() bool {
    var anterior *nodo = nil
    for n := bc.pri; n != nil; n = n.sig {
        if n.elem.Hash != calcularHash(n.elem) {
            return false // el bloque fue alterado
        }
        if anterior != nil && n.elem.HashPrevio != anterior.elem.Hash {
            return false // el enlace con el anterior se rompió
        }
        anterior = n
    }
    return true
}
```

## Observaciones

- **El hash cubre también el timestamp de la transacción:** `calcularHash` incluye tanto `b.Timestamp` (fecha del bloque) como `b.Data.Timestamp` (fecha de la transacción). Al principio faltaba este último, así que alterar solo `Data.Timestamp` de un bloque ya creado no se detectaba en `EsValida()`; ya está corregido.
- **Por qué `Blockchain` guarda un puntero `ult`:** sin él, para insertar un bloque nuevo (necesitás el hash del último para armar `HashPrevio`) habría que recorrer toda la cadena desde `pri` cada vez. Con `ult`, tanto `pushBack` como leer el hash previo quedan en O(1) en vez de O(n). Es la diferencia concreta entre esta implementación y una lista enlazada "genérica" sin ese puntero extra.
- **`AcuñarFondos` no está en el enunciado:** se agregó para poder darle saldo inicial a una billetera sin que la validación de saldo suficiente lo bloquee (si no existiera, ninguna billetera podría empezar con dinero, porque toda transacción pasa por `EnviarTransaccion` y esa sí valida saldo).
