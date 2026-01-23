package TTTEngine

import (
	"fmt"
	"math/rand"

	"github.com/utkarsh-naman/3T/src/Constants"
	"github.com/utkarsh-naman/3T/src/SymmetryReduction"
	"github.com/utkarsh-naman/3T/src/model"
)

func bestNextState(inputState model.State) model.State {
	var state = SymmetryReduction.CollapseEquivalents(inputState)
	var bestScore float32 = Constants.NEGINF
	var bestMoves []model.State
	var bestMove model.State

	if len(tttmap[state].NextStates) == 0 {
		fmt.Println("no children of state: ", state, "so quitting")
		return bestMove
	}

	for _, nextState := range tttmap[state].NextStates {
		bestScore = max(bestScore, tttmap[nextState].Score)
	}

	for _, nextState := range tttmap[state].NextStates {
		if tttmap[nextState].Score == bestScore {
			bestMoves = append(bestMoves, nextState)
		}
	}

	if bestScore == Constants.NEGINF { // we can be defeated, so pick a random one that takes maximum losing step optimally
		var loseSteps int8 = -1
		var moves []model.State
		for _, move := range bestMoves {
			if tttmap[move].LoseDepth > loseSteps {
				loseSteps = tttmap[move].LoseDepth
				moves = append(moves, move)
			}
		}
		fmt.Println("if bestScore == Constants.NEGINF")
		fmt.Println("should play any of:\t", moves)
		bestMove = moves[rand.Intn(len(moves))]
		fmt.Println("random choice:\t", bestMove)
	} else if bestScore == Constants.POSINF { // guaranteed win so brutally defeat by picking a random one take takes minimum wining steps optimally
		var winSteps int8 = 11
		var moves []model.State
		for _, move := range bestMoves {
			if tttmap[move].WinDepth < winSteps {
				winSteps = tttmap[move].WinDepth
				moves = append(moves, move)
			}
		}
		fmt.Println("if bestScore == Constants.POSINF")
		fmt.Println("should play any of:\t", moves)
		bestMove = moves[rand.Intn(len(moves))]
		fmt.Println("random choice:\t", bestMove)
	} else { // when max score is 0, so we can tie, a random one
		fmt.Printf("should play any of:\t%032b, \n", bestMoves)
		bestMove = bestMoves[rand.Intn(len(bestMoves))]
		fmt.Println("random choice:\t", bestMove)
	}
	fmt.Printf("\nmoved played by NextBestMove func:\t%018b\n", bestMove)
	return bestMove
}

func bestMove(inputState, finalState model.State) int8 {

	//if inputState == 2147484159 {
	//	return 9
	//}
	var move int8 = 0
	//for method 1
	//var board model.State = (1 << 9) - 1

	//for method 2
	var board model.State = (1 << 18) - 1
	for _, superposition := range expandEquivalents(finalState) {
		//var turnmask model.State = 1 << 27
		//var inputboardx = (inputState >> 18) & board
		//var inputboardo = (inputState >> 9) & board
		//var inputboardv = (inputState >> 9) & board
		//
		//for i := 0; i < int(inputboardv&(inputboardv-1)); i++ {
		//	var flipbitmask model.State = 1 << i
		//	if inputboardv&flipbitmask != 0 {
		//		var tempboardv = inputboardv | flipbitmask
		//		var nextBoard model.State
		//		if (inputState>>27)&1 == 0 {
		//			nextBoard = ((inputboardx | flipbitmask) << 18) | (inputboardo << 9) | tempboardv
		//		} else {
		//			nextBoard = inputboardx | ((inputboardo | flipbitmask) << 9) | tempboardv
		//		}
		//
		//		if (nextBoard^superposition)&(nextBoard^superposition) == 2 {
		//			return int8(i) + 1
		//		}
		//
		//	}
		//}

		//method 2
		for i := 0; i < 18; i++ {
			//fmt.Printf("superposition:\t%032b\n", superposition)
			var inputboardxo = (inputState >> 9) & board
			var superboardxo = (superposition >> 9) & board
			//fmt.Println("checking for:")
			//fmt.Printf("inuputboard:\t%018b\n", inputboardxo)
			//fmt.Printf("superboard:\t%018b\n", superboardxo)
			//fmt.Printf("result:\t%018b\n", inputboardxo^superboardxo)
			if (inputboardxo)^(superboardxo) == (1 << i) {
				//fmt.Println("returning move from superposition matched")
				return 10 - ((int8(i))%9 + 1)
			}
		}
	}
	return move
}

func expandEquivalents(state model.State) []model.State {
	var equivalents []model.State
	var head = state >> 27
	equivalents = append(equivalents, state)
	_, _, _, _, _, x, o, v := model.Extract(state)
	//fmt.Printf("expandequivalent got x o v as:\n")
	//fmt.Printf("x: \t%09b\to: %09b\tv: %09b\n", x, o, v)
	for _, f := range [][512]uint32{
		Constants.ROT90TABLE,
		Constants.ROT180TABLE,
		Constants.ROT270TABLE,
		Constants.FLIPHTABLE,
		Constants.FLIPVTABLE,
		Constants.DIAGTABLE,
		Constants.ANTIDIAGTABLE,
	} {
		eq := f[x]<<18 | f[o]<<9 | f[v]
		equivalents = append(equivalents, head<<27|model.State(eq))
	}
	return equivalents
}

func TTTNGinPlay(state model.State) int8 {
	return bestMove(state, bestNextState(state))
}
