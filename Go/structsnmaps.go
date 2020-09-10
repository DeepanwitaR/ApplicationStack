package main

import "fmt"

type Structure struct {
	number int
	word   string
}
type Structure2 struct {
	Structure
	data string `required max:"100"` //tag: conditions. https://youtu.be/YS4e4q9oBaU?t=9862
	id   int
} // embedded: sort of an inheritance feature. actually parent data are copied in child, composition.

//here other packs can see Structure but not see its field, we must capitalize to export them.
func boolreturn() bool {
	return true
}

//_ is a write only variable
func main() {
	//loops
	//for loop

	fmt.Println("for loop")
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Print(i)
		fmt.Print(j)
	}
	fmt.Println()
	// check out labels in Go, just like assembly language.
	array := []int{1, 2, 3, 4, 4, 5, 6}
	for k, v := range array {
		fmt.Printf("we have the index and value as: %v %v \n", k, v)
	}
}
func main2() {
	mapdata := map[string]int{
		"Dehli":   1,
		"Kolkata": 2,
		"Jersey":  3,
		"Osaka":   4,
	}
	_, ok := mapdata["Osaka"]
	fmt.Printf("Do we have the unknown data type: %v\n", ok)
	//we can have a conditional
	if pop, ok := mapdata["Osaka"]; ok {
		if 1 > 0 && 9 > 0 {
			fmt.Println("printing inner  block!")
		}
		fmt.Println("Osaka is present in the map and pop is:")
		fmt.Print(pop)
	}
	x := 2
	switch x {
	case 1, 5, 10:
		{
			fmt.Println(" \n printing 1\n")
		}
	case 2:
		{
			fmt.Println("\nprinting 2\n")
		}
	default:
		{
			fmt.Println("\nprinting the default block\n")
		}
	}
	//you may have an initializer
	switch i := x + 1; i {
	case 1, 5, 10:
		{
			fmt.Println(" \n printing 1\n")
			//fallthrough (what does it do?)
		}
	case 2: // you can have number<=10
		{
			fmt.Println("\nprinting 2\n")
		}
	default:
		{
			fmt.Println("\nprinting the default block\n")
		}
	}
	var typedata interface{} = 1 //interface can take in any data type.
	switch typedata.(type) {
	case int:
		fmt.Println("its an int")
	case bool:
		fmt.Println("its a bool")
	}
}
func main1() {
	fmt.Print("inside main1\n") // make can be involved here too.
	keymap := map[string]int{
		"stringA": 1,
		"stringB": 2,
	}
	keymap["stringC"] = 3
	delete(keymap, "stringB")
	fmt.Print(keymap)
	// an array is a valid key type, not slices
	// if you try to find a missing key, it returns 0.
	//map and slices, := leads to pointing to the same thing.
	_, ok := keymap["unknown"]
	structobj := Structure{
		number: 11,
		word:   "hello world",
	}
	// structobj := Structure{}
	structobj2 := Structure{
		111,
		"hello again!",
	} //causes probs if Structure is changed.
	structobj.number = 44

	aStructure := struct{ name string }{name: "anonymous"}

	fmt.Printf("here the ok value is: %v \n", ok)
	fmt.Printf("value of structobj: %v\n", structobj)
	fmt.Printf("value of structobj2: %v\n", structobj2)
	fmt.Printf("value of aStructure: %v\n", aStructure)
	fmt.Printf("value of number in structobj: %v\n", structobj.number)
}
