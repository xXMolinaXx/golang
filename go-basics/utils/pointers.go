package utils

import "fmt"
func PointerExample(){
	// Create a pointer to an int
    var intPointer *int
    var x int = 10
	// Assign the address of x to the pointer
    intPointer = &x
	intPointer2 := &10
	// Dereference the pointer to get the value
    fmt.Println(*intPointer) // 10
	// Change the value through the pointer
	*intPointer = 20
    fmt.Println(*intPointer, x) // 20
    // Punteros a structs
    // var ptr *utils.Person = &p
    // fmt.Println(ptr.Age) // 30
	pointerFunctionExample(intPointer)
	fmt.Println(*intPointer, x) // 20
}
// Example of a function that takes a pointer as an argument
// get the value of the pointer and change it
func pointerFunctionExample(p *int){
	*p = 5
}