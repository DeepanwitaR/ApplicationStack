package main

import "fmt"

//primitive onwards
//defaults set to 0 and equivalent.
//diff int and float types. int and uint. complex type no.
//real(i) and imag(n). to make 5 + 18i, we make it complex(5, 18)
// study strings and type aliases. chars in string and their letter are their respective utf chars when
// typecasted or if we do str[3].
func main() {
	n := 1 == 1
	r := 'a' //rune check them up.
	var com complex64 = 2 + 2i
	//constant data types. var name starts with capital are exported.
	const myConst int = 42
	const a = 4
	var b int16 = 6
	// fun fact, ennumeration is enabled, but there is a counter with a int variable called iota.
	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	var (
		i int = 11
		j int
		k int
	)
	//same as
	// const (
	// 	a1 = iota
	// 	a2
	// 	a3
	// )
	// const (
	// 	; = iota
	// )
	// const (
	// 	_ = iota or iota + 5
	// )
	// var data byte = .... ( is a thing with %b)
	fmt.Printf("value: %v and type is: %T\n", n, n)
	fmt.Printf("value: %v and type is: %T\n", com, com)
	fmt.Printf("value: %v and type is: %T\n", r, r)
	fmt.Printf("value: %v and type is: %T\n", myConst, myConst)
	fmt.Printf("printing values if a + b: %v\n", a+b) //compiler at runtime pastes a as 42 and this auto conversion happens.

	fmt.Println("printing the iota block values.")
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)
	fmt.Printf("%v\n", a4)

	fmt.Println("printing i j k")
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)

}
