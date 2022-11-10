package main

import (
	"fmt"
)

type MyInteger int

func main() {

	myIntegerLocal := MyInteger(100)
	fmt.Println("my integer is", myIntegerLocal.getInteger())
}

func (myInteger MyInteger) getInteger() int {
	return int(myInteger)
}
