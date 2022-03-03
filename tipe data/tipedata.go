package main

import "fmt"

func main() {
	//tipe data numerik non desimal
	var positiveNumber uint = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bialangan negatif: %d\n", negativeNumber)

	//%d digunakan untuk memformat data numerik non desimal

	//tipe data numerik desimal
	decimalNumber := 2.62

	fmt.Printf("bilangan desimal %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//%f digunakan untuk memformat data numerik desimal menjadi string %.nf = %.3f menghasilkan 3 angka dibelakang koma

	//tipe data boolean

	var exist bool = true
	nonExist := false

	fmt.Printf("exist? %t %t \n", exist, nonExist)

	//tipe data string

	message := `Nama saya "Ryan".
	Salam kenal.
	Mari belajar "Golang.`

	fmt.Println(message)

	//nilai nil dan zero value

	//variabel yang nilanya nil berarti memiliki nilai kosong
	//setiap tipe data yang dideklarasikan memiliki zero valuenya(default value)
	//string zero valuenya -> "" (string kosong)
	//bool -> false
	//numerik non desimal -> 0
	//numerik desimal - > 0.0

	//tipe data yang dapat diisi dengan nilai null
	//pointer, tipe data fungsi, slice, map, slice,channel, interface{}

}
