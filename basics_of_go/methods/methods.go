/*Learning from gobyexample.com
Ryan Schick

Go supports methods defined on struct types
*/

package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { //area has a "reciever type" of *rect
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Rectangle r = ", r)
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r

	fmt.Println("Rectangle rp = ", rp)
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	test := rect{height: 100, width: 100}

	fmt.Println("Rectangle test = ", r)
	fmt.Println("area: ", test.area())
	fmt.Println("perim: ", test.perim())
}
