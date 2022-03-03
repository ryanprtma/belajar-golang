package main

import (
	"fmt"
	"strings" //package untuk memanggil method join
)

func main() {
	names := []string{"ryan", "a"}
	printMessage("halo", names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")

	fmt.Println(message, nameString)
}
