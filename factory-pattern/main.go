package main

import (
	"factory-pattern/gun"
	"fmt"
)

func main() {
	fmt.Println("Run....")

	ak47, _ := getGun("ak-47")
	musket, _ := getGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g gun.IGun) {
	fmt.Println("========GUN=======")
	fmt.Println("Name : ", g.GetName())
	fmt.Println("Power : ", g.GetPower())
	fmt.Println()
}
