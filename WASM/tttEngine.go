package main

import (
	"fmt"
	"math/rand"
)

func bestNextState(inputState State) State {
	var state = CollapseEquivalents(inputState)
	var bestScore = NEGINF
	var bestMoves []State
	var bestMove State

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

	if bestScore == NEGINF { // we can be defeated, so pick a random one that takes maximum losing step optimally
		var loseSteps int8 = -1
		var moves []State
		for _, move := range bestMoves {
			if tttmap[move].LoseDepth > loseSteps {
				loseSteps = tttmap[move].LoseDepth
				moves = append(moves, move)
			}
		}
		fmt.Println("if bestScore == NEGINF")
		fmt.Println("should play any of:\t", moves)
		bestMove = moves[rand.Intn(len(moves))]
		fmt.Println("random choice:\t", bestMove)
	} else if bestScore == POSINF { // guaranteed win so brutally defeat by picking a random one take takes minimum wining steps optimally
		var winSteps int8 = 11
		var moves []State
		for _, move := range bestMoves {
			if tttmap[move].WinDepth < winSteps {
				winSteps = tttmap[move].WinDepth
				moves = append(moves, move)
			}
		}
		fmt.Println("if bestScore == POSINF")
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

func bestMove(inputState, finalState State) int8 {
	var move int8 = 0

	//for method 2
	var board State = (1 << 18) - 1
	for _, superposition := range expandEquivalents(finalState) {

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

func expandEquivalents(state State) []State {
	var equivalents []State
	var head = state >> 27
	equivalents = append(equivalents, state)
	_, _, _, _, _, x, o, v := Extract(state)
	//fmt.Printf("expandequivalent got x o v as:\n")
	//fmt.Printf("x: \t%09b\to: %09b\tv: %09b\n", x, o, v)
	for _, f := range [][512]uint32{
		ROT90TABLE,
		ROT180TABLE,
		ROT270TABLE,
		FLIPHTABLE,
		FLIPVTABLE,
		DIAGTABLE,
		ANTIDIAGTABLE,
	} {
		eq := f[x]<<18 | f[o]<<9 | f[v]
		equivalents = append(equivalents, head<<27|State(eq))
	}
	return equivalents
}

func stateFromBoard(boardState State) State {
	var gameState = finalTouch(boardState)
	return gameState
}

func TTTNGinPlay(state State) int8 {
	return bestMove(stateFromBoard(state), bestNextState(state))
}
