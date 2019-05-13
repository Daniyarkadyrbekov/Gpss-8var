package generator

import (
	"../model"
	"math"
	"math/rand"
	"time"
)

const expCarGenerator = 1250

func CarGenerator(addCar chan<- model.Car, tick, timer chan<- struct{}) model.Car {
	//ticker := time.NewTicker(expCarGenerator * time.Microsecond)
	//tickerSec := time.NewTicker(time.Millisecond)
	//timeOut := time.NewTimer(3600 * time.Millisecond)
	//
	//for {
	//	select {
	//	case <-ticker.C:
	//		addCar <- model.NewCar(timeGenerator(), roadGenerator(), roadGenerator())
	//	case <-tickerSec.C:
	//		tick <- struct{}{}
	//	case <-timeOut.C:
	//		timer <- struct{}{}
	//	}
	//}
	var addTime = .0
	var addTimes []float64

	for addTime <= 3600 {
		addTime += exponensialCarGenerator()
		addTimes = append(addTimes, addTime)
	}
	//logrus.WithField("addtimes", addTimes[:10]).Info("addTimes Generated")

	var timeNow = 0
	for {
		if timeNow > 3600 {
			timer <- struct{}{}
		}
		if len(addTimes) > 0{
			for int(addTimes[0]) <= timeNow {
				addCar<-model.NewCar(timeGenerator(), roadGenerator(), roadGenerator())
				if len(addTimes) > 2 {
					addTimes = addTimes[1:]
				}else{
					addTimes = []float64{}
					break
				}
			}
		}
		tick<- struct{}{}
		timeNow++
	}
}

func exponensialCarGenerator() float64 {
	var gen float64
	gen = -math.Log(u01()) * 1.25
	return gen
}

func timeGenerator() int {
	var time float64
	time = -math.Log(u01()) * 30
	if time < 20 {
		time = 20
	}
	return int(time) / 5
}

func roadGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % 5
}

func u01() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
