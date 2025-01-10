//Ryan Schick
//basics on how to implement functions in go
//help with gobyexample.com
package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a,b,c int) int {
	return a + b + c
}

func minus(a,b int) int {
	return a - b
}

func mult(a float32,b float32 ) float32{
	return a * b
} 

func divide(a,b float32) float32{
	return a/b
}

func main(){
	res := plus (1,2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 =",res)

	res = minus(70,1)
	fmt.Println("70-1 =", res)

	fRes := mult(50.98, 12.0)
	fmt.Println("50.98*12 =", fRes)

	fRes = divide(12.0,2.0)
	fmt.Println("12/2 = ", fRes)
}