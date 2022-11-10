package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() {
	v.X = 8
	fmt.Println("The updated val is:", v.X)
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	v.Abs()
	fmt.Println(v.X)
}
