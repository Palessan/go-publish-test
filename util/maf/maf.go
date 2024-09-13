package maf

import (
	"image"
	"math"
	"math/rand/v2"

	t "github.com/palessan/go-publish-test/util/types"
)

const (
	SMOOTHNESS = 0.20
)

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Generics
func RoundTo4[N ~float32 | ~float64](value N) N {
	return N(math.Round(float64(value)*10000) / 10000)
}
func RoundTo3[N ~float32 | ~float64](value N) N {
	return N(math.Round(float64(value)*1000) / 1000)
}

// GetRandomNumber returns a random integer between min and max (inclusive)
func GetRandomNumberMinMaxIncluded(min, max int) int {
	return rand.IntN(max+1-min) + min
}

// GetRandomNumberZeroToMaxMinusOne returns a random integer between 0 and max (inclusive)
func GetRandomNumberZeroToMaxMinusOne(max int) int {
	return rand.IntN(max)
}

func GetRandomNumberFrom1ToMaxIncluded(max int) int {
	return rand.IntN(max) + 1
}

func GetRandom[L any](list []L) L {
	return list[GetRandomNumberZeroToMaxMinusOne(len(list))]
}

// ClerpAngle performs a linear interpolation to smooth transitions between angles
// It normalizes the angle to the range −𝜋 to π by repeatedly adding or subtracting 2𝜋
// until the angle falls within the desired range.
func ClerpAngle(from, to float64) float64 {
	// Wrap the angle difference to the range -pi to pi
	delta := WrapAngle(to - from)
	// Calculate the transition step based on the smoothness
	t := delta * SMOOTHNESS
	// Return the interpolated angle
	return from + t
}
func ClerpDegrees(from, to float64) float64 {
	// Wrap the angle difference to the range -pi to pi
	delta := WrapDegrees(to - from)
	// Calculate the transition step based on the smoothness
	t := delta * SMOOTHNESS
	// Return the interpolated angle
	return from + t
}

// WrapAngle reduces the given angle to a value between -pi and pi
// The Clerp function calls WrapAngle to ensure the angle difference
// is normalized before performing the linear interpolation.
func WrapAngle(angle float64) float64 {
	for angle > math.Pi {
		angle -= 2 * math.Pi
	}
	for angle < -math.Pi {
		angle += 2 * math.Pi
	}
	return angle
}

// WrapAngle reduces the given angle to a value between -180° and 180°
// The Clerp function calls WrapAngle to ensure the angle difference
// is normalized before performing the linear interpolation.
func WrapDegrees(angle float64) float64 {
	for angle > 180 {
		angle -= 360
	}
	for angle < -180 {
		angle += 360
	}
	return angle
}

// all images are created having angle as 0 and looking up.
// to actually have the same degree orientation.
// https://www.wyzant.com/resources/lessons/math/trigonometry/unit-circle/
// 				 						(90° / -270°) (0, 1) (π/2 = 1.5708 / -3*π/2 radians = -4.71239)
// 																|
// 																|
// 																|
// 																|
// (180° / -180°)  (-1, 0) (π = 3.14159/ -π = -3.14159)	--------+---------  (0° / -0° / 360° / -360°) (1, 0) (0*π = 0)
// 																|
// 																|
// 																|
// 																|
// 																|
// 									(270° / -90°) (0, -1) (3*π/2 radians = 4.71239 / -π/2 = -1.5708)

// please note though that in ebiten the Y axis is reverse!!!
// +--------- +X
// |
// |
// |
// |
// +Y

// So the above is translated to:
// 				 					(270° / -90°) (0, -1) (3*π/2 radians = 4.71239 / -π/2 = -1.5708)
// 																|
// 																|
// 																|
// 																|
// (180° / -180°)  (-1, 0) (π = 3.14159/ -π = -3.14159)	--------+---------  (0° / -0° / 360° / -360°) (1, 0) (0*π = 0)
// 																|
// 																|
// 																|
// 																|
// 																|
// 									(90° / -270°) (0, 1) (π/2 = 1.5708 / -3*π/2 radians = -4.71239)

