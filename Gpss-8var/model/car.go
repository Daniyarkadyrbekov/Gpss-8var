package model

type DeltaType struct {
	DeltaCar Car
	DeltaTime float64
}

type Car struct {
	in int
	out int
	timePerRoad float64
	prioritet bool
}

func NewCar(timeForPart float64, enterRoad, outRoad int) Car {
	return Car{enterRoad, outRoad, timeForPart, false}
}
