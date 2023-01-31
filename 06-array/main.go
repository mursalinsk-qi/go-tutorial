package main

import "fmt"

func main() {
	nums := [5]int{1, 5, 8, 12}
	fruits:=[...]string{"apple","mango","banana","guava"}
	fmt.Println(nums)
	fmt.Println(fruits)
	fruits[2]="pineappple"
	fmt.Println(fruits)
	fmt.Println("length : ",len(fruits))
	for _,value:=range fruits{
		fmt.Print(value,",")
	}
}