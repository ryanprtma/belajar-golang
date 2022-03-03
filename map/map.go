package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}

	//deklarasi map kosong
	person := map[string]string{}

	//menambahkan data pada map
	person["title"] = "test"

	//megakses data pada map
	fmt.Println(person["title"])
	fmt.Println(person)

	//deklarasi map berisi
	orang := map[string]string{
		"nama": "ryan",
		"umur": "19",
	}

	fmt.Println(orang)

	//membuat map
	book := make(map[string]string)
	book["title"] = "test 1"
	book["author"] = "test 1"
	book["wrong"] = "salah"

	fmt.Println(book)

	//menghapus data pada map
	delete(book, "wrong")
	fmt.Println(book)

	//menghitung panjang map
	fmt.Println(len(book))

	//iterasi item map menggunakan for - range

	chicken := map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	//deteksi keberadaan item dengan key tertentu

	value, isExist := chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	//kombinasi slice dan map

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chichicken := range chickens {
		fmt.Println(chichicken["gender"], chichicken["name"])
	}

	//mempersingkat map
	var ayam = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	fmt.Println(ayam)

	//map dengan key yang berbeda
	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	data1 := data[0:2]

	fmt.Println(data1)

}
