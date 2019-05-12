package model

import "github.com/sirupsen/logrus"

type Car struct {
	Time int
	EnterRoad int
	OutRoad int
	iterations int
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
	if c.iterations % c.Time == 0 {
		logrus.Info("next true")
		return true
	}
	return false
}

func (c *Car) Terminate() bool {
	c.iterations = c.iterations - 1

	if c.iterations == 0 {
		logrus.WithField("car", *c).Info("car teminated")
		return true
	}
	return false
}
