package firtclass

import "fmt"

func FirtClass() {
	// declaracion de constate
	const pi float64 = 2.14
	const pi2 float64 = 3.33333
	fmt.Println(pi, "pi")
	fmt.Println(pi2, "pi2")

	// declar variables enteras
	base := 22
	var name string = "Andres"
	var lastName = "rodriguez"

	fmt.Println(name + " " + lastName)
	fmt.Println(base)
	// zero values
	/*var a int
	var b float64
	var c bools
	var d string
	fmt.Println("a", a, "b", b, "c", c, "d", d)
	*/

	const base1 int = 10
	// multiplicacion
	areaCuadrado := base1 * base1
	fmt.Println(areaCuadrado)
	// sumna
	y := 10
	z := 70
	result := y + z
	fmt.Println("suma", result)
	// resta
	result = y - z
	fmt.Println("resta", result)

	// diivsion
	result = y / z
	fmt.Println("division", result)
	// modulo
	result = y % z
	fmt.Println("modulo", result)
}

// Numeros enteros
//int = Depende del OS (32 o 64 bits)
//int8 = 8bits = -128 a 127
//int16 = 16bits = -2^15 a 2^15-1
//int32 = 32bits = -2^31 a 2^31-1
//int64 = 64bits = -2^63 a 2^63-1

//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
//uint = Depende del OS (32 o 64 bits)
//uint8 = 8bits = 0 a 127
//uint16 = 16bits = 0 a 2^15-1
//uint32 = 32bits = 0 a 2^31-1
//uint64 = 64bits = 0 a 2^63-1

//numeros decimales
// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

//textos y booleanos
//string = ""
//bool = true or false

//numeros complejos
//Complex64 = Real e Imaginario float32
//Complex128 = Real e Imaginario float64
//Ejemplo : c:=10 + 8i
