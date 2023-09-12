package main

import "fmt"

func main() {
	// declaracion de constate
	const pi float64 = 2.14
	const pi2 float64 = 3.33333
	fmt.Println(pi, "pi")
	fmt.Println(pi2, "pi2")

	// declar variables enteras
	base := 22
	var name string = "Andres"
	var apellido = "rodriguez"

	fmt.Println(name + " " + apellido)
	fmt.Println(base)
	// zero values
	var a int
	var b float64
	var c bool
	var d string

	fmt.Println("a", a, "b", b, "c", c, "d", d)

	const base1 int = 10
	areaCuadrado := base1 * base1

	fmt.Println(areaCuadrado)
}
