package cube

import (
	"testing"
)

func TestGetColor(t *testing.T) {
	for i := 0; i < 6; i++ {
		color := GetColor(i)
		if color.number != i {
			t.Errorf("Expected color number %d, got %d", i, color.number)
		}
	}
}

func testColors(t *testing.T, first, second Color, expected adjacent) {
	actual := adjoinedEdge(first, second)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestAdjoinedEdge(t *testing.T) {
	// test all possible combinations
	testColors(t, Green, Orange, adjacentLeft)
	testColors(t, Green, White, adjacentTop)
	testColors(t, Green, Red, adjacentRight)
	testColors(t, Green, Yellow, adjacentBottom)

	testColors(t, Red, Green, adjacentLeft)
	testColors(t, Red, White, adjacentTop)
	testColors(t, Red, Blue, adjacentRight)
	testColors(t, Red, Yellow, adjacentBottom)

	testColors(t, Blue, Red, adjacentLeft)
	testColors(t, Blue, White, adjacentTop)
	testColors(t, Blue, Orange, adjacentRight)
	testColors(t, Blue, Yellow, adjacentBottom)

	testColors(t, Orange, Blue, adjacentLeft)
	testColors(t, Orange, White, adjacentTop)
	testColors(t, Orange, Green, adjacentRight)
	testColors(t, Orange, Yellow, adjacentBottom)

	testColors(t, Yellow, Orange, adjacentLeft)
	testColors(t, Yellow, Green, adjacentTop)
	testColors(t, Yellow, Red, adjacentRight)
	testColors(t, Yellow, Blue, adjacentBottom)

	testColors(t, White, Orange, adjacentLeft)
	testColors(t, White, Blue, adjacentTop)
	testColors(t, White, Red, adjacentRight)
	testColors(t, White, Green, adjacentBottom)
}
