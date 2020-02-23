package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const num = 500000000000

	const d = 3e20 / num
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(num))
	const z float64 = 100
	const x float64 = 20.0
	fmt.Println(math.Pow(z, x))

}
