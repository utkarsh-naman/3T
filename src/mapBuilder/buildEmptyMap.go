package main

import (
	"fmt"

	"github.com/utkarsh-naman/3T/src/model"
	"github.com/utkarsh-naman/3T/src/utils"
)

func genMap(start model.State) model.GMap {
	gamemap := make(model.GMap)

	//bfs create map
	visited := map[model.State]bool{}
	visited[start] = true

	var toVisitQueue queue
	toVisitQueue.enqueue(start)

	for len(toVisitQueue) != 0 {
		var stateNode, _ = toVisitQueue.dequeue()
		var nodeprops model.StateProps
		nodeprops.Score = 0.0
		nodeprops.WinDepth = 10
		nodeprops.LoseDepth = 10
		var Children = utils.NextMoves(stateNode)
		nodeprops.NextStates = Children

		gamemap[stateNode] = nodeprops

		for _, childNode := range Children {
			if !visited[childNode] {
				visited[childNode] = true
				toVisitQueue.enqueue(childNode)
			}
		}
	}
	fmt.Println("Finished building map")
	return gamemap

}

func main() {
	const start model.State = 0b10000000000000000000000111111111
	result := genMap(start)
	fmt.Println("length of the map:", len(result))
	model.PrintGMap(result)
	err := model.SaveGMap(result, "bin/graph/unvalued/map0.ttt")
	if err != nil {
		return
	}
}

type queue []model.State

func (q *queue) enqueue(x model.State) {
	*q = append(*q, x)
}

func (q *queue) dequeue() (model.State, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x, true
}

//var q Queue
//
//q.Enqueue(10)
//q.Enqueue(20)
//
//v, ok := q.Dequeue() // 10, true
