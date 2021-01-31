package main

import "fmt"

func main() {

	a := 5

	//адрес в памяти - & (амперсанд)
	p := &a

	pp := &p

	fmt.Printf("value of a : %v | type of a : %T\n", a, a)
	fmt.Printf("value of p : %v | type of p : %T\n", p, p)
	fmt.Printf("value of pp : %v | type of pp : %T\n", pp, pp)

	//dereference - по адресу памяти получаете значение переменной

	//дай мне значение по адресу p
	b := *p //b = 5
	fmt.Printf("value of b : %v | type of b : %T\n", b, b)

	//------------------------------additional------------------------

	p0 := new(int)                       // p0 points to a zero int value.
	fmt.Println("hex address p0 : ", p0) // (a hex address string)
	fmt.Println(*p0)                     // 0

	// x is a copy of the value at
	// the address stored in p0.
	x := *p0
	// Both take the address of x.
	// x, *p1 and *p2 represent the same value.
	p1, p2 := &x, &x
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false
	p3 := &*p0            // <=> p3 := &(*p0) <=> p3 := p0
	// Now, p3 and p0 store the same address.
	fmt.Println(p0 == p3) // true
	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3) // 789 789 123

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}
