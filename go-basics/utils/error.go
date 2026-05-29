package utils

import (
	"errors"
	"fmt"
)

func ExecError() {
	defer func(){
        fmt.Println("defer call in ExecError")
        
    }()
	// Llamada a un método asociado al struct
    /*********************************************/
    /*********************************************/ 
    // NOTE: ERROR
    /*********************************************/ 
    /*********************************************/ 
    var err error
    err = errors.New("This is an error")
    fmt.Println(err)
    // imprimir error a tipo string
    fmt.Println(err.Error())
    err2 := fmt.Errorf("This is another error: %s", "something went wrong")
    fmt.Println(err2)
    
    x:= 4
    y:= 0
    z := x / y
    fmt.Println("Result of division:", z)
    
}
// this is a throw error example, it will cause a panic and the program will crash if not handled
func PanicExample() {
	panic("This is a panic example")
}