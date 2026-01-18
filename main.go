package main

import (
	"fmt"

	"github.com/utkarsh-naman/3T/src/model"
	"github.com/utkarsh-naman/3T/src/utils"
)

func main() {
	var start model.State = 0b10001000000001000000000111111110
	fmt.Printf("Cont\tTie?\tWon\tLost\tTurn\tX occupancy\tO occupancy\tVacant\n")

	for _, x := range utils.NextMoves(start) {
		model.PrintState(x)
	}

}
