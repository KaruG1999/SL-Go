package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
}

func New() Stack{
	return Stack{data:[]int{}}
}

func (s Stack) IsEmpty() bool{
	return len(s.data) == 0
}

func (s Stack) Len() int {
	return len(s.data)
}

func (s Stack) Top() (int, error){
	if s.IsEmpty() {
		return 0,errors.New("pila vacia")
	}
	return s.data[len(s.data)-1], nil
}

// El tope de la pila es el final 
func (s Stack) String() string {
	res := "tope -->"
	for i:= len(s.data) - 1; i>=0; i-- {
		res += fmt.Sprintf("[%d] ->", s.data[i])
	}
	return res
}

// Agrega atras 
func (s *Stack) Push (element int) {
	s.data = append(s.data, element)
} 

// Saca de atras
func (s *Stack) Pop()  (int, error){
	if s.IsEmpty(){
		return 0,errors.New("Pila vacia")
	}
	// sino guardo elemento a eliminar y decremento dimension
	top:= s.data[len(s.data)-1]
	// Slicing -> slice[inicio:fin] / Lenguaje asume que antes de : hay un 0 
	s.data = s.data[:len(s.data)-1]
	return top, nil
}

func (s *Stack) Iterate(f func(int) int) {
	for i:= range s.data {
		s.data[i] = f(s.data[i])
	}
}

func main() {
      s := New()
      s.Push(10)
      s.Push(20)
      s.Push(30)
      fmt.Printf("Pila: %s\n", s)

      top, err := s.Top()
      if err != nil {
          fmt.Println("Error:", err)
      } else {
          fmt.Printf("Tope: %d\n", top)
      }

      s.Iterate(func(n int) int { return n * 2 })
      fmt.Printf("Tras duplicar: %s\n", s)

      for !s.IsEmpty() {
          val, _ := s.Pop()
          fmt.Printf("Sacado: %d | Restante: %s\n", val, s)
      }

      _, err = s.Pop()
      fmt.Println("Pop en vacía:", err)
  }

