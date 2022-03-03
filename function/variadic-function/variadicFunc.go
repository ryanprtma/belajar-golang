package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg = calculate(2, 3, 4, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("rata-rata: %.2f", avg)
	//fmt.Sprintf() mengembalikan nilai dalam bentuk string
	fmt.Println(msg)

	//pengisian parameter fungsi varidic menggunakan data slice
	var numbers = []int{2, 3, 4, 5, 4, 3, 3, 5, 5, 3}
	var average = calculate(numbers...)
	var message = fmt.Sprintf("rata-rata%.2f", average)

	fmt.Println(message)

	//func biasa + func variadic
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
	//float64() berguna untuk castng, teknik untuk konversi tipe data
	//operasi matematika hany abisa dilakukan ketika tipe data sejenis
}

//kombinasi fungsi biasa dengan fungsi variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("hello, my name is: %s\n", name)
	fmt.Printf("my hobbies are: %s\n", hobbiesAsString)
}
