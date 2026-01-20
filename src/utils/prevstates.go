package utils

import (
	"slices"

	"github.com/utkarsh-naman/3T/src/SymmetryReduction"
	"github.com/utkarsh-naman/3T/src/model"
)

func PrevMoves(state model.State) []model.State {
	var prevStates []model.State
	_, _, _, _, turn, x, o, vacancy := model.Extract(state)

	for i := 0; i < 9; i++ {
		var bitflipmask model.State = model.State(1) << i
		if vacancy&bitflipmask == 0 {
			var prevState model.State

			if turn == 0 { // next turn of this state is of X therefore this move was played by O so, the turn of previous state = 1
				var board model.State = (x << 18) | ((o ^ bitflipmask) << 9) | (vacancy ^ bitflipmask)
				temporary := makeState(1, 0, 0, 0, turn^1, board)
				prevState = SymmetryReduction.CollapseEquivalents(finalTouch(temporary))
			} else {
				var board model.State = ((x ^ bitflipmask) << 18) | (o << 9) | (vacancy ^ bitflipmask)
				temporary := makeState(1, 0, 0, 0, turn^1, board)
				prevState = SymmetryReduction.CollapseEquivalents(finalTouch(temporary))
			}

			if !slices.Contains(prevStates, prevState) {
				prevStates = append(prevStates, prevState)
			}
		}
	}
	return prevStates
}
