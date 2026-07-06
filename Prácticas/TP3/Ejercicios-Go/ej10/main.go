package main

import (
	"fmt"
	"sync"
)

type Contact struct {
	Nombre            string
	Apellido          string
	CorreoElectronico string
	Telefono          string
}

type Agenda struct {
	mu        sync.RWMutex
	contactos map[string]Contact
}

func NewAgenda() *Agenda {
	return &Agenda{contactos: make(map[string]Contact)}
}

func (a *Agenda) AgregarContacto(c Contact) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.contactos[c.CorreoElectronico] = c
}

func (a *Agenda) EliminarContacto(correo string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.contactos, correo)
}

func (a *Agenda) BuscarContacto(correo string) Contact {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.contactos[correo]
}

func main() {
	agenda := NewAgenda()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			agenda.AgregarContacto(Contact{
				Nombre:            fmt.Sprintf("Nombre%d", i),
				Apellido:          fmt.Sprintf("Apellido%d", i),
				CorreoElectronico: fmt.Sprintf("correo%d@mail.com", i),
				Telefono:          fmt.Sprintf("11-0000-%04d", i),
			})
		}(i)
	}
	wg.Wait()

	wg.Add(2)
	go func() {
		defer wg.Done()
		agenda.EliminarContacto("correo2@mail.com")
	}()
	go func() {
		defer wg.Done()
		c := agenda.BuscarContacto("correo3@mail.com")
		fmt.Println("Encontrado:", c)
	}()
	wg.Wait()

	fmt.Println("\nAgenda final:")
	for correo, c := range agenda.contactos {
		fmt.Printf("  %s -> %+v\n", correo, c)
	}
}
