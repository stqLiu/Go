package main

import (
	"fmt"
)

func threadA(chanA chan string, str string) {
	for ;; {
		fmt.Println("A S");
		chanA <- "chanA exit"
		fmt.Println("A");
	}
}

func threadB(chanB chan string, str string) {
	for ;; {
		fmt.Println("B S");
		chanB <- "chanB exit"
		fmt.Println("B");
	}
}

func threadC(chanC chan string, str string) {
	for ;; {
		fmt.Println("C S");
		chanC <- "chanC exit"
		fmt.Println("C");
	}
}


func main() {
	var chanA = make(chan string)
	var chanB = make(chan string)
	var chanC = make(chan string)
	//var chanD = make(chan string)
	go threadA(chanA, "threadA")
	go threadB(chanB, "threadB")
	go threadC(chanC, "threadC")

	fmt.Println(<- chanA)
	fmt.Println(<- chanB)
	fmt.Println(<- chanC)

	//<- chanD
	for true {
		fmt.Println(<- chanA)
		fmt.Println(<- chanB)
		fmt.Println(<- chanC)
	}
}
