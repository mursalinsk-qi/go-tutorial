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
	// Prime number between a and b
	var a,b int
	fmt.Println("enter two number")
	fmt.Scan(&a,&b)
	printPrimes(a,b)
}