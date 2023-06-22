package cube

const (
	CubeSides    = 6 // number of sides on a cube
	sideCount3x3 = 9 // number of stickers on a side
)

type Cube3x3 struct {
	sides [CubeSides][sideCount3x3]Color
}

// NewCube3x3 returns a new Cube3x3 with the default colors.
func NewCube3x3() Cube3x3 {
	cube := Cube3x3{}
	for i := 0; i < CubeSides; i++ {
		for j := 0; j < sideCount3x3; j++ {
			cube.sides[i][j] = GetColor(i)
		}
	}
	return cube
}

func rotate(row, col int, clockwise bool) (int, int) {
	if clockwise {
		return col, 2 - row
	}
	return 2 - col, row
}

func (cube *Cube3x3) rotateSide(side Side, clockwise bool) {
	// rotate the side
	newSide := cube.sides[side]
	for i := 0; i < sideCount3x3; i++ {
		row, col := rotate(i/3, i%3, clockwise)
		index := row*3 + col
		newSide[index] = cube.sides[side][i]
	}
	cube.sides[side] = newSide
}

func (cube *Cube3x3) getEdge(side Side, edge adjacent) [3]Color {
	first, second, third := edge.edgeIndices()
	return [3]Color{
		cube.sides[side][first],
		cube.sides[side][second],
		cube.sides[side][third],
	}
}

func (cube *Cube3x3) replaceEdge(side Side, edge adjacent, colors [3]Color) {
	first, second, third := edge.edgeIndices()
	cube.sides[side][first] = colors[0]
	cube.sides[side][second] = colors[1]
	cube.sides[side][third] = colors[2]
}

func (cube *Cube3x3) rotateAdjacentSides(side Side, clockwise bool) {
	// rotate the adjacent sides
	color := GetColor(int(side))
	adjColors := adjacentColors(color)
	var edges [4][3]Color
	for i := 0; i < 4; i++ {
		adj := adjoinedEdge(adjColors[i], color)
		edge := cube.getEdge(adjColors[i].Side, adj)
		edges[i] = edge
	}

	for i := 0; i < 4; i++ {
		j := (i + 1) % 4
		if clockwise {
			j = (i + 3) % 4
		}
		adj := adjoinedEdge(adjColors[i], color)
		cube.replaceEdge(adjColors[i].Side, adj, edges[j])
	}
}

// rotate the side and the adjacent sides
func (cube *Cube3x3) Rotate(side Side, clockwise bool) {
	// rotate the side
	cube.rotateSide(side, clockwise)

	// rotate the adjacent sides
	cube.rotateAdjacentSides(side, clockwise)
}

func (cube *Cube3x3) String() string {
	result := "\n"
	for i := 0; i < CubeSides; i++ {
		for j := 0; j < sideCount3x3; j++ {
			result += cube.sides[i][j].name + " "
			if j%3 == 2 {
				result += "\n"
			}
		}
		result += "\n"
	}
	return result
}
