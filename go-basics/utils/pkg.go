package utils
import "strconv"
import "fmt"
import "net/http"
import "io"
import "reflect"
func Convert(){
	s := strconv.Itoa(-42)
	fmt.Println(s)
	s = strconv.FormatBool(true)
	fmt.Println(s)
	s = strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Println(s)
	s = strconv.FormatInt(-42, 10)
	fmt.Println(s)
	s = strconv.FormatUint(42, 10)
	fmt.Println(s)

	i, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println(i)
	}

	f, err := strconv.ParseFloat("3.14159", 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println(f)
	}
}

func NetExample(){
	resp, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response body:", string(body))
}

func ReflectExample(){
	// Reflection allows us to inspect the type and value of variables at runtime, it is useful for tasks like serialization, deserialization, and dynamic function calls.
	t := reflect.TypeOf(42)
	fmt.Println("Type:", t)

	v := reflect.ValueOf(42)
	fmt.Println("Value:", v)
	k := v.Kind()
	fmt.Println("Kind:", k)
}