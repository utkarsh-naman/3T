package model

import (
	"fmt"
)

type State uint32

// state representation in bit form
// Continuum	Tie,		Won Index,	Lost Index,		Turn,		X,			O,			Vacancy
// (1 bit)		(1 bit)		(1 bit)		(1 bit)			(1 bit)		(9 bits)	(9 bits)	(9 bits)

func Extract(state State) (Continuum, Tie, Won, Lost, Turn, X, O, Vacancy State) {
	const filter9 State = (1 << 9) - 1

	Continuum = (state >> 31) & 1 // 1st bit from left to right
	Tie = (state >> 30) & 1       // 2nd bit from left to right
	Won = (state >> 29) & 1       // 3rd bit from left to right
	Lost = (state >> 28) & 1      // 4th bit from left to right
	Turn = (state >> 27) & 1      // 5th bit from left to right
	X = (state >> 18) & filter9   // 6 to 14 bit from left to right
	O = (state >> 9) & filter9    // 15 to 23 bit from left to right
	Vacancy = state & filter9     // 24 to 32 bit from left to right
	return
}

func PrintState(state State) {
	c, t, w, l, tu, x, o, v := Extract(state)
	fmt.Printf("%01b\t%01b\t%01b\t%01b\t%01b\t%09b\t%09b\t%09b", c, t, w, l, tu, x, o, v)
	fmt.Println("")
}
