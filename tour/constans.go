package tour

import "fmt"

const Pi = 3.14

func Cnst()  {
	const World = "世界"
	fmt.Println("Hello",World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
