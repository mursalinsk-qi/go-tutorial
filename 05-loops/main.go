package main

import "fmt"

func isPrime(number int) bool{
	for i:=2;i<number;i++ {
		if number%i==0{
			return false
		}
	}
	return true
}
func printPrimes(a,b int){
	for num:=a;num<=b;num++ {
		if isPrime(num) {
			fmt.Print(num," ")
		}
	}
}
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

	// Prime number between a and b
	
	printPrimes(10,100)
}