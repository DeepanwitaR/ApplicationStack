package main

import (
	"fmt"
	"strconv"
)

//import abc?

var one int = 9
var two int

//can be as: in pacakage level only this syntax.

var (
	oneone int = 99
	twotwo int = 88
)

// var I and var i
// j = float32(i)
func main() {
	// int i = 10;
	var i int = 10
	// var b bool = true
	n := 1 == 1

	var str string
	str = strconv.Itoa(i)
	ie := 11 //leave it for the compiler to figure out and assign data. float32 and float64 example.
	fmt.Println("hello there")
	fmt.Println(n)
	fmt.Print(ie + i)
	fmt.Println(str)
	fmt.Printf("the type is %T and value is %v", oneone, oneone)
}
