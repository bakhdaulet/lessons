package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [5]int{0, 1, -2, 3, 4}

	//slice
	//array := []int{0, 1, -2, 3, 4}

	//fmt.Println(array[5])

	pointer := &array[0]
	fmt.Println(*pointer)
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Println(*pointer)
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer)
	fmt.Println()
}
