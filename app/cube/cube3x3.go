package cube

const (
	CubeSides    = 6
	sideCount3x3 = 9
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
