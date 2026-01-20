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
	//model.SaveGMap(gamemap, "bin/graph/valued/map2.ttt")
}

func wins() {
	if len(winStates) == 0 {
		return
	}
	loseStates = loseStates[:0]

	for _, stateKey := range winStates {
		props := gamemap[stateKey]
		props.Score = Constants.POSINF
		props.WinDepth = maxDepth(stateKey)
		gamemap[stateKey] = props

		for _, parentStateKey := range parentFromMap(stateKey) {
			if !workedHistory[parentStateKey] {
				parentProps := gamemap[parentStateKey]
				parentProps.Score = Constants.NEGINF
				gamemap[parentStateKey] = parentProps
				if len(parentProps.NextStates) > 0 {
					parentProps.LoseDepth = minDepth(parentStateKey) + 1
					gamemap[parentStateKey] = parentProps
				}
				loseStates = append(loseStates, parentStateKey)
				workedHistory[stateKey] = true
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
	for _, stateKey := range loseStates {
		for _, parent := range parentFromMap(stateKey) {
			if !workedHistory[parent] {
				if isAllNeg(gamemap[parent].NextStates) {
					winStates = append(winStates, parent)
				}
			}
		}
	}
	wins()
	return
}

func minDepth(state model.State) int8 {
	var mind int8 = gamemap[state].WinDepth
	for _, stateKey := range gamemap[state].NextStates {
		mind = min(mind, gamemap[stateKey].WinDepth)
	}
	return mind
}

func maxDepth(state model.State) int8 {
	var maxd int8 = gamemap[state].LoseDepth
	for _, stateKey := range gamemap[state].NextStates {
		maxd = max(maxd, gamemap[stateKey].LoseDepth)
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
