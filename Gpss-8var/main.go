package main

import (
	"./generator"
	"./model"
	"github.com/sirupsen/logrus"
)

var Cars = []model.Car{}

func main() {
	carGeneratorChan := make(chan model.Car)
	tickChan := make(chan struct{})
	timeoutChan := make(chan struct{})
	go generator.CarGenerator(carGeneratorChan, tickChan,timeoutChan)

	circle := model.NewCircle()
LOOP:
	for {
		select {
		case <-tickChan:
			circle.Next()
		case <-timeoutChan:
			//fmt.Printf("end ticker\n")
			break LOOP
		case car := <-carGeneratorChan:
			circle.Add(car)
		}
	}
	logrus.WithFields(logrus.Fields{
		"Пропускная способность в час": model.CarTerminated,
	}).Info("Вывод")
	logrus.WithFields(logrus.Fields{
		"Средняя длин очереди": float64(model.AvgQueue) / 3600.,
	}).Info("Вывод")
}
