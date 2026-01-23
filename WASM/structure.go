package main

import (
	"bytes"
	"encoding/gob"
)

type State uint32

func Extract(state State) (Continuum, Tie, Won, Lost, Turn, X, O, Vacancy State) {
	const filter9 State = (1 << 9) - 1

	Continuum = (state >> 31) & 1
	Tie = (state >> 30) & 1
	Won = (state >> 29) & 1
	Lost = (state >> 28) & 1
	Turn = (state >> 27) & 1
	X = (state >> 18) & filter9
	O = (state >> 9) & filter9
	Vacancy = state & filter9
	return
}

type StateProps struct {
	Score      float32
	WinDepth   int8
	LoseDepth  int8
	NextStates []State
}

type GMap map[State]StateProps

//func LoadGMap(filePath string) (GMap, error) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		return nil, err
//	}
//	defer func() {
//		if err := file.Close(); err != nil {
//			fmt.Println("warning: failed to close file:", err)
//		}
//	}()
//
//	decoder := gob.NewDecoder(file)
//
//	var gmap GMap
//	if err := decoder.Decode(&gmap); err != nil {
//		return nil, err
//	}
//
//	return gmap, nil
//}

func LoadGMapFromBytes(data []byte) GMap {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)

	var gmap GMap
	if err := decoder.Decode(&gmap); err != nil {
		panic(err)
	}
	return gmap
}
