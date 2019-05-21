package generator

import (
	"../model"
	"math"
	"math/rand"
	"time"
)

const expCarGenerator = 1250

var CarGenerated = 0

func CarGenerator( tick chan<- model.DeltaType, timer chan<- struct{}) {

	var timeNow = .0
	for {
		deltaTime := exponensialCarGenerator()
		timeNow += deltaTime
		if timeNow > 360 {
			timer <- struct{}{}
			break
		}

		//logrus.WithField("deltaTime", deltaTime).Info("deltaTimeGenerated Generated")
		CarGenerated++
		tick <- model.DeltaType{model.NewCar(timeGenerator(), roadGenerator(), roadGenerator()), deltaTime}
	}
}

func exponensialCarGenerator() float64 {
	var gen float64
	gen = -math.Log(u01()) * 1.3
	return gen
}

func timeGenerator() float64 {
	var time float64
	time = -math.Log(u01()) * 30
	if time < 20 {
		time = 20
	}
	return time / 5
}

func roadGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % 5
}

func u01() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
