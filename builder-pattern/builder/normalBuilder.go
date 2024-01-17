package builder

import "builder-pattern/building"

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() Ibuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) SetWindowType() {
	n.windowType = "Wooden Window"
}

func (n *NormalBuilder) SetDoorType() {
	n.doorType = "Wooden Door"
}

func (n *NormalBuilder) SetNumFloor() {
	n.floor = 2
}

func (n *NormalBuilder) GetHouse() building.House {
	return building.House{
		WindowType: n.windowType,
		DoorType:   n.doorType,
		Floor:      n.floor,
	}
}
