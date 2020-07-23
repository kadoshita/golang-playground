package main

import "fmt"

func main() {
	var n int = 10
	n1 := 0.01
	const str = "golang"
	var x, y bool = true, false
	var v1, v2, v3 = 1, -0.5, "str"

	var (
		Var1         = false
		Var2 int     = 10
		Var3 float32 = 0.002
	)

	fmt.Println(n)
	fmt.Println(n1)
	fmt.Println(str)
	fmt.Println(x, y)
	fmt.Printf("Type of v1: %T v2: %T v3: %T\n", v1, v2, v3)
	fmt.Println(v1, v2, v3)
	fmt.Printf("Type: %T Value: %v\n", Var1, Var1)
	fmt.Printf("Type: %T Value: %v\n", Var2, Var2)
	fmt.Printf("Type: %T Value: %v\n", Var3, Var3)
}
