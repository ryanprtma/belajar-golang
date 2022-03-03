package main

import "fmt"

func main() {
	//deklarasi variabel beserta tipe data && deklarasi variabel meunggunakan keyword var
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	fmt.Println("halo", firstName, lastName+"!")

	//deklarasi variabel tanpa tipe data

	var ryanA = "ryan"

	aPratama := "pratama"

	aPratama = "apratama"

	fmt.Printf("halo %s %s!\n", ryanA, aPratama)

	//deklarasi multi variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	//pengisian nilai bersamaan pada saat deklarasi
	var fourth, fifth, sixth string = "empat", "lima", "enam"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eight, ninth, one, isFriday, twoPointTwo, say)

	//variabel underscore atau variabel sampah
	_ = "belajar golang itu mudah"
	name, _ := "john", "wick"

	fmt.Printf("%s %s", name, say)

	//deklarasi variabel menggunakan keywword new

	nama := new(string)

	fmt.Println()
	fmt.Println(nama)
	fmt.Println(*nama)

	//deklarasi variabel menggunakan keyword make

	// Keyword ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu:

	// channel
	// slice
	// map

}
