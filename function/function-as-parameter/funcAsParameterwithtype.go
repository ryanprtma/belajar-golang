package main

import (
	"fmt"
	"strings"
)

//alias skema closure
//menggunakan type alias untuk mempersingkat
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
	var result []string

	for _, each := range data {
		filtered := callback(each) //true or false
		if filtered {
			result = append(result, each)
		}
	}
	return result
}

func cariHuruf(each string) bool {
	return strings.Contains(each, "o") //akan mengembalikan true or false
}

func main() {
	data := []string{"wick", "jason", "ethan"}

	var dataContainsO = filter(data, cariHuruf)

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
}
