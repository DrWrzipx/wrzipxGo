package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This will run !")
	}

	if !false {
		fmt.Println("This will not run !")
	}


	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	if false {
		fmt.Println("This will not run!")
	} else
	{
		fmt.Println("This will print this statement.")
	}


	for i := 0; i <= 100; i++ {
		if i % 2 != 0 {
			fmt.Println(i)
		}
	}
}
