package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("randm number:", randomValue)

	randomValue = randomWithRange(3, 10)
	fmt.Println("randm number:", randomValue)

	randomValue = randomWithRange(10, 10)
	fmt.Println("randm number:", randomValue)

}

func randomWithRange(min int, max int) int {
	rand.Seed(time.Now().Unix())
	var value = rand.Int()%(max-min+1) + min
	return value
}

//deklarasi parameter bertipe data sama
// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type) returnType

func nameOfFunc(paramA int, paramB int) int
func randomWithRanges(min, max int) int
