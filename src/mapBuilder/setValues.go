package main

import (
	"fmt"
	"slices"

	"github.com/utkarsh-naman/3T/src/Constants"
	"github.com/utkarsh-naman/3T/src/model"
)

var gamemap, _ = model.LoadGMap("bin/graph/valued/map1.ttt")
var winStates = getTerminal2(gamemap)
var loseStates []model.State

var workedHistory = make(map[model.State]bool)

func parentFromMap(state model.State) []model.State {
	var parentStates []model.State
	for stateKey := range gamemap {
		if slices.Contains(gamemap[stateKey].NextStates, state) {
			parentStates = append(parentStates, stateKey)
		}
	}
	return parentStates
}

func main() {
	wins()
	model.PrintGMap(gamemap)
	fmt.Println(len(gamemap))
	err := model.SaveGMap(gamemap, "bin/graph/valued/map3.ttt")
	if err != nil {
		return
	}
}

func wins() {
	if len(winStates) == 0 {
		return
	}
	loseStates = loseStates[:0]

	// marking the win states with +inf
	for _, winStateKey := range winStates {
		propsWin := gamemap[winStateKey]
		propsWin.Score = Constants.POSINF
		if len(gamemap[winStateKey].NextStates) != 0 {
			propsWin.WinDepth = maxDepth(winStateKey)
		} else {
			propsWin.WinDepth = 0
		}
		gamemap[winStateKey] = propsWin
	}

	for _, winStateKey := range winStates {
		// working for parent of +inf with -inf
		for _, parentStateKey := range parentFromMap(winStateKey) {
			if !workedHistory[parentStateKey] {
				parentProps := gamemap[parentStateKey]
				parentProps.Score = Constants.NEGINF
				gamemap[parentStateKey] = parentProps
				if len(parentProps.NextStates) > 0 {
					parentProps.LoseDepth = minDepth(parentStateKey) + 1
					gamemap[parentStateKey] = parentProps
				}
				loseStates = append(loseStates, parentStateKey)
				workedHistory[winStateKey] = true
			}
		}
	}
	loses()
	return
}

func loses() {
	if len(loseStates) == 0 {
		return
	}
	winStates = winStates[:0]
	for _, loseStateKey := range loseStates {
		for _, parent := range parentFromMap(loseStateKey) {
			if !workedHistory[parent] {
				if isAllNeg(gamemap[parent].NextStates) {
					winStates = append(winStates, parent)
				} else {
					if gamemap[parent].Score == 0 {
						parentProps := gamemap[parent]
						parentProps.Score = zeroScoreReset(parent)
						gamemap[parent] = parentProps
					}
				}
			}
		}
	}
	wins()
	return
}

func minDepth(state model.State) int8 {
	var mind int8 = 10
	for _, childStateKey := range gamemap[state].NextStates {
		mind = min(mind, gamemap[childStateKey].WinDepth)
	}
	return mind
}

func maxDepth(state model.State) int8 {
	var maxd int8 = 0
	for _, childStateKey := range gamemap[state].NextStates {
		maxd = max(maxd, gamemap[childStateKey].LoseDepth)
	}
	return maxd
}

func isAllNeg(children []model.State) bool {
	for _, child := range children {
		if gamemap[child].Score != Constants.NEGINF {
			return false
		}
	}
	return true
}

func getTerminal2(gamemap model.GMap) (terminalWinStates []model.State) {
	var terminalWinMask model.State = 1 << 31
	var terminalTieMask model.State = 1 << 30

	for stateKey := range gamemap {
		if stateKey&terminalWinMask == 0 { // game over
			if stateKey&terminalTieMask == 0 { // game over & not a tie, therefore def a win
				terminalWinStates = append(terminalWinStates, stateKey)
			}
		}
	}
	return
}

func zeroScoreReset(parent model.State) float32 {
	var score float32 = 0
	var children []model.State = gamemap[parent].NextStates
	//fmt.Println("for parent: ", parent)
	for _, childStateKey := range children {
		//fmt.Print("\tchild state:\t", childStateKey)
		childScore := gamemap[childStateKey].Score
		//fmt.Print("\tscore:", childScore, "\n")
		if childScore == Constants.NEGINF {
			score += 1
		} else if childScore == Constants.POSINF {
			score += -1
		} else if childScore > 0 {
			score += childScore
		}
	}
	score = score / float32(len(children))
	return score
}
