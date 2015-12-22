package main

import "fmt"
import least "github.com/StefanSchroeder/LeastSquareCircleFit"
import "math"

func main() {
	x := []float64{0.0, 0.5, 1.0, 1.5, 2.0, 2.5, 3.0}
	y := []float64{0.0, 0.25, 1.0, 2.25, 4.0, 6.25, 9.0}

	xc, yc, r := least.CalcLeastSquareCircleFit(x, y)

	for i := range x {
		fmt.Printf("%v %v # data\n", x[i], y[i])
	}
	fmt.Printf("\n\n%v %v # center\n\n\n", xc, yc)
	for i := 0; i < 360; i += 10 {
		xi := xc + r*math.Sin(float64(i)*math.Pi/180.0)
		yi := yc + r*math.Cos(float64(i)*math.Pi/180.0)
		fmt.Printf("%v %v #circle\n", xi, yi)
	}
	fmt.Printf("\n\n%v # radius\n\n\n", r)
}
