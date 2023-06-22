package cube

import "fmt"

// Color represents a color on a rubik's cube.
type Color int

// Colors for 6 sides of a rubik's cube.
const (
	Green Color = iota
	Red
	Blue
	Orange
	White
	Yellow
)

var colors = []string{
	"Green",
	"Red",
	"Blue",
	"Orange",
	"White",
	"Yellow",
}

func (c Color) String() (string, error) {
	if c >= Green && c <= Yellow {
		return colors[c], nil
	}
	return "", fmt.Errorf("unknown color: %d", c)
}
