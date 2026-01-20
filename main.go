package main

import (
	"fmt"

	"github.com/utkarsh-naman/3T/src/model"
)

func main() {
	var gamemap, _ = model.LoadGMap("bin/graph/valued/map2.ttt")
	model.PrintGMap(gamemap)
	fmt.Println(len(gamemap))
}
