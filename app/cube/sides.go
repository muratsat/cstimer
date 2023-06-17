package cube

import (
	"errors"
	"strings"
)

const (
	Clockwise        int = 1
	CounterClockwise int = -1
)

const (
	Left   = 0
	Top    = 1
	Right  = 2
	Bottom = 3
)

func positionToString(position int) string {
	switch position {
	case Left:
		return "Left"
	case Top:
		return "Top"
	case Right:
		return "Right"
	case Bottom:
		return "Bottom"
	}
	return ""
}

var colorToTheLeft = map[int]int{
	Green:  Orange,
	Red:    Green,
	Blue:   Red,
	Orange: Blue,
	White:  Orange,
	Yellow: Orange,
}

// face for 3x3 when at the green face with white on top
var face map[string]int = map[string]int{
	"F": Green,
	"R": Red,
	"B": Blue,
	"L": Orange,
	"U": White,
	"D": Yellow,
}

var faceByColor [6]string = [6]string{
	"F",
	"R",
	"B",
	"L",
	"U",
	"D",
}

func getFaceColor(side string) (int, error) {
	upper := strings.ToUpper(side)

	if _, ok := face[upper]; !ok {
		return -1, errors.New("invalid side")
	}
	return face[upper], nil
}

// adjacent colors in order of left, top, right, bottom
var Adj = [6][4]int{
	{Orange, White, Red, Yellow},
	{Green, White, Blue, Yellow},
	{Red, White, Orange, Yellow},
	{Blue, White, Green, Yellow},
	{Orange, Blue, Red, Green},
	{Orange, Green, Red, Blue},
}

func AreOppositeColors(color1, color2 int) bool {
	if color1 == color2 {
		return false
	}

	for i := 0; i < 4; i++ {
		if Adj[color1][i] == color2 {
			return false
		}
	}

	return true
}
