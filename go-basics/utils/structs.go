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


// this is a method of the struct Person, it can be called with a variable of type Person, it will print a greeting message with the name of the person
// explaining the sintaxis of the method, the receiver is the variable that will call the method, in this case p of type Person, the name of the method is Greet and it has no parameters and no return value
func (p Person) Greet() {
    fmt.Println("Hola, soy", p.Name)
}

func StructExample() {
    // Crear e inicializar un struct
    p := Person{Name: "Ana", Age: 30, City: "Madrid"}

    // Acceso a los campos
    fmt.Println(p.Name) // "Ana"

    
    e := Employee{
        Name: "Luis",
        Address: Address{
            Street: "Calle 1",
            City:   "Barcelona",
        },
    }
    fmt.Println(e) // "Barcelona"
    // Slice de structs
    people := []Person{
        {Name: "Juan", Age: 25, City: "Sevilla"},
        {Name: "Lucía", Age: 28, City: "Valencia"},
    }
    fmt.Println(people[0].Name) // "Juan"
}