// To convert degrees into directions (x, y) and radians:
// 0 degrees / right:
//    - Direction: (1, 0)
//    - Radians: 0*π radians or -π = -3.14159

// 90 degrees / -270 degrees / down:
//    - Direction: (0, -1)
//    - Radians: 3*π/2 radians = 4.71239 or -π/2 = -1.5708

// 180 degrees/-180 degrees / left:
//    - Direction: (-1, 0)
//    - Radians: π radians = 3.14159 or -π radians = -3.14159

// 270 degrees / -90 degrees / up:
//    - Direction: (0, 1)
//    - Radians: π/2 = 1.5708

func VectorToDegrees(vector t.Vector2) float64 {
	return RoundTo4(DegreesFromRadians(float64(math.Atan2(float64(vector.Y), float64(vector.X)))))
}

// DirectionFromRadians converts an angle to a vector
func DirectionFromRadians(angle float64) t.Vector2 {
	return t.Vector2{
		X: float64(RoundTo3(math.Cos(angle))),
		Y: float64(RoundTo3(math.Sin(angle))),
	}
}

// degrees=radians×(180/π)
// DegreesFromRadians converts an angle from radians to degrees
func DegreesFromRadians(radians float64) float64 {
	return RoundTo4(radians * (180 / math.Pi))
}

// RadiansFromDegrees converts an angle from degrees to radians
func RadiansFromDegrees(degrees float64) float64 {
	return RoundTo4(degrees * (math.Pi / 180))
}

// AngleToVector converts an angle in degrees to a vector
func DirectionFromDegrees(angleDegrees float64) t.Vector2 {
	angleRadians := RadiansFromDegrees(angleDegrees)
	return t.Vector2{
		X: float64(RoundTo3(math.Cos(angleRadians))),
		Y: float64(RoundTo3(math.Sin(angleRadians))),
	}
}

// AngleFromDirection converts a vector to an angle
// x and y will move from -1 to 1, which shows the angle
func AngleFromDirection(vector t.Vector2) float64 {
	return RoundTo4(float64(math.Atan2(float64(vector.Y), float64(vector.X))))
}

// AngleFromTwoDirections calculates the angle between two vectors
func AngleFromTwoDirections(from, to t.Vector2) float64 {
	return RoundTo4(float64(math.Atan2(float64(from.Y-to.Y), float64(from.X-to.X))))
}

// AngleFromTwoDirectionsReversed calculates the angle between two vectors (reversed order)
func AngleFromTwoDirectionsReversed(to, from t.Vector2) float64 {
	return RoundTo4(float64(math.Atan2(float64(from.Y-to.Y), float64(from.X-to.X))))
}

// gets middle of rectangle (size/2) for rotating.
// X and Y added to the position(upper left) of the rectangle
func OriginFromRectangle(r image.Rectangle) t.Vector2 {
	return t.Vector2{X: float64(r.Dx()) / 2, Y: float64(r.Dy()) / 2}
}

func OriginFromEbitenImage(rect *t.Rectangle) t.Vector2 {
	return t.Vector2{X: float64(rect.Dx()) / 2, Y: float64(rect.Dy()) / 2}
}

// t.Vector2Subtract - Subtract two vectors (v1 - v2)
func Vector2Subtract(v1, v2 t.Vector2) t.Vector2 {
	return t.Vector2{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

// CheckCollisionRecs checks collision between two rectangles
// deprecated use rlRect.ToImageRectangle
func RectFromRLRect(rlRect t.RLRectangle) image.Rectangle {
	min := image.Point{X: int(rlRect.X), Y: int(rlRect.Y)}
	max := image.Point{X: int(rlRect.X + rlRect.Width), Y: int(rlRect.Y + rlRect.Height)}
	return image.Rectangle{Min: min, Max: max}
}

// CheckCollisionRecs checks collision between two rectangles
func CheckCollisionRecs(rect1, rect2 image.Rectangle) bool {
	return rect1.Overlaps(rect2)
}
