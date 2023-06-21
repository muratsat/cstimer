package cube

// Color represents a color on a rubik's cube.
type Color struct {
	number int    // order number for color
	name   string // string name for color, e.g. "Green"
	side   Side   // side letter for color, e.g. "F" for front, "R" for right, etc.
}

// Colors for 6 sides of a rubik's cube.
var (
	Green  = Color{0, "Green", SideFront}
	Red    = Color{1, "Red", SideRight}
	Blue   = Color{2, "Blue", SideBack}
	Orange = Color{3, "Orange", SideLeft}
	White  = Color{4, "White", SideUp}
	Yellow = Color{5, "Yellow", SideDown}
)

var cubeColors = [6]Color{Green, Red, Blue, Orange, White, Yellow}

// get color of a cube by number
func GetColor(colorNumber int) Color {
	if colorNumber < 0 || colorNumber > 5 {
		return Color{}
	}
	return cubeColors[colorNumber]
}
