package cube

import (
	"testing"
)

func TestRotateFace(t *testing.T) {
	c := NewCube3x3()
	for i := 0; i < StickerCount3X3; i++ {
		c.faces[FaceFront][i] = Color(i)
	}
	t.Log(c)

	expectEq(t, c.faces[FaceFront], Stickers3x3{0, 1, 2, 3, 4, 5, 6, 7, 8})
	c.rotateFace(FaceFront, Clockwise)
	expectEq(t, c.faces[FaceFront], Stickers3x3{6, 3, 0, 7, 4, 1, 8, 5, 2})
	c.rotateFace(FaceFront, CounterClockwise)
	c.rotateFace(FaceFront, CounterClockwise)
	expectEq(t, c.faces[FaceFront], Stickers3x3{2, 5, 8, 1, 4, 7, 0, 3, 6})
}
