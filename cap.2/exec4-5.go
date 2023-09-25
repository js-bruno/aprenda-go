package main

import "fmt"

type celsius int

var ye int

func UsingTypes() {
	var x celsius

	fmt.Printf("%T", x)
	fmt.Println(x)

	x = 42

	ye = int(x)
	fmt.Println(ye)
	fmt.Printf("%T", ye)
}
