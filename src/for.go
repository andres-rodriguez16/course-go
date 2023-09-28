package main

import "fmt"

func for() {
	// for condition
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("\n")
	// while loop
	// counter := 0
	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }
	//for inverted
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}

}
