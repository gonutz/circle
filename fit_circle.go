/*
Package circle implements a circle least square fit to find the circle that best
fits a list of 2D points.

See this paper for reference:

Least-Squares Circle Fit by Randy Bullock, 2006
*/
package circle

import (
	"errors"
	"math"
)

// Fit computes a least square fit circle for a list of 2D-points. It takes the
// x and y coordinates as arguments. The xs and ys slices must have the same
// length. The function returns the center and radius of the circle that best
// fits the given points.
func Fit(x, y []float64) (centerX, centerY, radius float64, err error) {
	if len(x) != len(y) {
		err = errors.New("number of x and y coordinates must be the same")
		return
	}
	if len(x) < 3 {
		err = errors.New("need at least 3 points to fit a circle")
		return
	}

	if len(x) == 3 {
		return threePointCircle(x, y)
	}

	xAvg := avg(x)
	yAvg := avg(y)

	var suu, suv, svv, suuu, svvv, suvv, svuu float64
	for i := range x {
		u := x[i] - xAvg
		v := y[i] - yAvg

		suu += u * u
		suv += u * v
		svv += v * v

		suuu += u * u * u
		svvv += v * v * v

		suvv += u * v * v
		svuu += v * u * u
	}

	e4 := 0.5 * (suuu + suvv)
	e5 := 0.5 * (svvv + svuu)

	if zero(suu*svv - suv*suv) {
		err = errors.New("points lie on a line")
		return
	}
	uc := (svv*e4 - suv*e5) / (suu*svv - suv*suv)
	vc := e4 - uc*suu
	if !zero(suv) {
		vc /= suv
	}

	centerX = uc + xAvg
	centerY = vc + yAvg
	radius = math.Sqrt(uc*uc + vc*vc + (suu+svv)/float64(len(x)))
	return
}

func avg(x []float64) float64 {
	sum := 0.0
	for _, v := range x {
		sum += v
	}
	return sum / float64(len(x))
}

func zero(x float64) bool {
	const precision = 1e-6
	return -precision <= x && x <= precision
}

func threePointCircle(x, y []float64) (centerX, centerY, radius float64, err error) {
	// http://paulbourke.net/geometry/circlesphere
	x1, x2, x3 := x[0], x[1], x[2]
	y1, y2, y3 := y[0], y[1], y[2]
	if zero(x2 - x1) {
		x2, x3 = x3, x2
		y2, y3 = y3, y2
	}
	if zero(x2-x1) || zero(x3-x2) {
		err = errors.New("points lie on a line")
		return
	}
	ma := (y2 - y1) / (x2 - x1)
	mb := (y3 - y2) / (x3 - x2)
	if zero(ma - mb) {
		err = errors.New("points lie on a line")
		return
	}
	centerX = (ma*mb*(y1-y3) + mb*(x1+x2) - ma*(x2+x3)) / (2 * (mb - ma))
	centerY = (-1.0/mb)*(centerX-(x2+x3)/2) + (y2+y3)/2
	radius = math.Sqrt(square(centerX-x1) + square(centerY-y1))
	return
}

func square(x float64) float64 {
	return x * x
}
