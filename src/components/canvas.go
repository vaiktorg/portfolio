package components

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/maxence-charriere/go-app/pkg/app"
)

type Canvas struct {
	app.Compo
	canvasCtx app.Value
	w, h      int
	margin    string
	once      sync.Once
}

func (d *Canvas) Render() app.UI {
	d.once.Do(func() {
		d.margin = "10%"
		d.w, d.h = app.Window().Size()
	})
	return app.Canvas().Class("w3-border").ContentEditable(true).
		Width(d.w/2).Height(d.h/2).Style("margin-top", d.margin)
}

func (d *Canvas) OnMount(ctx app.Context) {
	d.canvasCtx = ctx.JSSrc.Call("getContext", "2d")

	maxXSquares := 10
	maxYSquares := 10

	for y := 0; y < maxYSquares; y++ {
		for x := 0; x < maxXSquares; x++ {
			xr := int(math.Floor(float64(255 - (256/maxXSquares)*y)))
			yg := int(math.Floor(float64(255 - (256/maxXSquares)*x)))
			col := d.rgbToHex(xr, yg, 0)

			d.SetFillStyle(col)
			d.FillRect(y*25, x*25, 25, 25)
		}
	}
}

// Canvas 2D API
// =====================================================================================================================
func (d *Canvas) SetFillStyle(arg string) {
	d.canvasCtx.Set("fillStyle", arg)
}
func (d *Canvas) SetStrokeStyle(arg int) {
	d.canvasCtx.Set("strokeStyle", arg)
}
func (d *Canvas) SetGlobalAlpha(arg int) {
	d.canvasCtx.Set("globalAlpha", arg)
}

func (d *Canvas) Rect(x, y, width, height int) {
	d.canvasCtx.Call("rect", x, y, width, height)
}

func (d *Canvas) FillRect(x, y, width, height int) {
	d.canvasCtx.Call("fillRect", x, y, width, height)
}

func (d *Canvas) StrokeRect(x, y, width, height int) {
	d.canvasCtx.Call("strokeRect", x, y, width, height)
}

func (d *Canvas) ClearRect(x, y, width, height int) {
	d.canvasCtx.Call("clearRect", x, y, width, height)
}

func (d *Canvas) BeginPath() {
	d.canvasCtx.Call("beginPath")
}

func (d *Canvas) ClosePath() {
	d.canvasCtx.Call("beginPath")
}

func (d *Canvas) Fill() {
	d.canvasCtx.Call("fill")
}

func (d *Canvas) Stroke() {
	d.canvasCtx.Call("stroke")
}

func (d *Canvas) MoveTo(x, y int) {
	d.canvasCtx.Call("moveTo", x, y)
}

func (d *Canvas) LineTo(x, y int) {
	d.canvasCtx.Call("lineTo", x, y)
}

func (d *Canvas) Arc(x, y, r, startAngle, endAngle int) {
	d.canvasCtx.Call("arc", x, y, r, startAngle, endAngle)
}

func (d *Canvas) ArcTo(x0, y0, x1, y1, radius int) {
	d.canvasCtx.Call("arc", x0, y0, x1, y1, radius)
}

// quadraticCurveTo(cp1x, cp1y, x, y)
func (d *Canvas) QuadraicCurveTo() {

}

//bezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y)
func (d *Canvas) BezierCurveTo() {

}

// =====================================================================================================================

func (d *Canvas) rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
}

func (d *Canvas) hexToRGB(hexColor string) (int, int, int, error) {
	hex := hexColor[1:]
	r, err := strconv.Atoi(hex[:2])
	g, err := strconv.Atoi(hex[2:4])
	b, err := strconv.Atoi(hex[4:6])
	if err != nil {
		return 0, 0, 0, err
	}

	return r, g, b, nil
}
