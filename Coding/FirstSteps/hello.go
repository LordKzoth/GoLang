// Объявление нового пакета
package main

// Импорт пакета
import "fmt"

func main() { // Entry point
	var (
		name, surname string
	)

	fmt.Println("Hi, who are you?")
	fmt.Scan(&name, &surname)

	fmt.Println("Hello, " + name + " " + surname + "!")
}
