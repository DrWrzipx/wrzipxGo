package main

import "fmt"

func main()  {
	var name string
	fmt.Print("What is your name ? ")
	fmt.Scan(&name)
	fmt.Print("Your name is ", name, "\n")


	var firstNumber float32
	var secondNumber float32
	var divideResult float32

	fmt.Print("Please enter first number ")
	fmt.Scan(&firstNumber)
	fmt.Print("Please enter second number ")
	fmt.Scan(&secondNumber)

	if (secondNumber == 0) {
		fmt.Print("Devide with zero is not allowed! ")
	} else {
		divideResult = secondNumber / firstNumber
	}
	fmt.Print("The devidion result is " , divideResult, "\n")
}
