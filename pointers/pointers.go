package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Scale() { //manipulates the copy of vertex which is passed
	v.X = v.X + 1
	v.Y = v.Y + 1
	fmt.Println("Inside scale the values are :", v.X, v.Y)
}

func (v *Vertex) Abs() { //manipulates the original vertex which is passed
	v.X = 8
	v.Y = 9
	fmt.Println("Inside abs updated val is:", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale()
	fmt.Println("After scale the values are: ", v.X, v.Y)

	vertex := Vertex{7, 8}
	p := &vertex //gives memory location
	p.Abs()
	fmt.Println("After abs the values are :", vertex.X, vertex.Y)
}
