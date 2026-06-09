package utils
import "runtime"
import "fmt"
import "os"

func OsExample() {
	// Obtener el nombre del sistema operativo
	osName := runtime.GOOS
	fmt.Println("Operating System:", osName)
	file , err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1) // Salir con un código de error
	}
	c, err := file.Read(make([]byte, 100))
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1) // Salir con un código de error 
	}
	fmt.Printf("Read %d bytes from file: %s\n", c, "file.txt")
	env1 := os.Getenv("HOME")
	fmt.Println("HOME environment variable:", env1)
	os.Setenv("MY_VAR", "Hello, World!")
	env2 := os.Getenv("MY_VAR")
	fmt.Println("MY_VAR environment variable:", env2)
	//  concatena el valor de MY_VAR con una cadena y lo asigna a dbUrl
	dbUrl := os.ExpandEnv("DB_URL=localhost:5432 ${MY_VAR}")
	fmt.Println("DB_URL environment variable:", dbUrl)
	defer file.Close()
}