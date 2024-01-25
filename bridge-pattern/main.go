package main

import (
	"bridge-pattern/abstraction"
	"bridge-pattern/implementation"
	"fmt"
)

func main() {
	fmt.Print("Run ...")
	hp := &implementation.Hp{}
	epson := &implementation.Epson{}

	fmt.Println("===== Mac =====")
	mac := &abstraction.Mac{}
	mac.SetPrinter(hp)
	mac.Print()
	fmt.Println()
	mac.SetPrinter(epson)
	mac.Print()

	fmt.Println("===== Windows =====")
	windows := &abstraction.Windows{}
	windows.SetPrinter(hp)
	windows.Print()
	fmt.Println()
	windows.SetPrinter(epson)
	windows.Print()

}
