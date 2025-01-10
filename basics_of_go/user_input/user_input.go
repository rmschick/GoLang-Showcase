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
	_, err := fmt.Scanln(&first)
	if err != nil {
		return
	}

	fmt.Println("Enter your last name: ")

	var last string
	_, err = fmt.Scanln(&last)
	if err != nil {
		return
	}

	fmt.Println("Your full name is", first, last)
}
