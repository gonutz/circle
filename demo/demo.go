package main

import (
	"github.com/gonutz/fit"
	"github.com/gonutz/prototype/draw"
)

func main() {
	const windowW, windowH = 800, 600
	var x, y []float64
	draw.RunWindow("fit demo", windowW, windowH, func(window draw.Window) {
		if window.WasKeyPressed(draw.KeyEscape) {
			window.Close()
		}
		window.DrawText("Click left to add a point.\nClick right to delete all points.", 0, 0, draw.White)
		for _, c := range window.Clicks() {
			if c.Button == draw.LeftButton {
				x = append(x, float64(c.X))
				y = append(y, float64(c.Y))
			}
			if c.Button == draw.RightButton {
				x, y = nil, nil
			}
		}
		for i := range x {
			xx, yy := int(x[i]), int(y[i])
			window.DrawLine(xx-3, yy-3, xx+3, yy+3, draw.Green)
			window.DrawLine(xx-3, yy+3, xx+3, yy-3, draw.Green)
		}

		x, y, r, err := fit.Circle(x, y)
		if err != nil {
			text := "Error: " + err.Error() + "."
			_, textH := window.GetTextSize(text)
			window.DrawText(text, 0, windowH-textH, draw.Red)
		} else {
			x, y := round(x-r), round(y-r)
			size := round(r * 2)
			window.DrawEllipse(x, y, size, size, draw.Purple)
		}
	})
}

func round(x float64) int {
	return int(x + 0.5)
}
