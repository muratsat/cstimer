package cube

import "testing"

func TestSideString(t *testing.T) {
	expectEqual("F", SideFront.String(), t)
	expectEqual("R", SideRight.String(), t)
	expectEqual("B", SideBack.String(), t)
	expectEqual("L", SideLeft.String(), t)
	expectEqual("U", SideUp.String(), t)
	expectEqual("D", SideDown.String(), t)
}
