package main

import _ "embed"

//go:embed map4.ttt
var mapData []byte

var tttmap = LoadGMapFromBytes(mapData)
