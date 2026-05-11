package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Print("Ingrese una secuencia de caracteres: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	digitos := make(map[rune]int)

	for _, c := range texto {
		if unicode.IsDigit(c) {
			digitos[c]++
		}
	}

	for d := '0'; d <= '9'; d++ {
		fmt.Printf("'%c': %d\n", d, digitos[d])
	}
}
