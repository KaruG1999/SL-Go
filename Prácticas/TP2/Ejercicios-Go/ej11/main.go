package main

import (
	"fmt"
	ibt "miarbol"
)

func main() {
	a := ibt.New()
	fmt.Printf("¿Vacío?: %v\n", a.IsEmpty())

	for _, v := range []int{5, 3, 7, 1, 4, 6, 8} {
		a = a.Add(v)
	}

	fmt.Printf("InOrder:  %s\n", a.String())
	fmt.Printf("Len: %d | Depth: %d\n", a.Len(), a.Depth())

	// duplicado no debe insertarse
	a = a.Add(5)
	fmt.Printf("Len tras Add(5) duplicado: %d\n", a.Len())

	fmt.Print("PreOrder:  ")
	a.Traverse(func(v int) { fmt.Printf("%d ", v) }, ibt.PreOrder)
	fmt.Println()

	fmt.Print("PostOrder: ")
	a.Traverse(func(v int) { fmt.Printf("%d ", v) }, ibt.PostOrder)
	fmt.Println()

	fmt.Printf("Includes(4): %v\n", a.Includes(4))
	fmt.Printf("Includes(9): %v\n", a.Includes(9))

	fmt.Printf("Find(par): %v\n", a.Find(func(v int) bool { return v%2 == 0 }))
	fmt.Printf("Find(>10): %v\n", a.Find(func(v int) bool { return v > 10 }))

	a.Apply(func(v int) int { return v * 2 })
	fmt.Printf("Apply(x2): %s\n", a.String())
}
