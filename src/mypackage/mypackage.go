package mypackage

import (
	"fmt"
	"strings"
)

// CarPublic car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// isPalindromo car con acceso publico
func IsPalindromo(text string) {
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
