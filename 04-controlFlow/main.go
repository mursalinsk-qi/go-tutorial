package main

import "fmt"

func isEvenOdd(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
func main() {
	var number int
	fmt.Scan(&number)
	result :=isEvenOdd(number)
	if result {
		fmt.Println(number ,"is even")
	} else {
		fmt.Println(number,"is odd")
	}
}