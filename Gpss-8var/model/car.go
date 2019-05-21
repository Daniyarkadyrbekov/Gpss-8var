package model

type DeltaType struct {
	DeltaCar Car
	DeltaTime float64
}

var CarTerminated = 0

type Car struct {
	TimeForPart float64
	AllTime float64
	TimeToGo float64
	TimeFixed float64
	Priority bool
	EnterRoad int
	outRoad int
}

func NewCar(time float64, enterRoad, outRoad int) Car {
	var roadsNum int
	if outRoad > enterRoad {
		roadsNum = outRoad - enterRoad
	}else{
		roadsNum = outRoad + roadsNumber - enterRoad
	}

	return Car{time, float64(roadsNum) * time,.0, .0, false, enterRoad, outRoad}
}

func(c *Car) Move(moveTime float64) {
	c.AllTime -= moveTime
	c.TimeToGo -= moveTime
	if c.TimeToGo < 0 {
		c.TimeToGo = .0
	}
	c.TimeFixed += moveTime
}

//func(c *Car) Next() bool {
//	if c.Iterations % c.TimeForPart == 0 {
//		return true
//	}
//	return false
//}
//
//func (c *Car) Terminate() bool {
//	c.Iterations = c.Iterations - 1
//
//	if c.Iterations == 0 {
//		//logrus.WithField("car", *c).Info("car teminated")
//		CarTerminated++
//		return true
//	}
//	return false
//}
