package main

import "fmt"

func main()  {
	n := average(45, 334, 23, 94, 34)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}