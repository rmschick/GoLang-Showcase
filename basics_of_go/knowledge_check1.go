package main

import (
	"errors"
	"fmt"
	"math"
)

type COD struct{
	gamerTag string
	KDR float32
	pogChamp bool
	favGun string
}

func getNumberFromUser(number *int){

	fmt.Println("Enter a number between 1-100: ")
	fmt.Scanln(number)

}

func newCODplayer(gamerTag string, KDR float32, favGun string) *COD{
	c := COD{gamerTag: gamerTag, KDR: KDR, pogChamp: false, favGun: favGun}

	if KDR > 2 {
		c.pogChamp = true
	}
	return &c
}

//variadic functions
func total(nums ...int){
	fmt.Print(nums, " ")
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func main(){
	fmt.Println("\nHello. This is the first knowledge check.")
	fmt.Println("We will use functions, math, user input, slices, for loops, etc... The basics.")

	fmt.Println("\nEnter your first name ")
	
	var firstName string
	fmt.Scanln(&firstName)

	var number int
	//function demonstration
	getNumberFromUser(&number) 
	for number < 1 || number > 100 {
		fmt.Println("ERROR: ", number, " is not in range.")
		getNumberFromUser(&number)
	}
	//for loop
	for i := 1; i<10; i++{
		fmt.Println("Your number", number, "*", i, "is: ", number * i)
	}
	//slice
	s := make([]int, 5)
	for j:= 0; j<5; j++{
		s[j] = j
	}
	fmt.Println("\n", s)

	s = append(s, 5, 6, 7, 8, 9, 10)
	fmt.Println("app:", s)

	sum := 0
	for _, num := range s{
		sum += num
	}

	fmt.Println("sum:", sum)
	//knowledge in structs
	fmt.Println("\n",*newCODplayer("Samurai RyGuy", 6.7, "Cooper"))

	///square root function with error handling

	result, err := sqrt(16)

	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("For the sqrt of 16 we get back:", result)
	}

	result, err = sqrt(-10)

	if err != nil{
		fmt.Println("For the sqrt of -10, we get back:", err)
	} else{
		fmt.Println(result)
	}
}