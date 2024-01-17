package main

import (
	"builder-pattern/builder"
	"builder-pattern/building"
)

type Director struct {
	builder builder.Ibuilder
}

func NewDirector(b builder.Ibuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) SetBuilder(b builder.Ibuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() building.House {
	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
