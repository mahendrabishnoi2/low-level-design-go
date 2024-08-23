package main

type iVehicle interface {
	drive()
}

func main() {
	vehicles := []iVehicle{
		newGoodsVehicle(),
		newPassengerVehicle(),
		newSportsVehicle(),
		newOffroadVehicle(),
	}

	for _, v := range vehicles {
		v.drive()
	}
}

type vehicle struct {
	driveStrategy driveStrategy
}

func newVehicle(driveStrategy driveStrategy) vehicle {
	return vehicle{
		driveStrategy: driveStrategy,
	}
}

func (v vehicle) drive() {
	v.driveStrategy.drive()
}

type goodsVehicle struct {
	vehicle
}

func newGoodsVehicle() goodsVehicle {
	return goodsVehicle{
		vehicle: newVehicle(normalDrive{}),
	}
}

type passengerVehicle struct {
	vehicle
}

func newPassengerVehicle() passengerVehicle {
	return passengerVehicle{
		vehicle: newVehicle(normalDrive{}),
	}
}

type sportsVehicle struct {
	vehicle
}

func newSportsVehicle() sportsVehicle {
	return sportsVehicle{
		vehicle: newVehicle(sportsDrive{}),
	}
}

type offroadVehicle struct {
	vehicle
}

func newOffroadVehicle() offroadVehicle {
	return offroadVehicle{
		vehicle: newVehicle(sportsDrive{}),
	}
}
