/*Learning from gobyexample.com
Ryan Schick

structs are typed collections of fields
28 Sept 2020
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}

func main() {
	fmt.Println("Enter your first name: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}

	fmt.Println("Enter your age: ")
	var age int
	_, err = fmt.Scanln(&age)
	if err != nil {
		return
	}

	fmt.Println(*newPerson(name, age))
}
