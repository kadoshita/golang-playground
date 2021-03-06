package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	X, Y int
}

type Coordinate struct {
	Lat, Long float64
}

var m = map[string]Coordinate{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

var (
	v1  = Vertex{1, 2}
	v2  = Vertex{X: 1}
	v3  = Vertex{}
	v1p = &Vertex{1, 2}
)

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	vp := &v
	vp.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, v1p, v2, v3)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a1 := names[0:2]
	a2 := names[1:3]
	fmt.Println(a1, a2)

	a2[0] = "XXX"
	fmt.Println(a1, a2)
	fmt.Println(names)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s1)

	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)
	s2 = s2[:0]
	printSlice(s2)
	s2 = s2[:4]
	printSlice(s2)
	s2 = s2[2:]
	printSlice(s2)

	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))

	if s3 == nil {
		fmt.Println("nil!")
	}

	s3 = append(s3, 0)
	fmt.Println(s3, len(s3), cap(s3))
	s3 = append(s3, 1)
	fmt.Println(s3, len(s3), cap(s3))
	s3 = append(s3, 2, 3, 4)
	fmt.Println(s3, len(s3), cap(s3))

	a3 := make([]int, 5)
	printSlice2("a3", a3)
	b3 := make([]int, 0, 5)
	printSlice2("b3", b3)
	c3 := b3[:2]
	printSlice2("c3", c3)
	d3 := c3[2:5]
	printSlice2("d3", d3)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var pow = make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	fmt.Println(m)

	m2 := make(map[string]int)

	m2["Answer"] = 42
	fmt.Println("The value: ", m2["Answer"])
	m2["Answer"] = 48
	fmt.Println("The value: ", m2["Answer"])
	delete(m2, "Answer")
	fmt.Println("The value: ", m2["Answer"])
	val, ok := m2["Answer"]
	fmt.Println("The value: ", val, "Present?", ok)

	wc.Test(WordCount)

	hypot := func(x1, y1 float64) float64 {
		return math.Sqrt(x1*x1 + y1*y1)
	}

	fmt.Println()
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
