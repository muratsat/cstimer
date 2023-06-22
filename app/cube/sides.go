package cube

// type that represents a side of the cube
// 0: front
// 1: right
// 2: back
// 3: left
// 4: up
// 5: down
type Side int

const (
	SideFront Side = iota
	SideRight
	SideBack
	SideLeft
	SideUp
	SideDown
)

const (
	sidesString = "FRBLUD"
)

func (side Side) String() string {
	color := GetColor(int(side))
	index := int(color.Side)
	return string(sidesString[index])
}

const (
	Clockwise        = true
	CounterClockwise = false
)
