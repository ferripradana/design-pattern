package main

import (
	"factory-pattern/gun"
	"fmt"
)

func getGun(gunType string) (gun.IGun, error) {
	if gunType == "ak-47" {
		return gun.NewAk47(), nil
	} else if gunType == "musket" {
		return gun.NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
