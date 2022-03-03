package main

import "fmt"

func main() {
	//slice adalah reference elemen array. SLice bisa dibuat atau juga bisa dihasilkan dari manipulasi array maupun slice lainnya. Perubahan data disetiap elemen berdampak pada slice lain

	//inisialisasi slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	//pada inisialisasi slice jumlah element tidak dideklarasikan

	fmt.Println(fruits[0])

	// /mengakses slice
	var newFruits = fruits[0:2]
	fmt.Println(newFruits) // ["apple", "grape"]

	//slice merupakan tipe data reference
	aFruits := fruits[0:3] //["apple", "grape", "banana"]
	bFruits := fruits[1:4] //["grape", "banana", "melon"]

	aaFruits := aFruits[1:2] //["grape"]
	baFruits := bFruits[0:1] //["grape"]

	//buah grape diubah menjadi pinaple
	baFruits[0] = "pineaple"

	fmt.Println(fruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	//fungsi len untuk menghitung jumlah elemen pada slice
	fmt.Println(len(fruits))

	//fungsi cap untuk menghitung lebar atau kapasitas maksimum slice
	fmt.Println(len(fruits)) //len : 4 [buah buah buah buah]
	fmt.Println(cap(fruits)) // cap: 4

	cFruits := fruits[0:3]
	fmt.Println(len(cFruits)) //len : 3
	fmt.Println(cap(cFruits)) // cap: 4 [buah buah buah ----]

	dFruits := fruits[1:4]
	fmt.Println(len(dFruits)) //len : 3
	fmt.Println(cap(dFruits)) // cap: 3 ---- [buah buah buah]

	//fungsi append
	eFruits := append(fruits, "papaya")
	fmt.Println(fruits)  // ["apple", "pinaple", "banana", "melon"]
	fmt.Println(eFruits) // ["apple", "pinaple", "banana", "melon", "papaya"]

	//kasus diatas ketika len (fruits) == cap (fruits) maka akan mengembalikan refernsi baru atau slice baru

	fFruits := fruits[0:2]

	fmt.Println(cap(fFruits)) // 4
	fmt.Println(len(fFruits)) // 2

	fmt.Println(fruits)  // ["apple", "pinaple", "banana", "melon"]
	fmt.Println(fFruits) // ["apple", "pinaple"]

	gFruits := append(fFruits, "papaya")
	fmt.Println(fruits)  // ["apple", "pineaple", "papaya", "melon"]
	fmt.Println(fFruits) // ["apple", "pineaple"]
	fmt.Println(gFruits) // ["apple", "pinaple", "papaya"]

	//ketika len(fFruits)<cap(fFruits) maka elemen yang akan diappend akan mengikuti slice terbaru fFruits := fruits[0:2] yang artinya "papaya" akan diisikan pada 2

	//fungsi copy
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	//contoh ke-2
	dst2 := []string{"potato", "potato", "potato"}
	src2 := []string{"watermelon", "pineaple"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2)     //watermelon pinnaple potato
	fmt.Println(src2)     //watermelon pinnaple
	fmt.Println(dst2, n2) //2

	//mengakses elemen slice dengan 3 indeks
	var nFruits = fruits[0:2]
	var mFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(nFruits)      // ["apple", "grape"]
	fmt.Println(len(nFruits)) // len: 2
	fmt.Println(cap(nFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(mFruits)) // len: 2
	fmt.Println(cap(mFruits)) // cap: 2

}
