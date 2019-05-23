package model

import (
	"math"
)

type Event struct {
	occurredTime float64
	addingCar Car
	freeRoad int
	typeIsFreeRoad bool
}

func newEvent(occurredTime float64, addinCar Car, freeRoad int, typeIsFreeRoad bool) Event{
	return Event{occurredTime, addinCar, freeRoad, typeIsFreeRoad}
}

type Events []Event

func (e *Events) getMinTime() (Events, float64) {
	minTime := math.MaxFloat64
	for _, event := range *e {
		if event.occurredTime < minTime {
			minTime = event.occurredTime
		}
	}
	var minEvents Events
	for i := 0; i < len(*e); i++ {
		if (*e)[i].occurredTime == minTime {
			minEvents = append(minEvents, (*e)[i])
			if i == len(*e) - 1{
				*e = (*e)[:i]
			}else{
				*e = append((*e)[:i], (*e)[i+1:]...)
			}
			i--
		}
	}
	return minEvents, minTime
}