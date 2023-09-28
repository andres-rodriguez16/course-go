package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valorOne := 4
	valortwo := 6

	// puerta logicas
	// &&
	if valorOne%2 == 0 {
		fmt.Println("es par")
	} else {
		fmt.Println("es impar")
	}
	// ||
	if valorOne == 1 || valortwo == 2 {
		fmt.Println("una condicion se cumplio")
	} else {
		fmt.Println("ambas condiciones no se cumplieron")
	}

	value, err := strconv.Atoi("7")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dweww", value)
	// switch

	switch number := 5 % 2; number {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("es impar")

	}
}
