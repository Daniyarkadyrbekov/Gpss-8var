package model

import (
	"math"
	"fmt"
	"../generator"
)

var AvgQueue = .0
var CarTerminated = 0

const roadsNumber = 5
var timeNow = .0

type Circle struct{
	futureEvents Events
	currentEvents Events
	roadIsFree []bool
}

func(c *Circle) Add(delta DeltaType){
	c.futureEvents = append(c.futureEvents, Event{delta.DeltaTime, delta.DeltaCar, 0, false})
}

var CarGenerated = 1

func(c *Circle) Next() {
	lastGeneratedTime := generator.ExponensialCarGenerator()
	deltaType := DeltaType{NewCar(generator.TimeGenerator(), generator.RoadGenerator(), generator.RoadGenerator()), timeNow}
	c.Add(deltaType)
	for timeNow < 36.0 {
		if timeNow >= lastGeneratedTime {
			lastGeneratedTime += generator.ExponensialCarGenerator()
			//CarGenerated++
			//deltaType := DeltaType{NewCar(generator.TimeGenerator(), generator.RoadGenerator(), generator.RoadGenerator()), timeNow}
			//c.Add(deltaType)
		}
		minEvents, minTime := c.futureEvents.getMinTime()
		if minTime  != math.MaxFloat64{
			timeNow = minTime
			c.currentEvents = append(c.currentEvents, minEvents...)
			c.processCurrentEvents()
		}
	}
	//timer <- struct{}{}
}

func (c *Circle) processCurrentEvents() {
	fmt.Printf("timeNow = %v\n", timeNow)
	fmt.Printf("roadIsFree = %v\n", c.roadIsFree)
	fmt.Printf("futureEvents = %v\n", c.futureEvents)
	fmt.Printf("CurrentEvents = %v\n", c.currentEvents)
	c.freeRoads()
	c.processInCircle()
}

func (c *Circle) freeRoads() {
	for i, event := range c.currentEvents{
		if event.typeIsFreeRoad{
			fmt.Printf("free Road %v\n", event.freeRoad)
			c.roadIsFree[event.freeRoad] = true
			if i == len(c.currentEvents) - 1{
				c.currentEvents = c.currentEvents[:i]
			}else{
				c.currentEvents  = append(c.currentEvents [:i], c.currentEvents [i+1:]...)
			}
		}
	}
}

func (c *Circle) processInCircle() {
	fmt.Println("\nprocess===")
	fmt.Printf("freeRoads = %v\n", c.roadIsFree)
	fmt.Printf("curentEvents = %v\n\n\n", c.currentEvents)
	for i := 0; i < len(c.currentEvents); i++{
		fmt.Printf("carTime = %v", c.currentEvents[i].addingCar.timePerRoad)
		//if event.addingCar.prioritet {
			if c.roadIsFree[c.currentEvents[i].addingCar.in] {
				c.roadIsFree[c.currentEvents[i].addingCar.in] = false
				occurredTime := timeNow + c.currentEvents[i].addingCar.timePerRoad

				event := Event{occurredTime:occurredTime, freeRoad: c.currentEvents[i].addingCar.in, typeIsFreeRoad:true}
				c.futureEvents = append(c.futureEvents, event)

				event.addingCar.in = (event.addingCar.in + 1) % roadsNumber
				if event.addingCar.in != event.addingCar.out{
					event := Event{occurredTime:occurredTime, freeRoad: event.addingCar.in, typeIsFreeRoad:false, addingCar:event.addingCar}
					c.futureEvents = append(c.futureEvents, event)
				}else{
					CarTerminated++
				}

				if i == len(c.currentEvents) - 1{
					c.currentEvents = c.currentEvents[:i]
				}else{
					c.currentEvents  = append(c.currentEvents [:i], c.currentEvents [i+1:]...)
				}
				i--
				//CarTerminated++
			}
		//}
	}
}

func (c *Circle) processInRoad() {

}

func NewCircle() Circle{
	return Circle{roadIsFree:[]bool{true, true, true, true, true}}
}
