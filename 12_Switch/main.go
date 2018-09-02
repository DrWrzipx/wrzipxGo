package main

import "fmt"

func main()  {
	name := " "
	fmt.Println("Please insert your name")
	fmt.Scanln(name)
	switch name {
	case "Daniel":
		fmt.Println("What's up Daniel!")
	case "Medhi":
		fmt.Println("What's up Medhi!")
	case "Matic":
		fmt.Println("What's up Matic")
	default:
		fmt.Println("What's up man!")
	}
}
