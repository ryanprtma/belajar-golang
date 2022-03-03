package main

import "fmt"

func main() {
	//mengambil nilai pointer pada variabel biasa dengan menambahkan & tepat sebelum nama variabel

	//nilai asli variabel pointer diambel dengan cara memanbahkan tanda * tepat sebelum nama variabel

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	numberA = 5
	fmt.Println("numberA (value)   :", numberA) // 5
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB) // 5
	fmt.Println("numberB (address) :", numberB)

	var angka = 4
	fmt.Println("before :", angka) // 4

	change(&angka, 10)
	fmt.Println("after  :", angka) // 10

	//angka(value(4))->&angka
	//original *int = &angka (address)
	//*original *int = value (value(10))

}

func change(original *int, value int) {
	*original = value
}
