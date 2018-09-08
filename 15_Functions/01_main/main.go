package main

import "fmt"

func main()  {
	greet("Matic")
	greetTwoParameters("Drobez", "Matic")
}

func greet(name string)  {
	fmt.Println(name)
}

func greetTwoParameters(lName string, fName string) {
	fmt.Println("Hello", fName, lName)
}
