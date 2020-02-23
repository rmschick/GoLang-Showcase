//Ryan Schick
//variadic_functions.go showcases that functions in go can have a dynamic number of arguments in a function
package main

import "fmt"

func sums(nums ...int){		//can have 0-infinite arguments sent
	fmt.Print(nums, " ")
	total :=0
	for _, num:= range nums{
		total += num
	}
	fmt.Println(total)
}
func main(){
	sums(1,2,3,4,5,6)
	sums(2*5,9,10)
	sums()
	nums:= []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,1,18,19}
	sums(nums...)	//can send an array full of numbers as an argument
}