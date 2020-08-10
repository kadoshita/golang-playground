package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}
type T struct {
	S string
}

type F float64

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (f F) M() {
	fmt.Println(f)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nul>")
		return
	}
	fmt.Println(t.S)
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {
	var a Abser
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
	f := MyFloat(-math.Sqrt2)

	a = f
	a = &v
	a = v
	fmt.Println(a.Abs())

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)

	f2, ok := i.(float64)
	fmt.Println(f2, ok)

	// f2 = i.(float64)
	// fmt.Println(f2)
	// var t *T
	// i = t
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)

	// i = &T{"hello"}
	// describe(i)
	// i.M()

	do(21)
	do("hello")
	do(true)

	p1 := Person{"Arthur Dent", 42}
	p2 := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(p1, p2)

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
