package stringers

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc Pc) String() string {
	return fmt.Sprintf("tengo un compuador %d GB Ram, %d GB y es un %s", myPc.ram, myPc.disk, myPc.brand)
}
func Stringers() {
	myPc := Pc{ram: 6, disk: 156, brand: "hp"}
	fmt.Println(myPc.String())
}
