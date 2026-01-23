package main

import (
	"fmt"

	"github.com/utkarsh-naman/3T/src/model"
)

func main() {
	gmap, _ := model.LoadGMap("bin/graph/valued/map2.ttt")
	printHardCodedGMap(gmap)
	fmt.Println(len(gmap))
}
func printHardCodedGMap(gmap model.GMap) {
	fmt.Println("var tttmap = model.GMap{")
	for key, value := range gmap {
		fmt.Println("\t", key, ": {")
		fmt.Println("\t\tScore: ", value.Score, ",")
		fmt.Println("\t\tWinDepth: ", value.WinDepth, ",")
		fmt.Println("\t\tLoseDepth: ", value.LoseDepth, ",")
		fmt.Print("\t\tNextStates: []model.State{")
		for _, nextState := range value.NextStates {
			fmt.Print(nextState, ",")
		}
		fmt.Println("},")
		fmt.Println("\t},")
	}
	fmt.Println("}")
}
