package cube

import "math/rand"

type Move struct {
	// Side to rotate (F, R, B, L, U, D)
	side int
	// Number of turns, if negative, number of CounterClockwise turns
	count int
}

func randomSide() int {
	return rand.Intn(6)
}

func randomDirection() int {
	max := 100
	if rand.Intn(max) < max/4 {
		return Clockwise
	}
	return CounterClockwise
}

func GenerateMoves() []Move {
	var moves []Move
	count := 0
	maxMoves := 25
	for i := 0; i < maxMoves; i++ {
		side := randomSide()
		direction := randomDirection()

		if count > 0 && side == moves[count-1].side {
			moves[count-1].count += direction
			continue
		}
		moves = append(moves, Move{side: side, count: 1})
		count++
	}
	return moves
}

func ScrambleString(moves []Move) string {
	total := len(moves)
	if total == 0 {
		return ""
	}

	result := ""

	for i := 0; i < total; i++ {
		count := moves[i].count
		if count == 0 {
			continue
		}
		if count%2 == 0 {
			result += "2"
		}
		result += faceByColor[moves[i].side]
		if count < 0 && count%2 != 0 {
			result += "'"
		}
		result += " "
	}

	return result
}

func GenerateScramble() string {
	moves := GenerateMoves()
	scramble := ScrambleString(moves)
	return scramble
}
