/*Learning from gobyexample.com
Ryan Schick
slices.go shows implementation of slices

Unlike arrays, slices are typed only by the elements they contain (not the number of elements).

Slices are more functional than arrays. You can append and copy with ease. These are used more than arrays for production
*/
package main

import "fmt"


func main(){
	s:= make([]string, 3)
	fmt.Println("emp: ",s)
	b := make([]string, 5)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	b = append(b, "I want tacos")
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("set: ", b)
	fmt.Println("get: ", b[5])

	fmt.Println("len: ", len(s)) //prints out length of slice s

	var word string = "h i j k elemeno p"
	s = append(s, "d")
	s= append(s, "e", "f", "g", word)
	fmt.Println("apd: ", s)

	c:= make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy: ", c)

	l:= s[2:5]
	fmt.Println("sl1: ",l)

	//This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2: ", l)

	//And this slices up from (and including) s[2].
	l = s[6:]
	fmt.Println("sl3: ", l)

	t := []string{"g","h","i"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 5)
	//Slices can be composed into multi-dimensional data structures. 
	//The length of the inner slices can vary, unlike with multi-dimensional arrays.
	for i := 0; i<5; i++{
		innerLen := i +1
		twoD[i] = make([]int, innerLen)
		for j:= 0; j<innerLen; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}