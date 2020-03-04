package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// 站位这个包 后面要用不用删除了
var _ = os.Args

//type P *int
//type Q *int

func main() {
	//var p P = new(int)
	//*p += 8
	//var x *int = p
	//var q Q = x
	//(*q)++
	//fmt.Println(*p, *q)
	//
	//true := false
	//fmt.Println(true)
	//
	//var f float64
	//fmt.Println(f)
	//
	//bits := *(*uint64)(unsafe.Pointer(&f))
	//
	//fmt.Println(bits)
	//
	//type ptr unsafe.Pointer
	//bits = *(*uint64)(ptr(&f))
	//fmt.Println(bits)
	//
	//var pp ptr = nil
	//
	//fmt.Println(pp)

	//var ret = watShadowDefer(5)
	//fmt.Println(ret)

	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//go func() {
	//	time.Sleep(time.Millisecond)
	//	wg.Done()
	//	//wg.Add(1)
	//}()
	//
	//wg.Wait()

	//	TestChan()
	//TestChan1()

	//Testkk()

	//flag.Parse()
	//fmt.Printf("ip %s port %d\n", ip, port)

	//	JsonTest()

	//testTT()

	//testf()
	//testfx()

	//link("seek", 1, 2, 3, 4)
	//
	//a := []int{1, 2, 3, 4}
	//
	//b := map[int]int{1: 1, 2: 2}
	//
	//fmt.Println(fmt.Sprint(a))
	//link("seek", strings.Trim(fmt.Sprint(a), "[]"))
	//
	//fmt.Println(fmt.Sprint(b))
	//link("seek", strings.Trim(fmt.Sprint(b), "[]"))

	var sli []int = []int{1}

	ss := sli[1:]

	fmt.Println(ss)
}

func link(p ...interface{}) {
	fmt.Println(p)
}

const Pi float64 = 3.14259265358979323846
const zero = 0.0

const (
	size int64 = 1024
	eof  int32 = -1
)

const (
//err1 = errors.New("haha").Error()
//err2 = errors.New("kkeke").Error()
)

const u, v float64 = 0, 3
const a, b, c = 3, 4, "foo"

func TestChan() {
	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()

	go func() {
		count++
		close(ch)
	}()

	<-ch
	fmt.Println(count)
}

func TestChan1() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()

	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)

	for range c {
		fmt.Printf("go num %d\n", runtime.NumGoroutine())
	}
}

func TestYsncMap() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	fmt.Println(1)
}

func Testkk() {
	var wg sync.WaitGroup

	wg.Add(2)
	var ints = make([]int, 0, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i) // append 是不安全的
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(len(ints))
}

func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		//ret := 10
		defer func() {
			fmt.Println("-----------------------")
			ret = ret + 1
		}()
	}

	var j int
	_ = j // 看成是一个占位符占用这个值

	return
}

var ip string
var port int

func init() {
	flag.StringVar(&ip, "ip", "0.0.0.0", "IP")
	flag.IntVar(&port, "port", 8000, "IP")
}

type People struct {
	Name string `json:"name"`
}

func JsonTest() {
	js := `{"name":"ll"}`

	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err ", err)
		return
	}

	fmt.Printf("People: %v", p)

}

type S struct {
	v int
}

func STest() {

}

type TT struct {
	Val int
}

func (t *TT) Incr(wg *sync.WaitGroup) {
	t.Val++

	fmt.Println("---------", t.Val)
	wg.Done()
}

func (t *TT) Print() {
	time.Sleep(time.Second)
	fmt.Println(t.Val)
}

func testTT() {
	var wg sync.WaitGroup
	wg.Add(10)
	var ts = make([]TT, 10)
	for i := 0; i < 10; i++ {
		ts[i] = TT{i}
	}

	for _, t := range ts {
		go t.Incr(&wg)
	}

	wg.Wait()

	fmt.Println(ts)
	for _, ts := range ts {
		go ts.Print()
	}

	time.Sleep(5 * time.Second)
}

func f(i int) {
	if i--; i == 0 {
		return
	}

	//fmt.Printf("%d-", i)

	f(i)
}

func testf() {
	var val int

	fmt.Println(&val)
	f(10000)
	fmt.Println(&val)

	//	println(&val)

}

func testfx() {
	var val1 int

	a := &val1
	println(a)
	f(10000)

	b := &val1
	println(b)

	println(a == b)
}
