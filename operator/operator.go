package main

import "fmt"

func main() {
	//operator perbandingan
	value := (((2+6)%3)*4 - 2) / 3
	isEqual := (value == 2)

	fmt.Printf("nilai %d = %t \n", value, isEqual)

	//operator logika
	var left = false
	right := true

	leftAndRight := left && right

	leftOrRight := left || right

	notLeft := !left
	fmt.Print(leftAndRight, leftOrRight, notLeft)
}
