package circle_test

import (
	"testing"

	"github.com/gonutz/check"
	"github.com/gonutz/circle"
)

func TestWeNeedAtLeast3PointsWithSameXAndYCounts(t *testing.T) {
	checkFail(t, nil, nil)
	checkFail(t, []float64{1}, []float64{1})
	checkFail(t, []float64{1}, []float64{1})
	checkFail(t, []float64{1}, []float64{1, 2})
	checkFail(t, []float64{1}, []float64{1})
	checkFail(t, []float64{1, 2}, []float64{1, 2})
	checkFail(t, []float64{1, 2, 3}, []float64{1, 2, 3, 4})
	checkFail(t, []float64{1, 2, 3, 4}, []float64{1, 2, 3})
}

func TestCircleFromThreePoints(t *testing.T) {
	checkCircle := func(xs, ys []float64) {
		x, y, r, err := circle.Fit(xs, ys)
		check.Eq(t, err, nil)
		check.Eq(t, x, 0)
		check.Eq(t, y, 0)
		check.Eq(t, r, 1)
	}

	checkCircle(
		[]float64{0, 1, 0},
		[]float64{1, 0, -1},
	)

	checkCircle(
		[]float64{-1, 0, 1},
		[]float64{0, 1, 0},
	)

	checkCircle(
		[]float64{0, 0, 1},
		[]float64{1, -1, 0},
	)

	checkCircle(
		[]float64{1, -1, 0},
		[]float64{0, 0, 1},
	)
}

func TestCircleFromFourPoints(t *testing.T) {
	// Points (0,1) (1,0) (0,-1) (-1,0) form a circle of radius 1 around the
	// origin.
	x, y, r, err := circle.Fit([]float64{0, 1, 0, -1}, []float64{1, 0, -1, 0})
	check.Eq(t, err, nil)
	check.Eq(t, x, 0)
	check.Eq(t, y, 0)
	check.Eq(t, r, 1)
}

func TestLineIsNotACircle(t *testing.T) {
	checkFail(t, []float64{1, 2, 3}, []float64{0, 0, 0})
	checkFail(t, []float64{1, 2, 3, 4}, []float64{0, 0, 0, 0})
	checkFail(t, []float64{0, 0, 0}, []float64{1, 2, 3})
	checkFail(t, []float64{0, 0, 0}, []float64{1, 2, 3, 4})
	checkFail(t, []float64{0, 0, 0}, []float64{0, 0, 0})
	checkFail(t, []float64{0, 0, 0, 0}, []float64{0, 0, 0, 0})
}

func checkFail(t *testing.T, xs, ys []float64) {
	t.Helper()
	_, _, _, err := circle.Fit(xs, ys)
	check.Neq(t, err, nil)
}
