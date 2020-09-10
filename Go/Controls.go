package main

import "fmt"

func main() {
	fmt.Println("in the main func1")
	//defer works in th stack order.
	defer fmt.Println("start")
	defer fmt.Println("middle")
	// panic("Function hit an error")
	defer fmt.Println("end")
	fmt.Println("in the main func2")
	// you may induce a panic, an exit error.
	// defer with anonymous function. https://youtu.be/YS4e4q9oBaU?t=14173
	//recover is like a catch for panic; or rethrowing panic
	//variatic parameter in func declarations.
	//returning pointers pointing to vars in the local scope. https://youtu.be/YS4e4q9oBaU?t=16806
	//annonymous function
	func() {
		fmt.Println("inside the annonymous func")
	}() // what is this for.
	structA := myStruct{}
	structA.name = "Structure Name"
	structA.number = 1234
	structA.StructMethod()
}

//methods, are any func executing in a known context.
type myStruct struct {
	number int
	name   string
}

func (m myStruct) StructMethod() {
	fmt.Println("I am the method for myStruct")
}
