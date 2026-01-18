package SymmetryReduction

import (
	"github.com/utkarsh-naman/3T/src/model"
)

func CollapseEquivalents(state model.State) model.State {
	var head = state >> 27

	_, _, _, _, _, x, o, v := model.Extract(state)
	var canonical = uint32(x<<18 | o<<9 | v)

	for _, f := range [][512]uint32{
		ROT90TABLE,
		ROT180TABLE,
		ROT270TABLE,
		FLIPHTABLE,
		FLIPVTABLE,
		DIAGTABLE,
		ANTIDIAGTABLE,
	} {
		if s := f[x]<<18 | f[o]<<9 | f[v]; s < canonical {
			canonical = s
		}
	}
	return head<<27 | model.State(canonical)
}
