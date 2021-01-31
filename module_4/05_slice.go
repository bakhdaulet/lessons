package main

import "fmt"

func main() {
	//var s1 = []int{0,0,0,0,0}
	s1 := []int{0, 0, 0, 0, 0}
	//s1 := make([]int, 5)

	//reSlice := s1[1:3] //3 не включительно

	//reSlice := s1[:3]
	reSlice := s1[2:]

	fmt.Println(s1)
	fmt.Println(reSlice)
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)

	//при создании вторичного среза не создается копия исходного среза.
	//базовый массив исходного среза будет сохраняться в памяти до тех пор, пока существует меньший, вторичный срез
}
