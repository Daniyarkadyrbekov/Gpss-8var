package main

import (
	"./model"

	"fmt"
)

func main() {
	//timeoutChan := make(chan struct{}, 2)

	circle := model.NewCircle()
	circle.Next()
//LOOP:
//	for {
//		select {
//		case <-timeoutChan:
//			break LOOP
//		}
//	}

	fmt.Println("carGenerated ", model.CarGenerated)
	fmt.Println("carterminated ", model.CarTerminated)
	//logrus.WithFields(logrus.Fields{
	//	"Пропускная способность": float64(model.CarTerminated) / 3600,
	//}).Info("Вывод")
	//logrus.WithFields(logrus.Fields{
	//	"Средняя длина очереди": float64(model.AvgQueue) / 3600,
	//}).Info("Вывод")
	//logrus.WithFields(logrus.Fields{
	//	//	"сгенерированно машин": generator.CarGenerated,
	//	//}).Info("Вывод")
}
