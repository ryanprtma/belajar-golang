package main

import (
	"fmt"
	"strings"
)

type siswa struct {
	name  string
	grade int
}

type BukanSiswa struct {
	name  string
	grade string
}

func (property siswa) sayHello() {
	fmt.Println("halo", property.name)
}

func (property siswa) getNameAt(i int) string {
	return strings.Split(property.name, " ")[i-1]
}

func main() {
	s1 := siswa{"ryan apratama", 22}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println(name)

	s2 := BukanSiswa{
		"orang lain",
		"b",
	}
}
