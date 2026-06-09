package utils
import "fmt"

func Variables() {
	basic()
	arrays()
	slices()
	mapExample()
}


func basic(){
	fmt.Println("Ejemplo de variables y tipos de datos")
	mynotDefinedVariable := "Hello, World!" 
    var myVariable string = "Hello, World!"
    var anotherVariable string
    anotherVariable = "Hello, World!"
	fmt.Println("creacion de variables y asignacion de valores mynotDefinedVariable := \"Hello, World!\" ")
	fmt.Println(mynotDefinedVariable)
	fmt.Println("creacion de variables y asignacion de valores var myVariable string = \"Hello, World!\" ")
	fmt.Println(myVariable)
	fmt.Println("creacion de variables y asignacion de valores var anotherVariable string; anotherVariable = \"Hello, World!\" ")
	fmt.Println(anotherVariable)
}
func arrays(){
	// vectors
	fmt.Println("Ejemplo de arrays")
	fmt.Println("creacion de arrays var varVector [3]string; varVector[0] = \"Hello\"; varVector[1] = \"World\"; varVector[2] = \"!\" ")
	fmt.Println("asingacion de valores a arrays var varVector [3]string; varVector[0] = \"Hello\"; varVector[1] = \"World\"; varVector[2] = \"!\" ")
	fmt.Println("creacion de arrays y asignacion de valores var varVector2 = [3]string{\"Hello\", \"World\", \"!\"} ")
    var varVector [3]string
    varVector[0] = "Hello"
    varVector[1] = "World"
    varVector[2] = "!"
    var varVector2 = [3]string{"Hello", "World", "!"}
	fmt.Println(varVector)
	fmt.Println(varVector2)
}

func slices(){
	fmt.Println("Ejemplo de slices")
	fmt.Println("creacion de slices var varSlice []string = []string{\"Hello\", \"World\", \"!\"} ")
	//slices
    var varSlice []string = []string{"Hello", "World", "!"}
	//The make() function can also be used to create a slice.
	slice_number := make([]int, 3, 5) // creates a slice of type int with length 3 and capacity 5
	fmt.Println(slice_number)
	fmt.Println("creacion de slices con make() function slice_number := make([]int, 3, 5) ")
	fmt.Println(slice_number)
	slice_number[2] = 50 // change slice element at index 2 to 50
    varSlice = append(varSlice, "Go")
	fmt.Println(varSlice)
	varSlice = append(varSlice, "Programming")
    fmt.Println(varSlice)
    fmt.Println(len(varSlice))//len() function - returns the length of the slice (the number of elements in the slice)
    fmt.Println(cap(varSlice)) //cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	// Iterate over a slice using a for loop
	for idx, val := range varSlice {
		fmt.Printf("%v\t%v\n", idx, val)
	 }
}
// uper camel case is used to export a function, if the first letter of the function name is uppercase, it can be called from other packages, if it is lowercase, it can only be called from the same package
func mapExample(){
	//maps
	fmt.Println("Ejemplo de maps")
    var myMap map[string]string
    myMap = make(map[string]string)
    myMap["key1"] = "value1"
    myMap["key2"] = "value2"
    fmt.Println(myMap)
    delete(myMap, "key1")
    fmt.Println(myMap)
    value, exists := myMap["key2"]
    _, exists2 := myMap["key1"] // _ is used to ignore the value
    fmt.Println("Value:", value, "Exists:", exists)
    fmt.Println("Exists2:", exists2)
    res := Add(5, 3)
    fmt.Println("Result of Add(5, 3):", res)
}