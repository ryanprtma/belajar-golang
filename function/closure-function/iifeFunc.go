package main

import "fmt"

func main() {
	var numbers = []int{2, 3, 4, 5, 0, 6, 7}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
				//keyword continue akan mengambil nilai selain e<min yang artinya e akan menghentikan e<min misalnya 2,0,1 (karena parameter yang diassign 3) dan mengembalikan nilai e>min
			}

			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number : ", numbers)
	fmt.Println("filtered number: ", newNumbers)
}
