package cube

import (
	"fmt"
)

type Cube3x3 struct {
	sides [CubeSides][sideCount3x3]int
}

// NewCube3x3 returns a new Cube3x3 with the default colors.
func NewCube3x3() Cube3x3 {
	return Cube3x3{
		sides: [CubeSides][sideCount3x3]int{
			{Green, Green, Green, Green, Green, Green, Green, Green, Green},
			{Red, Red, Red, Red, Red, Red, Red, Red, Red},
			{Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue},
			{Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange},
			{White, White, White, White, White, White, White, White, White},
			{Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow},
		},
	}
}

func Rotate3x3(stickers [9]int, direction int) [9]int {
	newStickers := [9]int{}
	for i := 0; i < 9; i++ {
		row := i / 3
		col := i % 3
		newRow := col
		newCol := 2 - row
		if direction == CounterClockwise {
			newRow = 2 - col
			newCol = row
		}
		newStickers[newRow*3+newCol] = stickers[i]
	}

	return newStickers
}

func getEdge(side [9]int, position int) [3]int {
	switch position {
	case Top:
		return [3]int{side[0], side[1], side[2]}
	case Right:
		return [3]int{side[2], side[5], side[8]}
	case Bottom:
		return [3]int{side[6], side[7], side[8]}
	case Left:
		return [3]int{side[0], side[3], side[6]}
	}
	return [3]int{}
}

func replaceEdge(side [9]int, edge [3]int, position int) [9]int {
	first := 0
	second := 1
	third := 2

	switch position {
	case Top:
		first = 0
		second = 1
		third = 2
	case Right:
		first = 2
		second = 5
		third = 8
	case Bottom:
		first = 6
		second = 7
		third = 8
	case Left:
		first = 0
		second = 3
		third = 6
	}

	side[first] = edge[0]
	side[second] = edge[1]
	side[third] = edge[2]

	return side
}

// side: "F", "R", "B", "L", "U", "D"
// rotation: Clockwise(1), CounterClockwise(-1)
func (cube *Cube3x3) Rotate(side string, rotation int) {
	faceColor, err := getFaceColor(side)
	if err != nil {
		return
	}
	cube.sides[faceColor] = Rotate3x3(cube.sides[faceColor], rotation)

	var edges [4][3]int
	var positionToReplace [4]int

	for i := 0; i < 4; i++ {
		color := Adj[faceColor][i]
		j := 0
		for faceColor != Adj[color][j] {
			j = (j + 1) % 4
		}
		positionToReplace[i] = j
		edges[i] = getEdge(cube.sides[color], j)
	}

	for i := 0; i < 4; i++ {
		color := Adj[faceColor][i]
		replaceIndex := (i - rotation + 4) % 4
		edge := edges[replaceIndex]
		cube.sides[color] = replaceEdge(cube.sides[color], edge, positionToReplace[i])
	}
}

func (cube *Cube3x3) String() string {
	result := ""

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			absoluteIndex := row*3 + col
			sticker := ColoredSymbol(cube.sides[White][absoluteIndex])
			result += fmt.Sprint(sticker, " ")
		}
		result += "\n"
	}

	for row := 0; row < 3; row++ {
		for color := Green; color < White; color++ {
			for col := 0; col < 3; col++ {
				absoluteIndex := row*3 + col
				sticker := ColoredSymbol(cube.sides[color][absoluteIndex])
				result += fmt.Sprint(sticker, " ")
			}
			result += "  "
		}
		result += "\n"
	}

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			absoluteIndex := row*3 + col
			sticker := ColoredSymbol(cube.sides[Yellow][absoluteIndex])
			result += fmt.Sprint(sticker, " ")
		}
		result += "\n"
	}

	return result
}

func (cube *Cube3x3) ApplyMoves(moves []Move) {
	for _, move := range moves {
		rotation := Clockwise
		if move.count < 0 {
			rotation = CounterClockwise
		}
		side := faceByColor[move.side]
		for i := 0; i < move.count; i++ {
			cube.Rotate(side, rotation)
		}
	}
}

/*
white
2 4 4
2 4 4
2 4 4
green   red     blue    orange
4 0 0   1 1 1   2 2 5   3 3 3
4 0 0   1 1 1   2 2 5   3 3 3
4 0 0   1 1 1   2 2 5   3 3 3
yellow
0 5 5
0 5 5
0 5 5
*/
