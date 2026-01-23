package main

//func NextMoves(state State) []State {
//	var NextSates []State
//	continuum, tie, won, lost, turn, x, o, vacancy := Extract(state)
//
//	if continuum != 1 {
//		return NextSates
//	}
//
//	for i := 0; i < 9; i++ {
//		var mask State = 1 << i
//		if vacancy&mask != 0 {
//			var nextstate State
//
//			if turn == 0 { // X turn
//				var board State = ((x ^ mask) << 18) | (o << 9) | (vacancy ^ mask)
//				temporary := makeState(continuum, tie, won, lost, turn^1, board)
//				nextstate = CollapseEquivalents(finalTouch(temporary))
//			} else {
//				var board State = (x << 18) | ((o ^ mask) << 9) | (vacancy ^ mask)
//				temporary := makeState(continuum, tie, won, lost, turn^1, board)
//				nextstate = CollapseEquivalents(finalTouch(temporary))
//			}
//
//			if !slices.Contains(NextSates, nextstate) {
//				NextSates = append(NextSates, nextstate)
//			}
//		}
//	}
//	return NextSates
//}

//func continueOrNot(state State) bool {
//	_, _, _, _, _, x, o, v := Extract(state)
//
//	if v == 0 || wonby(x) || wonby(o) {
//		return false
//	}
//	return true
//}
//
//func endAsTie(state State) bool {
//	_, _, _, _, _, x, o, v := Extract(state)
//
//	if v == 0 && !wonby(x) && !wonby(o) {
//		return true
//	}
//	return false
//}

//func wonby(player State) bool {
//	for _, winmask := range WinMasks {
//		if player&winmask == winmask {
//			return true
//		}
//	}
//	return false
//}

//func finalTouch(temporary State) State {
//	c, ti, w, l, tu, x, o, v := Extract(temporary)
//	if continueOrNot(temporary) {
//		c = 1
//	} else {
//		c = 0
//	}
//
//	if endAsTie(temporary) {
//		ti = 1
//	} else {
//		ti = 0
//	}
//
//	if wonby(x) {
//		w = 0
//		l = 1
//	}
//	if wonby(o) {
//		w = 1
//		l = 0
//	}
//
//
//	return makeState(c, ti, w, l, tu, x<<18|o<<9|v)
//
//}

//func makeState(continuum, tie, won, lost, turn, board State) State {
//	return continuum<<31 | tie<<30 | won<<29 | lost<<28 | turn<<27 | board
//}
