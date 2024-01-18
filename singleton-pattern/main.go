package main

import "fmt"

func main() {

	fmt.Println("Run ........ ......  .... .. ... ")
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
