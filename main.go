package main

import (
	"github.com/utkarsh-naman/3T/src/model"
	//"github.com/utkarsh-naman/3T/src/utils"
)

func main() {
	//var start model.State = 0b10000000000000000000000111111111
	gmap, _ := model.LoadGMap("bin/graph/unvalued/map0.ttt")
	model.PrintGMap(gmap)
	//fmt.Printf("Cont\tTie?\tWon\tLost\tTurn\tX occupancy\tO occupancy\tVacant\n")

	//for _, x := range utils.NextMoves(start) {
	//	model.PrintState(x)
	//}

}
