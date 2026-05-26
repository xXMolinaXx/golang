package utils
import "fmt"
// Definición básica de un struct
type Person struct {
    Name string
    Age  int
    City string
}



// Struct anidado
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Name    string
    Address Address
}


// Métodos asociados a structs
func (p Person) Greet() {
    fmt.Println("Hola, soy", p.Name)
}