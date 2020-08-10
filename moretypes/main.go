package main

import "fmt"

type Vertex struct {
	X, Y int
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
}
