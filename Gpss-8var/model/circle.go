package model

var AvgQueue = 0

const roadsNumber = 5

type Circle struct{
	Roads [roadsNumber][]Car
	//Circle [roadsNumber + 1][]Car
}

func(c *Circle) Add(car Car){
	c.Roads[car.EnterRoad] = append(c.Roads[car.EnterRoad], car)
}

func(c *Circle) Next(delta DeltaType) {
	//for i := roadsNumber - 1; i >= 0; i--{
	//	var roadIsEmpty =  true
	//	for j := 0; j < len(c.Circle[i]); j++{
	//		if c.Circle[i][j].Terminate(){
	//			c.Circle[i] = append(c.Circle[i][:j], c.Circle[i][j+1:]...)
	//			j--
	//		}else{
	//			if c.Circle[i][j].Next(){
	//				roadIsEmpty = false
	//				c.Circle[(i+1)%(roadsNumber + 1)] = append(c.Circle[(i+1)%(roadsNumber + 1)], c.Circle[i][j])
	//				c.Circle[i] = append(c.Circle[i][:j], c.Circle[i][j+1:]...)
	//				j--
	//			}
	//		}
	//	}
	//	if roadIsEmpty {
	//		if len(c.Roads[i]) != 0{
	//			var car Car
	//			car, c.Roads[i] = c.Roads[i][0], c.Roads[i][1:]
	//			c.Circle[(i + 1) % roadsNumber] = append(c.Circle[(i + 1) % roadsNumber], car)
	//		}
	//	}
	//	if i == 0 {
	//		for _, car := range c.Circle[roadsNumber] {
	//			c.Circle[0] = append(c.Circle[0], car)
	//		}
	//		c.Circle[roadsNumber] = []Car{}
	//	}
	//	AvgQueue += len(c.Roads[i])
	//}
	timeFixed := false
	for !timeFixed{
		timeFixed = true
		for i, Cars := range c.Roads{
			j := carNeedToGo(Cars)
			if j == -1{
				continue
			}
			c.Move(i, j)
		}
		for _, Cars := range c.Roads{
			for _, car := range Cars {
				if car.TimeFixed < delta.DeltaTime {
					timeFixed = false
				}
			}
		}
	}
	c.Add(delta.DeltaCar)
}

var CarTerminated = 0

func (c *Circle) Move(i, j int) {
	MoveTime := c.Roads[i][j].TimeForPart
	for k, _ := range c.Roads[i] {
		c.Roads[i][k].Move(MoveTime)
	}
	car := c.Roads[i][j]
	if c.Roads[i][j].AllTime > 0 {
		c.Roads[(i + 1) % roadsNumber] = append(c.Roads[(i + 1) % roadsNumber], car)
	}else{
		CarTerminated++
	}
	c.Roads[i] = append(c.Roads[i][:j], c.Roads[i][j+1:]...)
}

func carNeedToGo(Cars []Car) int{
	j := -1
	minTimeToGo := 10000.0
	minIndex := -1
	for i, car := range Cars {
		if car.TimeToGo < minTimeToGo {
			minTimeToGo = car.TimeToGo
			minIndex = i
		}
		if car.TimeToGo == .0 {
			if car.Priority {
				j = i
			}else{
				if j == -1 {
					j = i
				}
			}
		}
	}
	if j == -1 && minIndex != -1{
		for i, _ := range Cars{
			Cars[i].Move(minTimeToGo)
		}
		return minIndex
	}

	return j
}

func NewCircle() Circle{
	return Circle{}
}
