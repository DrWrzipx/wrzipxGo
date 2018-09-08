package main

import "fmt"

func main()  {
	fmt.Println(greet("Matic ", "Drobez"))
	fmt.Println(returnNaming("Matic ", "Drobez"))
	fmt.Println(returnMultiple("matic", "drobez"))
}

func greet(fName, lName string) string {
	return fmt.Sprint(fName, lName)
}


func returnNaming(fName string, lName string) (s string) {
	s = fmt.Sprint(fName, lName)
	return
}

func returnMultiple(fName, lName string) (string, string) {
	return fmt.Sprint(fName, lName), fmt.Sprint(lName, fName)
}
