package main

import "fmt"

func main()  {

	a := 45

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	c := *b * 28
	fmt.Println(c)
}
