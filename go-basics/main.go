package main

import "fmt"
import "example/example/utils"


func main() {
    var opcion string
    fmt.Println(`Menu seleciona ejemplo a ver:
    1. Variables y tipos de datos
    2. Structs
    3. Conversions
    4. Net package
    5. Reflection
    6. External packages
    7. Error handling
    8. Pointers`)
    fmt.Scanln(&opcion)
    switch opcion {
    case "1":
        utils.Variables()
    case "2":
        utils.StructExample()
    case "3":
        utils.Convert()
    case "4":
        utils.NetExample()
    case "5":
        utils.ReflectExample()
    case "6":
        utils.ExternalPackageExample()
    case "7":
        utils.ExecError()
    case "8":
        utils.PointerExample()
    default:
        panic("opcion invalida") // asi se tirar errores
    }
    
    return
    
}