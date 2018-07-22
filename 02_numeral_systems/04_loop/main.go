package main

import "fmt"

func main()  {
	for i := 65; i < 90; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i , i , i, i)
	}
}