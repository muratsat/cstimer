package cube

import (
	"testing"
)

func expectEq(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("%v != %v", a, b)
	}
}

func TestFaceShort(t *testing.T) {
	expectEq(t, FaceFront.Short(), "F")
	expectEq(t, FaceRight.Short(), "R")
	expectEq(t, FaceBack.Short(), "B")
	expectEq(t, FaceLeft.Short(), "L")
	expectEq(t, FaceUp.Short(), "U")
	expectEq(t, FaceDown.Short(), "D")
}

func TestFaceColor(t *testing.T) {
	expectEq(t, FaceFront.Color(), Green)
	expectEq(t, FaceRight.Color(), Red)
	expectEq(t, FaceBack.Color(), Blue)
	expectEq(t, FaceLeft.Color(), Orange)
	expectEq(t, FaceUp.Color(), White)
	expectEq(t, FaceDown.Color(), Yellow)
}

func TestRelative(t *testing.T) {
	expectEq(t, left.Opposite(), right)
	expectEq(t, top.Opposite(), bottom)
	expectEq(t, right.Opposite(), left)
	expectEq(t, bottom.Opposite(), top)
}

func TestAdjacent(t *testing.T) {
	expectEq(t, [4]Face{FaceOrange, FaceWhite, FaceRed, FaceYellow}, FaceGreen.AdjacentFaces())
	expectEq(t, [4]Face{FaceGreen, FaceWhite, FaceBlue, FaceYellow}, FaceRed.AdjacentFaces())
	expectEq(t, [4]Face{FaceRed, FaceWhite, FaceOrange, FaceYellow}, FaceBlue.AdjacentFaces())
	expectEq(t, [4]Face{FaceBlue, FaceWhite, FaceGreen, FaceYellow}, FaceOrange.AdjacentFaces())
	expectEq(t, [4]Face{FaceOrange, FaceBlue, FaceRed, FaceGreen}, FaceWhite.AdjacentFaces())
	expectEq(t, [4]Face{FaceOrange, FaceGreen, FaceRed, FaceBlue}, FaceYellow.AdjacentFaces())
}

func TestGetAdjacentIndex(t *testing.T) {
	expectEq(t, FaceGreen.GetAdjacentIndex(FaceOrange), left)
	expectEq(t, FaceGreen.GetAdjacentIndex(FaceWhite), top)
	expectEq(t, FaceGreen.GetAdjacentIndex(FaceRed), right)
	expectEq(t, FaceGreen.GetAdjacentIndex(FaceYellow), bottom)

	expectEq(t, FaceRed.GetAdjacentIndex(FaceGreen), left)
	expectEq(t, FaceRed.GetAdjacentIndex(FaceWhite), top)
	expectEq(t, FaceRed.GetAdjacentIndex(FaceBlue), right)
	expectEq(t, FaceRed.GetAdjacentIndex(FaceYellow), bottom)

	expectEq(t, FaceBlue.GetAdjacentIndex(FaceRed), left)
	expectEq(t, FaceBlue.GetAdjacentIndex(FaceWhite), top)
	expectEq(t, FaceBlue.GetAdjacentIndex(FaceOrange), right)
	expectEq(t, FaceBlue.GetAdjacentIndex(FaceYellow), bottom)

	expectEq(t, FaceOrange.GetAdjacentIndex(FaceBlue), left)
	expectEq(t, FaceOrange.GetAdjacentIndex(FaceWhite), top)
	expectEq(t, FaceOrange.GetAdjacentIndex(FaceGreen), right)
	expectEq(t, FaceOrange.GetAdjacentIndex(FaceYellow), bottom)

	expectEq(t, FaceWhite.GetAdjacentIndex(FaceOrange), left)
	expectEq(t, FaceWhite.GetAdjacentIndex(FaceBlue), top)
	expectEq(t, FaceWhite.GetAdjacentIndex(FaceRed), right)
	expectEq(t, FaceWhite.GetAdjacentIndex(FaceGreen), bottom)

	expectEq(t, FaceYellow.GetAdjacentIndex(FaceOrange), left)
	expectEq(t, FaceYellow.GetAdjacentIndex(FaceGreen), top)
	expectEq(t, FaceYellow.GetAdjacentIndex(FaceRed), right)
	expectEq(t, FaceYellow.GetAdjacentIndex(FaceBlue), bottom)
}
