package main

//
//import (
//	"fmt"
//	"slices"
//
//	"github.com/utkarsh-naman/3T/src/model"
//	"github.com/utkarsh-naman/3T/src/utils"
//)
//
//var gamemap1, _ = model.LoadGMap("bin/graph/valued/map1.ttt")
//
//func main() {
//
//	for stateKey := range gamemap1 {
//		if isEqual(parentFromMap(stateKey), utils.PrevMoves(stateKey)) {
//			fmt.Println("correct")
//		} else {
//			fmt.Println("incorrect prevmoves func")
//			for _, a := range utils.PrevMoves(stateKey) {
//				fmt.Println("prevmoves: ", a)
//			}
//			for _, b := range parentFromMap(stateKey) {
//				fmt.Println("from Func: ", b)
//			}
//		}
//
//	}
//}
//
//func parentFromMap(state model.State) []model.State {
//	var parentStates []model.State
//	for stateKey := range gamemap1 {
//		if slices.Contains(gamemap1[stateKey].NextStates, state) {
//			parentStates = append(parentStates, stateKey)
//		}
//	}
//	return parentStates
//}
//
//func isEqual(fromMap, fromFunc []model.State) bool {
//	slices.Sort(fromMap)
//	slices.Sort(fromFunc)
//	if len(fromMap) != len(fromFunc) {
//		return false
//	}
//	for i := 0; i < len(fromMap); i++ {
//		if fromMap[i] != fromFunc[i] {
//			return false
//		}
//	}
//	return true
//}
