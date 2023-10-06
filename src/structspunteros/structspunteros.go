package structspunteros

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func (p pc) Saludar() {
	fmt.Println(p.brand, p.ram, "pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

// PunteroPublic punteros con acceso publico
func Punteros(name string) {
	a := 10
	b := *&a
	println(b)
	//fmt.Println(*b)
	// fmt.Println(a)
	newPc := pc{ram: 20, disk: 168, brand: "hp"}
	fmt.Println(newPc)
	fmt.Println(newPc.ram)
	newPc.Saludar()
	newPc.duplicateRam()
	fmt.Println(newPc.ram)
}
