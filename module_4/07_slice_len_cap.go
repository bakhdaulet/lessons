package main

import (
	"fmt"
)

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func main() {
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	bSlice := []int{-2, -3, -4}
	//почему в append надо передовать сам aSlice
	aSlice = append(aSlice, bSlice...)

	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
}
