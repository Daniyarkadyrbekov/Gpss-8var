package main

import (
	"./generator"
	"./model"
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
)

var Cars = []model.Car{}

func main() {
	carGeneratorChan := make(chan model.Car)
	go generator.CarGenerator(carGeneratorChan)

	ticker := time.NewTicker(time.Millisecond)
	timer := time.NewTimer(3600 * time.Millisecond)
	circle := model.NewCircle()
	//logCtx := log.WithFields(log.Fields{
	//	"RoadsQueue": circle.Roads,
	//	"circle": circle.Circle,
	//})
LOOP:
	for {
		select {
		case <-ticker.C:
			circle.Next()
			log.WithFields(log.Fields{
				//"RoadsQueue": circle.Roads,
				"circle": circle.Circle,
			}).Info("next1")
			log.WithFields(log.Fields{
				"RoadsQueue": circle.Roads,
				//"circle": circle.Circle,
			}).Info("next2")
		case <-timer.C:
			ticker.Stop()
			fmt.Printf("end ticker\n")
			break LOOP
		case car := <-carGeneratorChan:
			circle.Add(car)
			//log.WithFields(log.Fields{
			//	"RoadsQueue": circle.Roads,
			//	"circle": circle.Circle,
			//}).Info("Add")
		}
	}
}
