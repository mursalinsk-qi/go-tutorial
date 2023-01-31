package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func main() {
	rect := Rectangle{15, 28}

	// access the length of the struct
	fmt.Println("Length:", rect.length)

	// access the breadth of the struct
	fmt.Println("Breadth:", rect.breadth)

	area := rect.length * rect.breadth
	fmt.Println("Area:", area)

}
