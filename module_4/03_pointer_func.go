package main

import "fmt"

func main() {
	a := 5

	multiplyIntBy(3, &a)

	fmt.Printf("value of a : %v | type of a : %T\n", a, a)
}

//go is most of the times PASS BY VALUE
func multiplyIntBy(multiplier int, i *int) {
	//i - копия из функции вызова
	//i = i * multiplier
	fmt.Printf("value of i : %v | type of i : %T\n", i, i)

	*i *= multiplier
	//значение по этому адресу = значение по этому адресу * multiplier
}
