package main

import "fmt"

func main() {
	//for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// i := 0

	// //for tanpa dengan argumen hanya kondisi
	// for i <= 5 {
	// 	fmt.Println("angka", i)
	// 	i++
	// }

	//for tanpa argumen
	// for {
	// 	fmt.Println("angka", i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	//for range digunakan untuk pengulangan pada array

	//penggunaan break dan continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
			//berhenti pada kondisi if i%2 ==  0
			//lanjut pada kondisi selain kondisi if i%2 == 0
		}

		if i > 8 {
			break
			//benar-benar berhenti pada kondisi i>8
		}

		fmt.Println("angka", i)
	}

	//perulangan bersarang
	for j := 0; j < 5; j++ {
		for k := j; k < 5; k++ {
			fmt.Print(k, "k")
		}
		fmt.Println("j")
	}

	//pemanfaatan label dalam perulangan
outerLoop:
	for l := 0; l < 5; l++ {
		for m := 0; m < 5; m++ {
			if l == 3 {
				break outerLoop
			}
			fmt.Print(l, m, "\n")
		}
	}
}
