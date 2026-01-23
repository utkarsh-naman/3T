package main

import "math"

const winRow1Mask State = 0b111000000
const winRow2Mask State = 0b000111000
const winRow3Mask State = 0b000000111
const winCol1Mask State = 0b100100100
const winCol2Mask State = 0b010010010
const winCol3Mask State = 0b001001001
const winDia1Mask State = 0b100010001
const winDia2Mask State = 0b001010100

var WinMasks = [8]State{winRow1Mask, winRow2Mask, winRow3Mask, winCol1Mask, winCol2Mask, winCol3Mask, winDia1Mask, winDia2Mask}

var POSINF = float32(math.Inf(1)) // +∞
var NEGINF = float32(math.Inf(-1))
