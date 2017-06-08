package tour

import "fmt"

func Fo()  {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func FoExt()  {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}

	fmt.Println(sum)
}

func FoExt2()  {
	sum := 1
	for sum < 1000 {
			sum += sum
	}

	fmt.Println(sum)
}
