package main

import (
	"github.com/utkarsh-naman/3T/src/Constants"
	"github.com/utkarsh-naman/3T/src/model"
)

func main() {
	gamemap, _ := model.LoadGMap("bin/graph/unvalued/map0.ttt")
	setTerminal(gamemap)
}

func setTerminal(gamemap model.GMap) {
	terminalWinState := getTerminal(gamemap)

	for _, stateKey := range terminalWinState {
		props := gamemap[stateKey]
		props.Score = Constants.POSINF
		props.WinDepth = 0
		gamemap[stateKey] = props
	}

	//model.PrintGMap(gamemap)
	model.SaveGMap(gamemap, "bin/graph/valued/map1.ttt")
}

func getTerminal(gamemap model.GMap) (terminalWinStates []model.State) {
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
