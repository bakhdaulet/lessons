package main

import (
	"fmt"
	"strconv"
)

//maps in go are hashmaps

func main() {
	//m := make(map[string]int)
	//var m = map[string]int{}

	var m = map[string]int{
		"vasya": 123213,
		"kolya": 321321,
	}

	m["petya"] = 424242
	m["egor"] = 0

	//foreach
	for key, value := range m {
		fmt.Println(key + " " + strconv.Itoa(value))
	}

	delete(m, "kolya")

	//проверка есть ли значение в map
	if val, ok := m["kolya"]; ok {
		fmt.Println("val is : ", val)
	} else {
		fmt.Println("no exits but val is : ", val)
	}

}
