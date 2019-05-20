package main

import (
	"./generator"
	"./model"
	"github.com/sirupsen/logrus"
)

var Cars = []model.Car{}

func main() {
	//carGeneratorChan := make(chan model.Car)
	tickChan := make(chan model.DeltaType)
	timeoutChan := make(chan struct{})
	go generator.CarGenerator(tickChan, timeoutChan)

	circle := model.NewCircle()
LOOP:
	for {
		select {
		case delta := <-tickChan:
			circle.Next(delta)
		case <-timeoutChan:
			//fmt.Printf("end ticker\n")
			break LOOP
		//case car := <-carGeneratorChan:
		//	circle.Add(car)
		}
	}
	logrus.WithFields(logrus.Fields{
		"Пропускная способность": float64(model.CarTerminated) / 3600.,
	}).Info("Вывод")
	logrus.WithFields(logrus.Fields{
		"Средняя длин очереди": float64(model.AvgQueue) / 3600.,
	}).Info("Вывод")
}
