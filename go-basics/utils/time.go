package utils
import "time"
import "fmt"
func TimeExample(){
 	now := time.Now()
 	fmt.Println("Current time:", now)

 	// Formatear la fecha y hora
 	formattedTime := now.Format("2006-01-02 15:04:05")
 	fmt.Println("Formatted time:", formattedTime)

 	// Obtener partes específicas de la fecha y hora
 	year := now.Year()
 	month := now.Month()
 	day := now.Day()
 	hour := now.Hour()
 	minute := now.Minute()
 	second := now.Second()
	fmt.Printf("Year: %d, Month: %s, Day: %d, Hour: %d, Minute: %d, Second: %d\n",
	 year, month, day, hour, minute, second)
	 then := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	 fmt.Println("Specific date:", then) 
	 then =  then.AddDate(0, 0, 10) // Agregar 10 días
	 fmt.Println("New date after adding 10 days:", then)
	 fmt.Println("Difference between now and then:", then.Before(now))
	 
	 fmt.Println("Difference between now and then:", then.After(now))
	 fmt.Println("Difference between now and then:", then.Equal(now))
	 fmt.Println("Difference between now and then:", then.Sub(now))
}