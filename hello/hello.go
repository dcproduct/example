package main

import (
		"fmt"
		"time"
		"math/rand"
		"math"
		"math/cmplx"
		"github.com/dcproduct/example/stringutil"
	)

var c, python, java = true, false, "no!"

var (
		ToBe bool = false
		MaxInt uint64 = 1<<64 - 1
		z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main()  {
	fmt.Println(stringutil.Reverse("!oG, olleH"))
	fmt.Println("My favourite number: " , rand.Intn(10))
	fmt.Println("The time is", time.Now())
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	fmt.Printf("\n")
	fmt.Println(math.Pi)

	fmt.Println(add(33,22))

	fmt.Println(swap("hello", "world"))
	fmt.Println(split(33))


	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)


}

func add(x, y int) int  {
	return x + y
}

func swap(x, y string) (string, string)  {
	return y, x
}

func split(sum int) (x, y int)  {
	x = sum * 4/9
	y = sum - x
	return
}
