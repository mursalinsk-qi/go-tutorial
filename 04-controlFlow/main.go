package main

import "fmt"

func isEvenOdd(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
func calculator(number1,number2 int,operator string) int{
	switch operator {
	case "+":
		return number1+number2
	case "-":
		return number1-number2
	case "*":
		return number1*number2
	case "/":
		return number1/number2
	default:
		fmt.Println("invalid operation")
	}
	return -1
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