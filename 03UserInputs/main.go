package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome bros"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza: ")

	//comma ok syntax|| error ok syntax

	input, _ := reader.ReadString('\n')
	// _, err := reader.ReadString('\n')  //caseswhere we only care about errors only
	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("Type of ratingis %T", input)
}
