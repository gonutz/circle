package LeastSquareCircleFit

/*

This package implements a Circle Least Square Fit for a
list of 2D-coordinates

  ->        x1, x2, x3, x4, x5 ...
  x  =      y1, y2, y3, y4, y5 ...
 so that the resulting circle is a "best fit to the points given.

 The only exported function is

 CalcLeastSquareCircleFit

 which takes two arrays as arguments: the x-coords in the first
 and the y-coords in the second; it returns three float64:
 the x-coord of the circle center,
 the y-coord of the circle center,
 the radius.

 Author: Stefan Schroeder
 Date  : 2013-07-01

 Implemented following the paper:

 Least-Squares Circle Fit by R. Bullock, October 24, 2006 10:22 am MDT

 Caveats:

 There are some divisions involved which may provoke a division by zero error.
 But I didn't take the time to figure out how this can be done. Perhaps
 if you supply not enough points or all the points are identical; it's definitely
 a pathological case.

*/

import "math"

func bar(in []float64) float64 {
	sum := 0.0
	for _, v := range in {
		sum = sum + v
	}
	return sum / float64(len(in))
}

func calcSuu(in []float64) float64 {
	sum := 0.0
	for _, v := range in {
		sum = sum + (v * v)
	}
	return sum
}

func calcSuuu(in []float64) float64 {
	sum := 0.0
	for _, v := range in {
		sum = sum + (v * v * v)
	}
	return sum
}

func calcSuv(in1 []float64, in2 []float64) float64 {
	sum := 0.0
	for k, _ := range in1 {
		sum = sum + in1[k]*in2[k]
	}
	return sum
}

func calcSuvv(in1 []float64, in2 []float64) float64 {
	sum := 0.0
	for k, _ := range in1 {
		sum = sum + in1[k]*in2[k]*in2[k]
	}
	return sum
}

func calc_u(in []float64, bar float64) []float64 {
	out := make([]float64, len(in))

	for k, v := range in {
		out[k] = v - bar
	}
	return out
}


// CalcLeastSquareCircleFit computes a least square fit circle for a list of 2d-coordinates.
// It takes the x and y coordinates as arguments. Obviously the two 
// argument arrays must have the same length. 
// The function returns three values: The x,y location of the circle center
// and the radius of the circle.
func CalcLeastSquareCircleFit(x []float64, y []float64) (float64, float64, float64) {

	N := len(x)

	xbar := bar(x)
	ybar := bar(y)

	u := calc_u(x, xbar)
	v := calc_u(y, ybar)

	suu := calcSuu(u)
	suv := calcSuv(u, v)
	svv := calcSuu(v)

	suuu := calcSuuu(u)
	svvv := calcSuuu(v)

	suvv := calcSuvv(u, v)
	svuu := calcSuvv(v, u)

	e4 := 0.5 * (suuu + suvv)
	e5 := 0.5 * (svvv + svuu)

	uc := (svv*e4 - suv*e5) / (suu*svv - suv*suv)
	vc := (e4 - uc*suu) / suv

	xc := uc + xbar
	yc := vc + ybar
	r := math.Sqrt(uc*uc + vc*vc + (suu+svv)/float64(N))

	return xc, yc, r
}
