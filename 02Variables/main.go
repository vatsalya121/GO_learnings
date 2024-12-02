package main

import "fmt"

// jwttoken := 5646   // here this novar methpd is nolt allowed as this isnt inside any declared method

func main() {
	var username string = "Vatsalya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var islogged bool = true
	fmt.Println(islogged)
	fmt.Printf("Variable is of type: %T \n", islogged)

	var smlVal int = 256
	fmt.Println(smlVal)
	fmt.Printf("Variable is of type: %T \n", smlVal)

	//implicit type
	var website = 12345
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var method
	systemm := "munna" //only allowed inside any declared method
	fmt.Println(systemm)
	fmt.Printf("Variable is of type: %T \n", systemm)

}
