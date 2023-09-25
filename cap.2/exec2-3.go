package main

import "fmt"

var (
	x int
	y string
	z bool
)

func UsingPackageLevelVars() {
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprint(x, y, z)
	fmt.Println(s)
}
