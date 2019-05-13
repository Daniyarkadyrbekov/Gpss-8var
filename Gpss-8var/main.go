package main

import (
	"./generator"
	"./model"
	"github.com/sirupsen/logrus"
	"time"
)

var Cars = []model.Car{}

func main() {
	carGeneratorChan := make(chan model.Car)
	go generator.CarGenerator(carGeneratorChan)

	ticker := time.NewTicker(time.Millisecond)
	timer := time.NewTimer(3600 * time.Millisecond)
	circle := model.NewCircle()
LOOP:
	for {
		select {
		case <-ticker.C:
			circle.Next()
		case <-timer.C:
			ticker.Stop()
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
