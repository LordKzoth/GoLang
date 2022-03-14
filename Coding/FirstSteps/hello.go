// Объявление нового пакета
package main

// Импорт пакета
import "fmt"

func main() { // Entry point
	var (
		num1, num2 int = 10, 12
		someString string
	)

	someString = "Some value"
	fmt.Println(someString)

	// Another variable declaration style
	newVariable := 1234

	fmt.Println(num1 + num2 + newVariable) // Print some text in console
}
