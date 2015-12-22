# LeastSquareCircleFit
Golang library to compute a least square circle fit

[![GoDoc](https://godoc.org/github.com/StefanSchroeder/LeastSquareCircleFit?status.png)](https://godoc.org/github.com/StefanSchroeder/LeastSquareCircleFit)
[![Go Report Card](http://goreportcard.com/badge/StefanSchroeder/LeastSquareCircleFit)](http://goreportcard.com/report/StefanSchroeder/LeastSquareCircleFit) 

This package implements a Circle Least Square Fit for a
list of 2D-coordinates

    →        ⎡ x1, x2, x3, x4, x5 ...⎤
    x  =     ⎣ y1, y2, y3, y4, y5 ...⎦	


 so that the resulting circle is a "best fit" to the points given.

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

You can find it with some googling.

 Caveats:

 There are some divisions involved which may provoke a division by zero error.
 But I didn't check out how this can be done. Perhaps
 if you supply not enough points or all the points are identical; it's definitely
 a pathological case.

![Graph](https://github.com/StefanSchroeder/LeastSquareCircleFit/blob/master/demo/least.png?raw=true)

