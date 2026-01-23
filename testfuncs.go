package main

import (
	"fmt"

	"github.com/utkarsh-naman/3T/src/Constants"
	"github.com/utkarsh-naman/3T/src/model"
)

func main() {
	var gmapp, _ = model.LoadGMap("bin/graph/valued/map1.ttt")
	var count int16 = 0
	for key := range gmapp {
		if gmapp[key].Score == Constants.POSINF {
			count++
			fmt.Printf("%01b, %01b, %01b, %01b, %01b, %09b, %09b, %09b \n", (key>>31)&1, (key>>30)&1, (key>>29)&1, (key>>28)&1, (key>>27)&1, (key>>18)&0b111111111, (key>>9)&0b111111111, key&0b111111111)
		}
	}
	fmt.Println(count)
}
