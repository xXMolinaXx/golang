package main

import "fmt"
import "example/example/utils"


func main() {
    mynotDefinedVariable := "Hello, World!" 
    var myVariable string = "Hello, World!"
    var anotherVariable string
    anotherVariable = "Hello, World!"
    // vectors
    var varVector [3]string
    varVector[0] = "Hello"
    varVector[1] = "World"
    varVector[2] = "!"
    var varVector2 = [3]string{"Hello", "World", "!"}
    //slices
    var varSlice []string = []string{"Hello", "World", "!"}
    fmt.Println("Hello, World!")
    fmt.Println(myVariable)
    fmt.Println(anotherVariable)
    fmt.Println(varVector)
    fmt.Println(varVector2)
    fmt.Println(varSlice)
    fmt.Println(mynotDefinedVariable)
    //maps
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
    res := utils.Add(5, 3)
    fmt.Println("Result of Add(5, 3):", res)

    // Crear e inicializar un struct
    p := utils.Person{Name: "Ana", Age: 30, City: "Madrid"}

    // Acceso a los campos
    fmt.Println(p.Name) // "Ana"

    
    e := utils.Employee{
        Name: "Luis",
        Address: utils.Address{
            Street: "Calle 1",
            City:   "Barcelona",
        },
    }
    
    // Punteros a structs
    var ptr *utils.Person = &p
    fmt.Println(ptr.Age) // 30
    
    // Slice de structs
    people := []utils.Person{
        {Name: "Juan", Age: 25, City: "Sevilla"},
        {Name: "Lucía", Age: 28, City: "Valencia"},
    }
    fmt.Println(people[0].Name) // "Juan"
    fmt.Println(e) // "Valencia"


    // these function ocurrus when the end, defer is used to execute a function after the main function has finished, even if there is a panic(error), recover give us the ability to handle the panic and prevent the program from crashing
    defer func(){
        fmt.Println("This is a function that will cause a panic due to division by zero")
        r:= recover() // recover captura el panic y permite manejarlo sin que el programa se detenga
        if r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    utils.ExecError()
    utils.PanicExample()
    
}