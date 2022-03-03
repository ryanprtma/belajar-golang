package main

import "fmt"

type student struct {
	name  string
	grade int
}

type person struct {
	name string
	age  int
}

type students struct {
	grade int
	age   int
	person
}

//tag dimanfaatkan untuk encode/decode data json
type orangs struct {
	name string `tag1`
	age  int    `tag2`
}

//type alias
type People = person

//type alias untuk anon struct
type People1 = struct {
	name string
}

func main() {
	//deklarasi struct
	//hal ini sama seperti var s1 string
	//lalu var s1 diisikan valuenya
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	//inisilisaisi object
	var s2 = student{}
	s2.name = "wick"
	s2.grade = 2

	var s3 = student{"ethan", 2}

	var s4 = student{name: "ethan"}

	fmt.Println(s1.name, s2.name, s3, s4)

	s5 := student{name: "wick", grade: 2}

	var s6 *student = &s5

	fmt.Println(s5.name) //wick
	fmt.Println(s6.name) //untuk mengakses nilai asli pada variabel selain struct maka ditambahkan * didepan nama variabelnya, pada struct tidak perlu melakukan itu

	s6.name = "ethan"
	fmt.Println(s5.name, s6.name)

	//embedded struct
	//mekanisme menempelkan struct pada struct lain
	var s7 = students{}
	s7.name = "wick"
	s7.age = 21
	s7.grade = 2

	//untuk tiga kondisi dibawah s7.person.age memiliki memori yang sama dengan s7.age
	var s8 = s7.person.age
	s8 = 9
	s7.person.age = 9

	fmt.Println(s7, s8)

	//embedded struct dengan nama property yang sama
	//jika parent struct dan embedded struct memilii nama property yang sama, pengaksesan struct harus jelas

	s9 := students{}
	s9.name = "wick"
	s9.age = 21        //age of students
	s9.person.age = 22 //age of person

	fmt.Println(s9)

	//pengisian nilai sub struct
	//sub-struct bisa langsung diletakan pada struct parent asalkan strukturnya sudah dideklarasikan diatas
	var p1 = person{name: "wick", age: 21}
	var s10 = students{person: p1, grade: 2}

	fmt.Println(s10)

	//anonymous struct
	//anonymous struct adalah struct yang tidak didefinisikan diawal sebagai tipe data baru hal ini berguna ketika struct yang digunakan hanya sekali

	var s11 = struct {
		person     //inisialisasi struktur
		grade  int //inisialisasi struktur
	}{}

	//pengisian / pengaksesan anon struct
	s11.person = person{"wick", 21}
	s11.grade = 2

	fmt.Println(s11)

	//anon struct dengan pengisian langsung
	s12 := struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

	fmt.Println(s12)

	//deklarasi anon struct menggunakan keyword var
	//anon struct tidak didefinisikan sebagai type melainkan didefinisikan dengan variabel dan bertype anon

	//deklarasi
	//sama halnya dengan deklarasi type data yang lain, contoh :
	//deklarasi ->var huruf string
	//pengisian var huruf -> huruf = "contoh"
	//deklarasi sekaligus pengisian varriabel -> var huruf = "contoh" dengan tidak menuliskan lagi type datanya

	var student struct {
		person
		grade int
	}

	student.person = person{"wick", 21}

	//deklarasi sekalgus pengisian
	var students = struct {
		grade int
	}{
		12,
	}

	fmt.Println(students)

	//kombinasi slice dan struct

	var allStudents = []person{{
		name: "ryan", age: 22},
		{name: "dyah", age: 22},
		{name: "john", age: 45},
	}

	for _, student := range allStudents {
		fmt.Println(student.name)
	}

	fmt.Println(allStudents)

	//inisialisasi slice anon struct
	var allStudent = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}

	fmt.Println(allStudent)
}
