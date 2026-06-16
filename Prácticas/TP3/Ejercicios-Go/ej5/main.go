package main

import "fmt"

func ping(pingCh chan struct{}, pongCh chan struct{}) {
	for i := 0; i < 4; i++ {
		<-pingCh             // Espera la ficha para jugar
		fmt.Println("PING")
		pongCh <- struct{}{} // Le pasa la ficha a PONG
	}
}

func pong(pingCh chan struct{}, pongCh chan struct{}, done chan bool) {
	for i := 0; i < 4; i++ {
		<-pongCh             // Espera la ficha para jugar
		fmt.Println("PONG")
		pingCh <- struct{}{} // Le devuelve la ficha a PING (así ping no se traba)
	}
	done <- true // Avisa al main que el juego terminó
}

func main() {
	// Canales de control 
	pingCh := make(chan struct{}, 1)
	pongCh := make(chan struct{})
	done := make(chan bool)

	// Inicialización: metemos la primera ficha para que arranque PING
	pingCh <- struct{}{}

	// Lanzamos los jugadores en segundo plano pasándole sus dependencias
	go ping(pingCh, pongCh)
	go pong(pingCh, pongCh, done)

	// Freno de mano: el main espera acá hasta que pong avise que terminó
	<-done
}