package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Andres"] = 20
	m["camilo"] = 50

	fmt.Println(m)
	// recorrer un map
	for i, valor := range m {
		fmt.Println(i, valor)
	}
	// Encontar un valo
	_, ok := m["Camilo"]
	fmt.Println(ok)

}
