package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	minc := strings.ToUpper(text)
	var textReversed string
	for i := len(minc) - 1; i >= 0; i-- {
		textReversed += string(minc[i])
	}
	if textReversed == minc {
		fmt.Println("Es un palidromo")
	} else {
		fmt.Println("No es un palidromo")
	}
}

func main() {
	// arrays: toca indicarle el tipado y la longitud, son inmutables
	var arr [4]int
	arr[0] = 1
	arr[1] = 5

	// leng() devuelve la longitud del arreglo y cap me duevuelve la capacidad maxima del arreglo
	fmt.Println(arr, len(arr), cap(arr))

	// slice: son mutables y no se les indican la longitud
	slice := []int{2, 4, 5, 67, 6, 6, 6, 9}
	fmt.Println(slice, len(slice), cap(slice))

	// metodos de slice, se utiliza con el valor del indice
	fmt.Println(slice[:2])  // 2, 4 // no incluye desde le indicador
	fmt.Println(slice[1:5]) // 4, 5, 67, 6,incluye desde el indicador
	fmt.Println(slice[1:])  // 4, 5, 67, 6, 6, 6 incluye desde el indicador

	// append

	slice2 := []int{100, 200}
	slice = append(slice, slice2...)
	// slice = append(slice, slice2...)
	fmt.Println(slice, "nuevo")

	// recorrer un slice

	slice3 := []string{"hola", "maria", "how are you"}

	for _, valor := range slice3 {
		fmt.Println(valor)
	}
	isPalindromo("Oso")

	for i, valor := range slice {
		fmt.Println(i, valor)
	}
}
