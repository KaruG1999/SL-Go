package main

import (
	"fmt"
)

func FactorialIterativo (n int) int{
	res := 1
	for i:= 1; i<=n; i++ {
		res *= i
	}
	return res
}

func FactorialRecursivo (n int) int {
	// Caso base 
	if n==0 {
		return 1
	}
	return n* FactorialRecursivo(n-1)
}

func main (){

	for i := 0; i <= 9; i++ {
		iter := FactorialIterativo(i)
		recurs := FactorialRecursivo(i)
		// \t es para que quede alineado
		fmt.Printf("%d\t%d\t\t%d\n", i, iter,recurs)
	}

}