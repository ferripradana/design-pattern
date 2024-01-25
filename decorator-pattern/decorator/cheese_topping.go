package decorator

import "decorator-pattern/component"

type CheeseTopping struct {
	Pizza component.IPizza
}

func (c *CheeseTopping) GetPrice() float32 {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
