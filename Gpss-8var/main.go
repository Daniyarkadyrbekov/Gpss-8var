package main

import (
	"./model"
	"fmt"
)

func main() {
	circle := model.NewCircle()
	circle.Next()

	outPutResults()
}

func outPutResults() {
	fmt.Println("carGenerated ", model.CarGenerated)
	fmt.Println("carterminated ", model.CarTerminated)
	fmt.Println("carQueue ", model.AvgQueue / model.Iterations)
}
