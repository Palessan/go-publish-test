package maf

import (
	"math"
	"testing"

	tp "github.com/palessan/go-publish-test/util/types"
)

// assertEqual checks if expected is equal to actual, if not, it fails the test
func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Failed. Expected: %v, but got: %v", expected, actual)
	}
}

func assertEqualM(t *testing.T, expected interface{}, actual interface{}, message string) {
	if actual != expected {
		t.Errorf("Failed: %s - Expected: %v, but got: %v", message, expected, actual)
	}
}

func TestGetRandomNumberZeroToMaxMinusOne(t *testing.T) {
	for i := 1; i < 100; i++ {
		result := GetRandomNumberZeroToMaxMinusOne(2)
		assertEqualM(t, true, (result >= 0 && result < 2), "in bounds on 0 and 1")
	}

}

func TestGetDegreesFromRadians(t *testing.T) {
	assertEqualM(t, 45.0, DegreesFromRadians(math.Pi/4), "This means it will fully go right")
}

// In Go, test functions must start with the prefix Test, followed by a name that starts with a capital letter.
// testing: warning: no tests to run
// func AngleToVectorTest(t *testing.T) {
func Test1AngleToVector(t *testing.T) {
	// degrees=radians×(180/π)

	assertEqualM(t, tp.Vector2{X: 1, Y: 0}, DirectionFromRadians(0), "This means it will fully go right")
	assertEqualM(t, tp.Vector2{X: 1, Y: 0}, DirectionFromRadians(math.Pi*2), "This means it will fully go right1")

	assertEqualM(t, tp.Vector2{X: -1, Y: 0}, DirectionFromRadians(math.Pi), "This means it will fully go left")
	assertEqualM(t, tp.Vector2{X: -1, Y: 0}, DirectionFromRadians(-math.Pi), "This means it will fully go left1")

	assertEqualM(t, tp.Vector2{X: 0, Y: 1}, DirectionFromRadians(math.Pi/2), "This means it will fully go up")
	assertEqualM(t, tp.Vector2{X: 0, Y: -1}, DirectionFromRadians(-math.Pi/2), "This means it will fully go down1")

	assertEqualM(t, tp.Vector2{X: 0.707, Y: 0.707}, DirectionFromRadians(math.Pi/4), "This means it will go right down ")
}

func Test1DegreesToVector(t *testing.T) {
	assertEqualM(t, tp.Vector2{X: 1, Y: 0}, DirectionFromDegrees(0), "This means it will fully go right")
	assertEqualM(t, tp.Vector2{X: 0.707, Y: 0.707}, DirectionFromDegrees(45),
		"This means it will go right and down. Its based on sin and cos so it will not be 0.5 0.5")
	assertEqualM(t, tp.Vector2{X: 0, Y: 1}, DirectionFromDegrees(90), "This means it will fully go down1")
	assertEqualM(t, tp.Vector2{X: -1, Y: 0}, DirectionFromDegrees(180), "This means it will fully go left")
	assertEqualM(t, tp.Vector2{X: 0, Y: -1}, DirectionFromDegrees(270), "This means it will fully go up")
}

func Test1VectorToAngle(t *testing.T) {
	assertEqualM(t, float64(0), AngleFromDirection(tp.Vector2{X: 1, Y: 0}), "This means it will fully go right")
}

func Test1VectorToDegrees(t *testing.T) {
	assertEqualM(t, float64(0), VectorToDegrees(tp.Vector2{X: 1, Y: 0}), "This means it will fully go right")
	assertEqualM(t, float64(90), VectorToDegrees(tp.Vector2{X: 0, Y: 1}), "This means it will fully go down")
}
