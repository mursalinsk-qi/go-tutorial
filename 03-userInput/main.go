package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// string input
	reader:=bufio.NewReader(os.Stdin)
	name,_:=reader.ReadString('\n')
	name=strings.TrimSpace(name)
	fmt.Println("hi",name,"good morning")

	// int input
	fmt.Println("enter two number")
	var number1,number2 int
	fmt.Scan(&number1,&number2)
	sum:=number1+number2
	print("sum = ",sum)
	
}
