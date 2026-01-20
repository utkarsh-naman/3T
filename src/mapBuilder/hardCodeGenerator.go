package main

import (
	"fmt"
	"sort"

	"github.com/utkarsh-naman/3T/src/model"
)

func main() {
	gmap, _ := model.LoadGMap("bin/graph/valued/map2.ttt")
	printHardCodedGMap(gmap)
	fmt.Println(len(gmap))
}
func printHardCodedGMap(gmap model.GMap) {
	fmt.Println("func LoadHardCodedGMap() GMap {")
	fmt.Println("\treturn GMap{")

	// Stable order for reproducible output
	keys := make([]model.State, 0, len(gmap))
	for k := range gmap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, state := range keys {
		props := gmap[state]

		fmt.Printf("\t\t%#032b: {\n", uint32(state))
		fmt.Printf("\t\t\tScore: %g,\n", props.Score)
		fmt.Printf("\t\t\tWinDepth: %d,\n", props.WinDepth)
		fmt.Printf("\t\t\tLoseDepth: %d,\n", props.LoseDepth)

		fmt.Printf("\t\t\tNextStates: []State{")
		for i, child := range props.NextStates {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%#032b", uint32(child))
		}
		fmt.Println("},")
		fmt.Println("\t\t},")
	}

	fmt.Println("\t}")
	fmt.Println("}")
}
