package main

import (
	"fmt"
)

type vehicle struct {
}

func (v vehicle) drive() {
	fmt.Println("normal drive capability")
}

type goodsVehicle struct {
	vehicle
}

type passengerVehicle struct {
	vehicle
}

type sportsVehicle struct {
}

// sportsVehicle has a different drive capability
func (sv sportsVehicle) drive() {
	// overwriting the drive method of vehicle
	fmt.Println("sports drive capability")
}

type offroadVehicle struct {
}

// offroadVehicle has a different drive capability
func (ov offroadVehicle) drive() {
	// same as sportsVehicle, duplicated code
	fmt.Println("sports drive capability")
}

type iVehicle interface {
	drive()
}

func main() {
	var vehicles []iVehicle

	vehicles = append(vehicles, vehicle{})
	vehicles = append(vehicles, goodsVehicle{})
	vehicles = append(vehicles, passengerVehicle{})
	vehicles = append(vehicles, sportsVehicle{})
	vehicles = append(vehicles, offroadVehicle{})

	for _, v := range vehicles {
		v.drive()
	}
}
