package cube

import "math/rand"

type movetype int

const (
	movetypeNormal = iota
	movetypePrime
	movetypeDouble
)

type Move struct {
	// Side to rotate (F, R, B, L, U, D)
	side int
	// type of rotation (normal, prime, double)
	moveType int
}

func (m Move) String() string {
	color := GetColor(m.side)
	switch m.moveType {
	case movetypeNormal:
		return color.side.String()
	case movetypePrime:
		return color.side.String() + "'"
	case movetypeDouble:
		return "2" + color.side.String()
	}
	return ""
}

func randomMove(previous Move) Move {
	side := rand.Intn(6)
	moveType := rand.Intn(3)
	for side == previous.side {
		side = rand.Intn(6)
	}
	return Move{side: side, moveType: moveType}
}

func GenerateMoves() []Move {
	var moves []Move
	maxMoves := 20
	var lastMove Move
	for i := 0; i < maxMoves; i++ {
		move := randomMove(lastMove)
		lastMove = move
		moves = append(moves, move)
	}
	return moves
}

func ScrambleString(moves []Move) string {
	total := len(moves)
	result := ""

	for i := 0; i < total; i++ {
		move := moves[i]
		result += move.String()
		if i < total-1 {
			result += " "
		}
	}

	return result
}

func GenerateScramble() (string, []Move) {
	moves := GenerateMoves()
	scramble := ScrambleString(moves)
	return scramble, moves
}
