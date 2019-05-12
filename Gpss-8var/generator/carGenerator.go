package generator

import (
	"math/rand"
	"../model"
	"time"
)


func CarGenerator(out chan<- model.Car) model.Car {
	ticker := time.NewTicker(833 * time.Microsecond)

	for{
		select {
		case <-ticker.C:
			out <- model.NewCar(int(rand.ExpFloat64()) % 3 + 4,int(rand.ExpFloat64()) % 5,int(rand.ExpFloat64()) % 5)
		}
	}
}

func timeGenerator() int {
	return int(rand.ExpFloat64()) % 3 + 4
}

func exponentialRandomNumber() int {
	return 0
}

func u01() float64{
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}