package model

type DeltaType struct {
	DeltaCar Car
	DeltaTime float64
}

var CarTerminated = 0

type Car struct {
	Time int
	EnterRoad int
	OutRoad int
	Iterations int
}

func NewCar(time, enterRoad, outRoad int) Car {
	var roadsNum int
	if outRoad > enterRoad {
		roadsNum = outRoad - enterRoad
	}else{
		roadsNum = outRoad + roadsNumber - enterRoad
	}

	return Car{time, enterRoad,outRoad, roadsNum * time}
}

func(c *Car) Next() bool {
	if c.Iterations % c.Time == 0 {
		//logrus.Info("next true")
		return true
	}
	return false
}

func (c *Car) Terminate() bool {
	c.Iterations = c.Iterations - 1

	if c.Iterations == 0 {
		//logrus.WithField("car", *c).Info("car teminated")
		CarTerminated++
		return true
	}
	return false
}
