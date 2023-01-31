package main

import "fmt"

func main() {
	var studentName string = "Mursalin Sk"
	var roolNumber int = 23
	var totalMarks float32 = 426

	percentageMarks := totalMarks / 5
	fmt.Println(studentName,"roll no - ",roolNumber,"got",percentageMarks,"%")
}