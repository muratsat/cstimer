package cube

// Color represents a color on a rubik's cube.
type Color struct {
	number int    // order number for color
	name   string // string name for color, e.g. "Green"
	side   string // side letter for color, e.g. "F" for front, "R" for right, etc.
}

// Colors for 6 sides of a rubik's cube.
var (
	Green  = Color{0, "Green ", "F"}
	Red    = Color{1, "Red", "R"}
	Blue   = Color{2, "Blue", "B"}
	Orange = Color{3, "Orange", "L"}
	White  = Color{4, "White", "U"}
	Yellow = Color{5, "Yellow", "D"}
)

var cubeColors = [6]Color{Green, Red, Blue, Orange, White, Yellow}

// get color of a cube by number
func GetColor(colorNumber int) Color {
	if colorNumber < 0 || colorNumber > 5 {
		return Color{}
	}
	return cubeColors[colorNumber]
}
