package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	n := 43210
	// 运算符的运算顺序 ->
	fmt.Println(n/60*60, "hours and ", n%60*60, " seconds")

	//err := Foo()
	var err *os.PathError = nil

	var e interface{} = nil
	fmt.Println(err, " ", reflect.TypeOf(err), " ", err)
	fmt.Println(nil, " ", reflect.TypeOf(nil), " ", nil)
	fmt.Println(err == e)
}

func Foo() error {
	var err *os.PathError = nil

	return err
}
