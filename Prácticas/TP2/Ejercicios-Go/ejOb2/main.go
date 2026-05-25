package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

// ── Tipos de dominio ──────────────────────────────────────────────────────────

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

// ── Lista enlazada de Bloque (vi: reemplaza al slice) ─────────────────────────
// Guardamos puntero al último nodo para que PushBack sea O(1).

type nodo struct {
	elem Bloque
	sig  *nodo
}

type Blockchain struct {
	pri *nodo
	ult *nodo
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

// ── i) Crear billetera ────────────────────────────────────────────────────────

func NuevaBilletera(id, nombre, apellido string) Billetera {
	return Billetera{ID: id, Nombre: nombre, Apellido: apellido}
}

// ── Hash ──────────────────────────────────────────────────────────────────────

func calcularHash(b Bloque) string {
	datos := fmt.Sprintf("%s|%s|%s|%.2f|%v",
		b.HashPrevio, b.Data.EmisorID, b.Data.ReceptorID, b.Data.Monto, b.Timestamp)
	hash := sha256.Sum256([]byte(datos))
	return fmt.Sprintf("%x", hash)
}

// ── iv) Obtener saldo ─────────────────────────────────────────────────────────

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

// ── insertar bloque interno ───────────────────────────────────────────────────

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

// ── ii+iii) Enviar transacción (con validación de saldo) ─────────────────────

func (bc *Blockchain) EnviarTransaccion(tx Transaccion) error {
	// vii) validar saldo suficiente
	if bc.ObtenerSaldo(tx.EmisorID) < tx.Monto {
		return errors.New("saldo insuficiente")
	}
	bc.insertarBloque(tx)
	return nil
}

// AcuñarFondos permite al sistema emitir moneda sin validación de saldo.
// Representa el bloque génesis o depósito inicial del banco.
func (bc *Blockchain) AcuñarFondos(receptorID string, monto float64) {
	bc.insertarBloque(Transaccion{
		Monto:      monto,
		EmisorID:   "SISTEMA",
		ReceptorID: receptorID,
		Timestamp:  time.Now(),
	})
}

// ── v) Validar consistencia de la cadena ─────────────────────────────────────

func (bc *Blockchain) EsValida() bool {
	var anterior *nodo = nil
	for n := bc.pri; n != nil; n = n.sig {
		// el hash del bloque debe coincidir con su contenido
		if n.elem.Hash != calcularHash(n.elem) {
			return false
		}
		// el hash previo debe coincidir con el hash del bloque anterior
		if anterior != nil && n.elem.HashPrevio != anterior.elem.Hash {
			return false
		}
		anterior = n
	}
	return true
}

// ── Utilidades de impresión ───────────────────────────────────────────────────

func (bc *Blockchain) Imprimir() {
	fmt.Printf("Blockchain (%d bloques):\n", bc.len)
	for n := bc.pri; n != nil; n = n.sig {
		b := n.elem
		fmt.Printf("  Hash:     %.16s...\n", b.Hash)
		fmt.Printf("  Previo:   %.16s...\n", b.HashPrevio)
		fmt.Printf("  TX:       %s -> %s  $%.2f\n", b.Data.EmisorID, b.Data.ReceptorID, b.Data.Monto)
		fmt.Println("  ──────────────────────────")
	}
}

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	// i) crear billeteras
	alice := NuevaBilletera("alice", "Alice", "Smith")
	bob := NuevaBilletera("bob", "Bob", "Jones")

	bc := Blockchain{}

	// Fondos iniciales: el sistema acuña moneda para cada billetera
	fmt.Println("=== Cargando fondos iniciales ===")
	bc.AcuñarFondos(alice.ID, 1000)
	fmt.Printf("Fondos a %s: $1000\n", alice.Nombre)
	bc.AcuñarFondos(bob.ID, 500)
	fmt.Printf("Fondos a %s: $500\n", bob.Nombre)

	fmt.Println("\n=== Transacciones entre usuarios ===")
	// alice -> bob: 200
	err := bc.EnviarTransaccion(Transaccion{Monto: 200, EmisorID: alice.ID, ReceptorID: bob.ID, Timestamp: time.Now()})
	fmt.Printf("Alice -> Bob $200: %v\n", err)

	// bob -> alice: 9999 (saldo insuficiente)
	err = bc.EnviarTransaccion(Transaccion{Monto: 9999, EmisorID: bob.ID, ReceptorID: alice.ID, Timestamp: time.Now()})
	fmt.Printf("Bob -> Alice $9999: %v\n", err)

	fmt.Println("\n=== Saldos ===")
	fmt.Printf("Saldo %s (%s): $%.2f\n", alice.Nombre, alice.ID, bc.ObtenerSaldo(alice.ID))
	fmt.Printf("Saldo %s (%s): $%.2f\n", bob.Nombre, bob.ID, bc.ObtenerSaldo(bob.ID))

	fmt.Println("\n=== Cadena ===")
	bc.Imprimir()

	fmt.Println("=== Validación ===")
	fmt.Printf("Cadena válida: %v\n", bc.EsValida())

	// Tamperear un bloque para probar que la validación lo detecta
	bc.pri.elem.Data.Monto = 9999
	fmt.Printf("Cadena válida tras tampereo: %v\n", bc.EsValida())
}
