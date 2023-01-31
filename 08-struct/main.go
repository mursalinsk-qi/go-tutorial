package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}
func calculateArea(rect Rectangle) int{
	area:=rect.length*rect.breadth
	return area
}
func main() {
	rect := Rectangle{15, 28}

	// access the length of the struct
	fmt.Println("Length:", rect.length)

	// access the breadth of the struct
	fmt.Println("Breadth:", rect.breadth)

	area := calculateArea(rect)
	fmt.Println("Area:", area)

}
