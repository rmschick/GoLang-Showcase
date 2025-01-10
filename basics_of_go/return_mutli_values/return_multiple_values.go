//Ryan Schick
//return_multiple_values.go showcases that functions in go are able to return multiple values
package main

import "fmt"

func vars() (int, int){
	return 3,7
}
func main(){
	a,b := vars()		//will assign a and b to be 3 and 7 respectively
	fmt.Println(a)
	fmt.Println(b)

	_, c := vars()	// use the blank identifier _. to get a subset of the returned values
	fmt.Println(c)
}

