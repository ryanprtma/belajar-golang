package main

import "fmt"

func main() {
	//closure adalah function yang bisa disimpan didalm sebuah variabel dalam bentuk func anonymous

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
				fmt.Print(max, min)
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\n min:%v\n max : %v\n", numbers, min, max)

	//penggunaan template string %v
	// Template %v digunakan untuk menampilkan segala jenis data. Bisa array, int, float, bool, dan lainnya. Bisa dilihat pada statement di atas, data bertipe array dan numerik ditampilkan menggunakan %v. Template ini biasa dimanfaatkan untuk menampilkan sebuah data yang tipe nya bisa dinamis atau belum diketahui. Sangat tepat jika digunakan pada data bertipe interface{}
}
