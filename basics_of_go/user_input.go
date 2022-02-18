/*
Author: Ryan Schick
Date: 02/17/2022

Taking in user input example. Utilized https://www.geeksforgeeks.org/how-to-take-input-from-the-user-in-golang/


*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your first name: ")

	var first string
	fmt.Scanln(&first)
	
	fmt.Println("Enter your last name: ")
	
	var last string
	fmt.Scanln(&last)

	fmt.Println("Your full name is", first, last)
}

