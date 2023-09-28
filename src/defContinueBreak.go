package main

import "fmt"

func main() {
	// defer
	defer fmt.Println("termina la funcion")
	fmt.Println("inicia la funcion")
	// continue
	for i := 0; i <= 5; i++ {
		if i == 4 {
			fmt.Println("es 4", i)
			continue
		}
	}
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("i:", i)
			break
		}
	}
}
