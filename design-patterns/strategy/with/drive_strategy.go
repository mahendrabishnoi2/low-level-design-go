package main

import (
	"fmt"
)

type driveStrategy interface {
	drive()
}

type normalDrive struct {
}

func (nd normalDrive) drive() {
	// normal drive capability
	fmt.Println("normal drive capability")
}

type sportsDrive struct {
}

func (sd sportsDrive) drive() {
	// sports drive capability
	fmt.Println("sports drive capability")
}
