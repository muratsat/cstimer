package cube

import (
	"testing"
)

func TestRandomMove(t *testing.T) {
	lastMove := Move{side: 0, moveType: 0}

	wrong := 0
	total := 100
	for i := 0; i < total; i++ {
		move := randomMove(lastMove)
		if move.side == lastMove.side {
			// t.Errorf("move.side == lastMove.side: %d == %d", move.side, lastMove.side)
			wrong++
		}
	}

	if wrong > 0 {
		t.Errorf("result: %d/%d", wrong, total)
	}
}

func expectEqual(expected, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestMoveString(t *testing.T) {
	var moves []Move
	for i := 0; i < 6; i++ {
		for j := 0; j < 3; j++ {
			moves = append(moves, Move{side: i, moveType: j})
		}
	}
	for _, move := range moves {
		switch move.moveType {
		case movetypeNormal:
			expectEqual(move.String(), GetColor(move.side).side, t)
		case movetypePrime:
			expectEqual(move.String(), GetColor(move.side).side+"'", t)
		case movetypeDouble:
			expectEqual(move.String(), "2"+GetColor(move.side).side, t)
		}
	}
}

func TestScrambleString(t *testing.T) {
	moves := []Move{
		{side: 0, moveType: 0},
		{side: 1, moveType: 1},
		{side: 2, moveType: 2},
		{side: 3, moveType: 0},
		{side: 4, moveType: 1},
		{side: 5, moveType: 2},
	}

	expectEqual("F R' 2B L U' 2D", ScrambleString(moves), t)
}
