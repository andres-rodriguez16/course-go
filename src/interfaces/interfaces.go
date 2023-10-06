package interfaces

import "fmt"

type figuraGeometrica interface {
	area() float64
}

type Cuadrado struct {
	Base float64
}

type Reactangulo struct {
	base   float64
	altura float64
}

func (c Reactangulo) area() float64 {
	return c.base * c.altura
}
func (c Cuadrado) area() float64 {
	return c.Base * c.Base
}

func calcular(f figuraGeometrica) {
	fmt.Println("area", f.area())
}
func Interfaces() {
	fmt.Println("hello")
	myCuadrado := Cuadrado{Base: 40}
	myRectangulo := Reactangulo{base: 30, altura: 50}
	calcular(myCuadrado)
	calcular(myRectangulo)
	slice := []interface{}{1, 393, "sss", false}
	fmt.Println(slice...)
}
