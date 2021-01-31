package main

import (
	"fmt"
)

func changeSlice(sl []int) {
	for i := 0; i < len(sl); i++ {
		sl[i] += 1
	}

	sl = append(sl, -22)
}

func main() {
	array1 := make([]int, 0, 10)
	array2 := []int{0, 1, 2, 3, 4, 5, 6}

	array1 = append(array1, array2...)
	fmt.Println("cap = ", cap(array1), "  ", array1)
	changeSlicePtr(&array1)
	fmt.Println("cap = ", cap(array1), "  ", array1)

	//pointer := &array1[6]
	//memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array1[0])
	//pointer = (*int)(unsafe.Pointer(memoryAddress))
	//fmt.Printf("Array[%d] = %d \n",7,*pointer)

}

func changeSlicePtr(slptr *[]int) {
	slice := *slptr

	for i := 0; i < len(slice); i++ {
		slice[i] += 1
	}

	*slptr = append(slice, -22)
}
