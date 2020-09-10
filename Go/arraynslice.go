package main

import "fmt"

func main() {
	fmt.Print("inside main\n")
	grades := [3][3]int{} // or [...]int{1,2,3}
	var grades1 [3]int
	fmt.Printf("grades array contain: %v\n", grades)
	fmt.Printf("grades1 array contain: %v\n", grades1)
	fmt.Printf("lenght of grades is: %v\n", len(grades)) //cap(a) func as well

	//a := b[3:] etc syntax
	//lets looks at a 2D array.
	var idmatrix [3][3]int = [3][3]int{[3]int{}, [3]int{}, [3]int{}}
	fmt.Println("idmatrix: \n")
	fmt.Println(idmatrix)
	fmt.Println(idmatrix[2])

	//you can copy arrays, but here references of 1st element of same memory
	// are not passed, there is a literal copy.
	array := []int{1, 2} //splice syntax
	array2 := array
	array2[1] = 3
	//if you just want to pass reference
	array3 := &array //both pointing to the same.
	//slice is like a chunk of an array; logically; you can manipulate the trait to add to the size.
	array4 := make([]int, 3)      //type of data, size
	array5 := make([]int, 3, 100) //type of data, size, capacity.
	array5 = append(array5, 2)
	array5 = append(array5, 2, 4, 6, 8)
	// there is something else called spreading.
	// append(array5, []int{2,3,4,5}...)
	fmt.Print(array3)
	fmt.Print(array4)
	fmt.Print("---")
	fmt.Print(array5)
}
