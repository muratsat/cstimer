package cube

// Color represents a color on a rubik's cube.
type Color struct {
	number int    // order number for color
	name   string // string name for color, e.g. "Green"
	Side   Side   // side letter for color, e.g. "F" for front, "R" for right, etc.
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

type adjacent int

const (
	adjacentLeft adjacent = iota
	adjacentTop
	adjacentRight
	adjacentBottom
)

func (a adjacent) opposite() adjacent {
	return (a + 2) % 4
}

// get adjacent colors of a color
// in order of left, up, right, down
func adjacentColors(color Color) []Color {
	switch color {
	case Green:
		return []Color{Orange, White, Red, Yellow}
	case Red:
		return []Color{Green, White, Blue, Yellow}
	case Blue:
		return []Color{Red, White, Orange, Yellow}
	case Orange:
		return []Color{Blue, White, Green, Yellow}
	case White:
		return []Color{Orange, Blue, Red, Green}
	case Yellow:
		return []Color{Orange, Green, Red, Blue}
	}
	return []Color{}
}

// get the side of the first color that is joined to the second color
func adjoinedEdge(first, second Color) adjacent {
	adj := adjacentColors(first)
	for i := 0; i < 4; i++ {
		if adj[i] == second {
			return adjacent(i)
		}
	}
	return -1
}

func (a adjacent) edgeIndices() (int, int, int) {
	var first, second, third int
	switch a {
	case adjacentLeft:
		first, second, third = 0, 3, 6
	case adjacentTop:
		first, second, third = 0, 1, 2
	case adjacentRight:
		first, second, third = 2, 5, 8
	case adjacentBottom:
		first, second, third = 6, 7, 8
	}
	return first, second, third
}
