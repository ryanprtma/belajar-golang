package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	//names[4] baris ini menghsilkan error

	fmt.Println(names[0], names[1], names[2], names[3])

	//inisialisasi nilai awal array
	fruits := [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("jumlah element \t\t:", len(fruits))
	fmt.Println("isi semua element \t:", fruits)
	// \t digunakan untuk tab

	//penulisan inisialisasi array vertikal
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println(fruits)

	//inisialisasi array tanpa jumlah elemen
	numbers := [...]int{2, 2, 2, 2}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//array multidimensi
	numbers1 := [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}

	numbers2 := [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//perulangan array dnegna for
	buah := [...]string{
		"apel",
		"anggur",
		"pisang",
		"melon",
	}

	for i := 0; i <= len(buah)-1; i++ {
		fmt.Println(buah[i])
	}

	//perulangan elemen array menggunakan for-range
	for j, fruit := range buah {
		fmt.Println(j, fruit)
	}

	//penggunaan underscore dalam for-range
	for _, buahs := range buah {
		fmt.Println(buahs)
	}

	//alokasi elemen array menggunakan keyword make
	//deklarasi sekaligus alokasi dat array juga bisa dilakukan lewat keyword make

	//parameter pertama adalah tipe data elemen, panjang array

	var sayur = make([]string, 2)

	sayur[0] = "tomat"
	sayur[1] = "timun"

	fmt.Println((sayur))
}
