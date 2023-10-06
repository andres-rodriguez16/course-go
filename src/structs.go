package main

import "fmt"

type car struct {
	brand string
	year  int
}
type Product struct {
	name     string
	price    float64
	quantity uint32
	brand    string
	genres   []string
}

func main() {
	myCar := car{brand: "ford", year: 10}
	slice := []car{}
	slice = append(slice, myCar)

	fmt.Println(myCar)
	// otra manera
	var otherCar car
	otherCar.brand = "ferrary"
	otherCar.year = 2222
	slice = append(slice, otherCar)
	fmt.Println(slice)
}
