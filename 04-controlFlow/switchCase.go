package main

import (
	"fmt"
)
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
	var number1, number2 int
	var operator string
	fmt.Println("Enter two number :")
	fmt.Scanln(&number1,&number2)
	fmt.Println("enter operator (+,-,*,/")
	fmt.Scan(&operator)
	result:=calculator(number1,number2,operator)
	fmt.Println(result)


}