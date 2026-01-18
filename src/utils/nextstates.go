package utils

import (
	"slices"

	"github.com/utkarsh-naman/3T/src/Constants"
	"github.com/utkarsh-naman/3T/src/SymmetryReduction"
	"github.com/utkarsh-naman/3T/src/model"
)

func NextMoves(state model.State) []model.State {
	var NextSates []model.State
	continuum, tie, won, lost, turn, x, o, vacancy := model.Extract(state)

	if continuum != 1 {
		return NextSates
	}

	for i := 0; i < 9; i++ {
		var mask model.State = 1 << i
		if vacancy&mask != 0 {
			var nextstate model.State

			if turn == 0 { // X turn
				var board model.State = ((x ^ mask) << 18) | (o << 9) | (vacancy ^ mask)
				temporary := makeState(continuum, tie, won, lost, turn^1, board)
				nextstate = SymmetryReduction.CollapseEquivalents(finalTouch(temporary))
			} else {
				var board model.State = (x << 18) | ((o ^ mask) << 9) | (vacancy ^ mask)
				temporary := makeState(continuum, tie, won, lost, turn^1, board)
				nextstate = SymmetryReduction.CollapseEquivalents(finalTouch(temporary))
			}

			if !slices.Contains(NextSates, nextstate) {
				NextSates = append(NextSates, nextstate)
			}
		}
	}
	return NextSates
}

func wonby(player model.State) bool {
	for _, winmask := range Constants.WinMasks {
		if player&winmask == winmask {
			return true
		}
	}
	return false
}

func endAsTie(state model.State) bool {
	_, _, _, _, _, x, o, v := model.Extract(state)

	if v == 0 && !wonby(x) && !wonby(o) {
		return true
	}
	return false
}

func continueOrNot(state model.State) bool {
	_, _, _, _, _, x, o, v := model.Extract(state)

	if v == 0 || wonby(x) || wonby(o) {
		return false
	}
	return true
}

func finalTouch(temporary model.State) model.State {
	c, ti, w, l, tu, x, o, v := model.Extract(temporary)
	if continueOrNot(temporary) {
		c = 1
	} else {
		c = 0
	}

	if endAsTie(temporary) {
		ti = 1
	} else {
		ti = 0
	}

	if wonby(x) {
		w = 0
		l = 1
	}
	if wonby(o) {
		w = 1
		l = 0
	}

	return makeState(c, ti, w, l, tu, x<<18|o<<9|v)

}

func makeState(continuum, tie, won, lost, turn, board model.State) model.State {
	return continuum<<31 | tie<<30 | won<<29 | lost<<28 | turn<<27 | board
}
