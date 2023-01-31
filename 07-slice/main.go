package main

import "fmt"

func main() {
	fruits:=[]string{"apple","mango","banana","guava"}
	fmt.Println(fruits)
	fmt.Printf("type of fruits %T",fruits)
	fmt.Println("length of fruits",len(fruits))
	// Adding New element
	fmt.Println("----after adding------")
	fruits=append(fruits, "pineapple","orange","lemon")
	fmt.Println(fruits)
	fmt.Println("length of fruits",len(fruits))
}