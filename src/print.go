package main

import (
	"fmt"
)

// func main() {
// 	helloMessage := "Hello"
// 	worldMessage := "world"

// 	// Println: Salto de Linea Automatico
// 	fmt.Println(helloMessage, worldMessage)
// 	fmt.Println(helloMessage, worldMessage)

// 	// Printf
// 	nombre := "Platzi"
// 	cursos := 500
// 	// Con valores seguros
// 	// expecipicar que tipo de dato va a resivir
// 	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
// 	// Con valores inseguros
// 	// cuando no sabemos que tipo de dato va a llegar
// 	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

// 	// Sprintf
// 	// convierte los valores que se le pacen en un string
// 	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
// 	fmt.Println(message)
// 	// para agregar salto de linea se usa \n
// 	// Tipo de datos:
// 	// podemos saber el tipo de con %T
// 	fmt.Printf("helloMessage: %T\n", helloMessage)
// 	fmt.Printf("cursos: %T\n", cursos)
// }

// normal
func hello(name string) {
	fmt.Printf("hello %s \n", name)
	fmt.Printf("%T", name)
}

// con varios parametros
func threeArgument(a, c string, b int) {
	fmt.Printf("\na=%s;b=%d;\nc=%s;", a, b, c)
}

// tipo de dato que retorna la funcion

func returnFunc(n int) int {
	return n
}

// return dos tipos de datos difentes
func returnDifferent(f int, x int) (a int, b int) {
	return f, x + x
}
func main() {
	// fmt.Println("hello world")
	// hello("andres")
	// threeArgurment("rodriguez", "andres", 20)
	// fmt.Println(returnFunc(3))
	valueOne, _ := returnDifferent(10, 20)
	fmt.Println(valueOne)
}
