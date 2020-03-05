//Ryan Schick
/*a closure is a function that has an environment of its own. In this environment,
there is at least one bound variable (a name that has a value, such as a number).
The closure's environment keeps the bound variables in memory between uses of the closure. */
package main

import "fmt"

func intSeq() func() int {
	i := 1
	return func() int {
		i *= 2
		return i
	}
}
func main() {
	nextInt := intSeq()

	fmt.Println("nextInt() starting value is ", nextInt())

	newInts := intSeq()
	fmt.Println("newInt() starting value is ", newInts())

	for j := 0; j < 19; j++ {
		if j%2 == 0 {
			nextInt()
		} else {
			newInts()
		}
	}

	fmt.Println("\nnextInt() ending value is ", nextInt())
	fmt.Println("newInts() ending value is ", newInts())
}
