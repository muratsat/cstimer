package cube

import (
	"testing"
)

func TestRotateSide(t *testing.T) {
	cube := NewCube3x3()
	cube.sides[SideFront][0] = Red
	cube.sides[SideFront][2] = Orange
	cube.sides[SideFront][6] = Yellow
	cube.sides[SideFront][8] = White
	// R G O
	// G G G
	// Y G W
	//  to:
	// Y G R
	// G G G
	// W G O
	cube.rotateSide(SideFront, Clockwise)

	passed := true
	passed = passed && cube.sides[SideFront][0] == Yellow
	passed = passed && cube.sides[SideFront][2] == Red
	passed = passed && cube.sides[SideFront][6] == White
	passed = passed && cube.sides[SideFront][8] == Orange
	if !passed {
		t.Error("RotateSide failed: ", cube.sides[SideFront])
	}

	cube.rotateSide(SideFront, CounterClockwise)
	passed = true
	passed = passed && cube.sides[SideFront][0] == Red
	passed = passed && cube.sides[SideFront][2] == Orange
	passed = passed && cube.sides[SideFront][6] == Yellow
	passed = passed && cube.sides[SideFront][8] == White
	if !passed {
		t.Error("RotateSide failed: ", cube.sides[SideFront])
	}
}

// func TestReplaceEdge(t *testing.T) {
// }

func testEdge(t *testing.T, edge [3]Color, expected [3]Color) {
	passed := edge[0] == expected[0] && edge[1] == expected[1] && edge[2] == expected[2]
	if !passed {
		t.Error("testEdge failed: ", edge, expected)
	}
}
