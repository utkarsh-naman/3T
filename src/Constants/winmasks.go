package Constants

import (
	"github.com/utkarsh-naman/3T/src/model"
)

const winRow1Mask model.State = 0b111000000
const winRow2Mask model.State = 0b000111000
const winRow3Mask model.State = 0b000000111
const winCol1Mask model.State = 0b100100100
const winCol2Mask model.State = 0b010010010
const winCol3Mask model.State = 0b001001001
const winDia1Mask model.State = 0b100010001
const winDia2Mask model.State = 0b001010100

var WinMasks = [8]model.State{winRow1Mask, winRow2Mask, winRow3Mask, winCol1Mask, winCol2Mask, winCol3Mask, winDia1Mask, winDia2Mask}
