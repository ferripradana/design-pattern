package decorator

import "decorator-pattern/component"

type TomatoTopping struct {
	Pizza component.IPizza
}

func (t *TomatoTopping) GetPrice() float32 {
	pizzaPrice := t.Pizza.GetPrice()
	return pizzaPrice + 3
}
