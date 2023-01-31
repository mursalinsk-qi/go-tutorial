package main

import (
	"fmt"
	"time"
)

func displayMessage(message string) {
	for i:=0;i<5;i++{
		fmt.Println(message)
		time.Sleep(5*time.Second)
	}
	
}

func main() {
	go displayMessage("process 1 running")
	displayMessage("process 2 running")
	
}
