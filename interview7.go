package main

import "fmt"

func Minx(a int8, b uint8) {
	var min uint8 = 0

	//min = uint8(copy(make([]struct{}, ), make([]struct{}, b)))

	//min = uint8(math.Min(float64(uint8(a)), float64(b)))
	fmt.Printf("the min of %d and %d is %d\n", a, b, min)

}
