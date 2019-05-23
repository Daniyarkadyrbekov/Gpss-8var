package generator

import (
	"math"
	"math/rand"
	"time"
)

const expCarGenerator = 1250

var CarGenerated = 0

//func CarGenerator( tick chan<- model.DeltaType) {
//
//	var timeNow = .0
//	for {
//		deltaTime := exponensialCarGenerator()
//		timeNow += deltaTime
//
//		fmt.Println(deltaTime)
//		fmt.Println(timeNow)
//		if timeNow > 36.0 {
//			break
//		}
//		CarGenerated++
//		tick <- model.DeltaType{model.NewCar(timeGenerator(), roadGenerator(), roadGenerator()), timeNow}
//	}
//}

func ExponensialCarGenerator() float64 {
	var gen float64
	gen = -math.Log(u01()) * 1.3
	return gen
}

func TimeGenerator() float64 {
	var time float64
	time = -math.Log(u01()) * 30
	if time < 20 {
		time = 20
	}
	return time / 5
}

func RoadGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % 5
}

func u01() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
