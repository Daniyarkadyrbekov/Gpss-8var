package model

const roadsNumber = 5

type Circle struct{
	Roads [roadsNumber][]Car
	Circle [roadsNumber + 1][]Car
}

func(c *Circle) Add(car Car){
	c.Roads[car.EnterRoad] = append(c.Roads[car.EnterRoad], car)
	//logrus.WithFields(logrus.Fields{
	//	"roads": c.Roads,
	//	"circle": c.Circle,
	//}).Info("add internal")
}

func(c *Circle) Next() {
	for i := roadsNumber - 1; i >= 0; i--{
		//if len(c.Circle[i]) != 0 {
		//	var car *Car
		//	car = &c.Circle[i][0]
		//	if car.Terminate() {
		//		c.Circle[i] = c.Circle[i][1:]
		//	}else{
		//		if car.Next() {
		//			c.Circle[i] = c.Circle[i][1:]
		//			c.Circle[(i+1)%(roadsNumber + 1)] = append(c.Circle[(i+1)%(roadsNumber + 1)], *car)
		//		}
		//	}
		//}else{
		//	if len(c.Roads[i]) != 0{
		//		var car Car
		//		car, c.Roads[i] = c.Roads[i][0], c.Roads[i][1:]
		//		c.Circle[(i + 1) % roadsNumber] = append(c.Circle[(i + 1) % roadsNumber], car)
		//	}
		//}
		var roadIsEmpty =  true
		for j := 0; j < len(c.Circle[i]); j++{
			//logrus.WithFields(logrus.Fields{
			//	"i": i,
			//	"j": j,
			//	"circle": c.Circle,
			//}).Info("debug")
			if c.Circle[i][j].Terminate(){
				c.Circle[i] = append(c.Circle[i][:j], c.Circle[i][j+1:]...)
				j--
			}else{
				if c.Circle[i][j].Next(){
					roadIsEmpty = false
					c.Circle[(i+1)%(roadsNumber + 1)] = append(c.Circle[(i+1)%(roadsNumber + 1)], c.Circle[i][j])
					c.Circle[i] = append(c.Circle[i][:j], c.Circle[i][j+1:]...)
					j--
				}
			}
		}
		if roadIsEmpty {
			if len(c.Roads[i]) != 0{
				var car Car
				car, c.Roads[i] = c.Roads[i][0], c.Roads[i][1:]
				c.Circle[(i + 1) % roadsNumber] = append(c.Circle[(i + 1) % roadsNumber], car)
			}
		}
		if i == 0 {
			for _, car := range c.Circle[roadsNumber] {
				c.Circle[0] = append(c.Circle[0], car)
			}
			c.Circle[roadsNumber] = []Car{}
		}
	}
}

func NewCircle() Circle{
	return Circle{}
}
