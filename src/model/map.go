package model

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type StateProps struct {
	Score      float32
	WinDepth   int8
	LoseDepth  int8
	NextStates []State
}

type GMap map[State]StateProps

// use m := make(utils.GMap)

func PrintGMap(data GMap) {
	if len(data) == 0 {
		fmt.Println("GMap is empty.")
		return
	}

	fmt.Println("--- GMap Contents ---")

	// 1. Extract Keys to ensure deterministic order
	keys := make([]State, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	// 2. Sort the keys (State is uint32, so we verify order)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// 3. Iterate over sorted keys and print
	for _, stateID := range keys {
		props := data[stateID]

		// Header for the State
		fmt.Printf("State [%d]:\n", stateID)

		// Body (Props) indented for readability
		fmt.Printf("  Score:      %.2f\n", props.Score)
		fmt.Printf("  WinDepth:   %d\n", props.WinDepth)
		fmt.Printf("  LoseDepth:  %d\n", props.LoseDepth)

		// Format NextStates specially
		if len(props.NextStates) > 0 {
			fmt.Printf("  NextStates: %v\n", props.NextStates)
		} else {
			fmt.Println("  NextStates: (None)")
		}

		// Add a separator for visual clarity
		fmt.Println("- - - - - -")
	}
	fmt.Println("--- End of Map ---")
}

// --- Serialization Functions ---
// SaveGMap serializes the game map into a .ttt binary file

func SaveGMap(gmap GMap, filePath string) error {
	// Ensure directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("warning: failed to close file:", err)
		}
	}()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(gmap); err != nil {
		return err
	}

	return nil
}

func LoadGMap(filePath string) (GMap, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("warning: failed to close file:", err)
		}
	}()

	decoder := gob.NewDecoder(file)

	var gmap GMap
	if err := decoder.Decode(&gmap); err != nil {
		return nil, err
	}

	return gmap, nil
}
