package main

import "fmt"

func main()  {
	for i := 65; i <= 90; i++ {
		//fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

	foo := "A" // Best option is using '' - rune literals
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	//fmt.Printf("%T \n", rune(foo))
}
