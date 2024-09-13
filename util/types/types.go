package types

import (
	"image"
	"strconv"
)

type Vector2 struct {
	X, Y float64
}

// Rectangle type
type RLRectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (r RLRectangle) ToImageRectangle() image.Rectangle {
	min := image.Point{X: int(r.X), Y: int(r.Y)}
	max := image.Point{X: int(r.X + r.Width), Y: int(r.Y + r.Height)}
	return image.Rectangle{Min: min, Max: max}
}

// A Point is an X, Y coordinate pair. The axes increase right and down.
type Point struct {
	X, Y int
}

// String returns a string representation of p like "(3,4)".
func (p Point) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

// A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y.
// It is well-formed if Min.X <= Max.X and likewise for Y. Points are always
// well-formed. A rectangle's methods always return well-formed outputs for
// well-formed inputs.
//
// A Rectangle is also an [Image] whose bounds are the rectangle itself. At
// returns color.Opaque for points in the rectangle and color.Transparent
// otherwise.
type Rectangle struct {
	Min, Max Point
}

// String returns a string representation of r like "(3,4)-(6,5)".
func (r Rectangle) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Dx returns r's width.
func (r Rectangle) Dx() int {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r Rectangle) Dy() int {
	return r.Max.Y - r.Min.Y
}

// Size returns r's width and height.
func (r Rectangle) Size() Point {
	return Point{
		r.Max.X - r.Min.X,
		r.Max.Y - r.Min.Y,
	}
}
