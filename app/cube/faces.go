package cube

import "fmt"

type Face int

// Face constants
const (
	FaceFront Face = iota
	FaceRight
	FaceBack
	FaceLeft
	FaceUp
	FaceDown
)

// Face constants with their corresponding color
const (
	FaceGreen Face = iota
	FaceRed
	FaceBlue
	FaceOrange
	FaceWhite
	FaceYellow
)

// relative is a type for the relative position of a face
// relative(0) is left, relative(1) is top, relative(2) is right, relative(3) is bottom
type relative int

const (
	left relative = iota
	top
	right
	bottom
)

func (r relative) Opposite() relative {
	return (r + 2) % 4
}

const totalFaces = 6

var faces = [totalFaces]string{
	"FaceFront",
	"FaceRight",
	"FaceBack",
	"FaceLeft",
	"FaceUp",
	"FaceDown",
}

func (f Face) Color() Color {
	return Color(f)
}

func (f Face) String() string {
	if f >= FaceFront && f <= FaceDown {
		return faces[f]
	}
	return ""
}

// One letter representation of the face:
// f - front, r - right, b - back, l - left, u - up, d - down
func (f Face) Short() string {
	s := f.String()
	if s == "" {
		return s
	}
	return s[4:5]
}

// adjacentFaces is a 2D array of faces, where the first index is the face
// and the second index is the adjacent face.
// The order of the adjacent faces is clockwise from the left
var adjacentFaces = [totalFaces][4]Face{
	{FaceOrange, FaceWhite, FaceRed, FaceYellow}, // Front(Green)
	{FaceGreen, FaceWhite, FaceBlue, FaceYellow}, // Right(Red)
	{FaceRed, FaceWhite, FaceOrange, FaceYellow}, //  Back(Blue)
	{FaceBlue, FaceWhite, FaceGreen, FaceYellow}, //  Left(Orange)
	{FaceOrange, FaceBlue, FaceRed, FaceGreen},   //    Up(White)
	{FaceOrange, FaceGreen, FaceRed, FaceBlue},   //  Down(Yellow)
}

func (f Face) AdjacentFaces() [4]Face {
	if f < FaceFront || f > FaceDown {
		panic(fmt.Sprintf("Invalid face: %v", f))
	}
	return adjacentFaces[f]
}

func (f Face) GetAdjacentIndex(other Face) relative {
	if f < FaceFront || f > FaceDown {
		panic(fmt.Sprintf("Invalid face: %v", f))
	}

	for i, face := range adjacentFaces[f] {
		if face == other {
			return relative(i)
		}
	}

	return -1
}
