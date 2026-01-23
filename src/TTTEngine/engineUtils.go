package TTTEngine

import (
	"github.com/utkarsh-naman/3T/src/Constants"
	"github.com/utkarsh-naman/3T/src/model"
)

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
