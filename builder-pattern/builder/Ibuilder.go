package builder

import "builder-pattern/building"

type Ibuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() building.House
}

func GetBuilder(builderType string) Ibuilder {
	if builderType == "normal" {
		return NewNormalBuilder()
	}
	if builderType == "igloo" {
		return NewIglooBuilder()
	}
	return nil
}
