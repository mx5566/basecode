package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 1
	dl := data{"one"}
	dl.setName("three") // 指针方法 name被修改了 three
	dl.setName1("four") // 值方法 name没有修改 只是修改了nama副本的值
	dl.print()

	// 2
	var dl1 = &data{"one"}
	dl1.setName("three") // 指针方法 被修改成功
	dl1.setName1("four") // 值方法 name没有修改 只是修改了nama副本的值
	dl1.print()

	// 3
	// data继承了printer实现的接口方法通过指针方法实现 所以printer只能接受指针类型
	var in printer = &data{"two"}
	in.setName("three")
	in.print()
	fmt.Println("type %T", reflect.TypeOf(in))

	// 4
	// 下面两个结果是two
	// 原因是go检测到只实现了结构方法、这样结构方法不能够修改只能修改副本的值
	// 下面虽然赋值的是指针类型、但是go语言会检测到这是个结构体方法 他会隐式的转换为
	// (*in1).setName("three")
	// 还是只副本改变得值
	var in1 printer = &data1{"two"}
	in1.setName("three")
	in1.print()

	fmt.Println("type %T", reflect.TypeOf(in1))
	// 5
	var in2 printer = data1{"two"}
	in2.setName("three")
	in2.print()

	fmt.Println("type %T", reflect.TypeOf(in2))

}

type data struct {
	name string
}

type data1 struct {
	name string
}

func (p *data) print() {
	fmt.Println("name: ", p.name)
}
func (p *data) setName(name string) {
	p.name = name
}

func (p data) setName1(name string) {
	p.name = name
}

func (p data1) print() {
	fmt.Println("name: ", p.name)
}

func (p data1) setName(name string) {
	p.name = name
}

type printer interface {
	print()
	setName(name string)
}
