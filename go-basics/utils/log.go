package utils
import "log"
import "os"
func LogExample(){
	log.Println("This is a log example")
	// log.Fatal("This is a fatal log example") // This will exit the program
	// log.Panic("This is a panic log example") // This will panic the program
	file, err := os.Open("file.txt")
	if err != nil {
		log.Printf("Error opening file: %v", err)
	}
	l := log.New(file, "", 0)
	l.Println("This is a log example with a custom logger")
	l.Printf("This is a log example with a custom logger and a formatted message: %s", "Hello, World!")
}