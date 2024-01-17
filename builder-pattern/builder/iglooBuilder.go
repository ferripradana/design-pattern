package builder

import "builder-pattern/building"

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() Ibuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) SetWindowType() {
	i.windowType = "Snow Window"
}

func (i *IglooBuilder) SetDoorType() {
	i.doorType = "Snow Door"
}

func (i *IglooBuilder) SetNumFloor() {
	i.floor = 1
}

func (i *IglooBuilder) GetHouse() building.House {
	return building.House{
		WindowType: i.windowType,
		DoorType:   i.doorType,
		Floor:      i.floor,
	}
}
