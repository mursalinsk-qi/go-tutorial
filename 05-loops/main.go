package main

import "fmt"

func main() {
	var number int
	// Sum of n numbers
	fmt.Println("Enter number ")
	fmt.Scan(&number)
	sum:=0
	for i:=1;i<=number;i++ {
		sum+=i
	}
	fmt.Println("sum of",number,"numbers : ",sum)
}