package main

import (
	"decorator-pattern/component"
	"decorator-pattern/decorator"
	"fmt"
)

func main() {
	fmt.Println("Run...")

	fmt.Println("=============================")

	meatPizza := &component.Meat{}
	meatPizzaWithCheese := &decorator.CheeseTopping{Pizza: meatPizza}
	meatPizzaWithCheeseAndTomato := &decorator.TomatoTopping{Pizza: meatPizzaWithCheese}

	fmt.Printf("Price of meat pizza with tomato and cheese topping is %f\n", meatPizzaWithCheeseAndTomato.GetPrice())

	fmt.Println("=============================")

	veggiePizza := &component.Veggie{}
	veggiePizzaWithCheese := &decorator.CheeseTopping{Pizza: veggiePizza}
	veggiePizzaWithCheeseAndTomato := &decorator.TomatoTopping{Pizza: veggiePizzaWithCheese}
	fmt.Printf("Price of veggie pizza with tomato and cheese topping is %f\n", veggiePizzaWithCheeseAndTomato.GetPrice())

}
