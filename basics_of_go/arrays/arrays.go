package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 130

	fmt.Println("set: ", a)      //printing the whole array after seting last spot to 130
	fmt.Println("get: ", a[4])   //printing the last spot in the array
	fmt.Println("len: ", len(a)) //length of array

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [10][5]int
	for i := 1; i < 11; i++ {
		for j := 1; j < 6; j++ {
			twoD[i-1][j-1] = i * j
		}
	}
	fmt.Println("2d: ", twoD)
}
