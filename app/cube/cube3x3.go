package cube

const (
	CubeSides       = 6 // number of sides on a cube
	StickerCount3X3 = 9 // number of stickers on a side
)

const (
	Clockwise        = true
	CounterClockwise = false
)

type Stickers3x3 [StickerCount3X3]Color

type Cube3x3 struct {
	faces [CubeSides]Stickers3x3
}

func NewCube3x3() Cube3x3 {
	c := Cube3x3{}
	for i := 0; i < CubeSides; i++ {
		for j := 0; j < StickerCount3X3; j++ {
			c.faces[i][j] = Color(i)
		}
	}
	return c
}

func rotate(row, col int, clockwise bool) (int, int) {
	if clockwise {
		return col, 2 - row
	}
	return 2 - col, row
}

func (c *Cube3x3) rotateFace(f Face, clockwise bool) {
	old := c.faces[f]
	for i := 0; i < StickerCount3X3; i++ {
		row, col := rotate(i/3, i%3, clockwise)
		newIndex := row*3 + col
		c.faces[f][newIndex] = old[i]
	}
}
