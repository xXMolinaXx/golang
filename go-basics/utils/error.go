package utils

import (
	// "errors"
	"fmt"
)

func ExecError() {
    // Defer function to recover from panic
    // this will be called when the function returns, either normally or through a panic
    // recover() is a built-in function that regains control of a panicking goroutine. It can only be used inside deferred functions. During normal execution, a call to recover() will return nil and have no other effect. If the current goroutine is panicking, a call to recover() will capture the value given to panic() and resume normal execution.
	defer func(){
        fmt.Println("defer call in ExecError")
        error := recover()
        if error != nil {
            fmt.Println("Recovered from panic:", error)
        }
    }()
    var valueX int
    var valueY int
    fmt.Println("Enter value for x:")
    fmt.Scanln(&valueX)
    fmt.Println("Enter value for y:")
    fmt.Scanln(&valueY)
    if valueY == 0 {
        panic("division by zero")
    }
    z := valueX / valueY
    fmt.Println("Result of division:", z)
    // var err error
    // err = errors.New("This is an error")
    // fmt.Println(err)
    // fmt.Println(err.Error())
    // err2 := fmt.Errorf("This is another error: %s", "something went wrong")
    // fmt.Println(err2)
}