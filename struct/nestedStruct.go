package main

import "fmt"

type murid struct {
	name string
	age  int
}

type orang struct {
	suku   string
	murids murid
}

func main() {
	//nested struct
	//nested struct adalah anon struct yang di embed ke sebuah struct

	var murid2 orang
	murid2.suku = "melayu"
	murid2.murids.name = "ryan"
	murid2.murids.age = 22
	fmt.Println(murid2)

	murid3 := murid{
		"dyah",
		21,
	}

	murid4 := orang{
		suku:   "jawa",
		murids: murid3,
	}

	fmt.Println(murid3)
	fmt.Println(murid4)
}
