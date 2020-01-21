package main

import (
	"testing"
)

func TestMinx(t *testing.T) {
	Minx(-128, 255)
	//fmt.Printf("%f", math.Min(math.Pow(float64(-2), float64(31)), math.Pow(float64(2), float64(31))-float64(1)))
}
