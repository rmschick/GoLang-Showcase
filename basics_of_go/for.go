/*
Author: Ryan Schick
Date: 15 January 2020

This programs uses 4 types of for loops to showcase purposes of each in different ways
*/
package main

import "fmt"

func main() {
	var i int = 1
	num := 10

	//for loop thats only one condition
	for i <= 3 {
		fmt.Print(i, " ")
		i = i + 1
	}
	fmt.Println()

	//traditional for loop with three parameters
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//for loop goes until it hits a break
	for {
		fmt.Println("This will break")
		fmt.Println("78.32/3.2 =", 78.32/3.2)
		break
	}

	//can set conditions like continue to "skip"
	for n := 10; n >= 5; n-- {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//more practice
	for k := 0; k<num; k++ {
		fmt.Println(k, "*", num, " = ", k*num)
	}

	for a := i; a<=num; a++ {
		fmt.Println("I want ", a, "hamburgers!")
	}
}
