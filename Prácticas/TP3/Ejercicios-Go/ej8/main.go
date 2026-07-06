package main

import "fmt"

func generar(ch chan<- int, valores []int) {
	for _, v := range valores {
		ch <- v
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go generar(ch1, []int{1, 2, 3})
	go generar(ch2, []int{10, 20, 30, 40})
	go generar(ch3, []int{100, 200})

	count1, count2, count3 := 0, 0, 0

	for ch1 != nil || ch2 != nil || ch3 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("ch1:", v)
			count1++
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("ch2:", v)
			count2++
		case v, ok := <-ch3:
			if !ok {
				ch3 = nil
				continue
			}
			fmt.Println("ch3:", v)
			count3++
		}
	}

	fmt.Printf("Totales - ch1: %d, ch2: %d, ch3: %d\n", count1, count2, count3)
}
