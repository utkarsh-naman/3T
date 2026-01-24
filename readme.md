<div align="center">
  <img src="https://github-production-user-asset-6210df.s3.amazonaws.com/114791877/540058020-628801f6-b8de-46ae-974d-5920efc663b9.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20260124%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260124T084019Z&X-Amz-Expires=300&X-Amz-Signature=2fbbbfda0206c7991cb3b5837d0dec475eb91561edc87cdf436848edeab1cb68&X-Amz-SignedHeaders=host" width="50%" alt="roman men playing tic-tac-toe" />
</div>


>3T is a game engine for TIC-TAC-TOE for the most resource optimized tic-tac-toe game bot.</br>


>For understanding the project and model structure, please refer [Build Instructions](instructions.md) 

## Table of Contents
- [Build Instructions](#build-instructions)
- - [Building GAME](#building-the-game-graph)
- - [Building WASM GameEngine](#wasm)
- [Testing the WASM engine](#test-the-engine-on)



## Build Instructions
### Building the game graph
```bash
go run "src/mapBuilder/buidEmptyMap.go"
```
```bash
go run "src/mapBuilder/setTerminalValues.go"
```
```bash
go run "src/mapBuilder/setValues.go"
```



### WASM
```bash
cd WASM
$env:GOOS = "js"; $env:GOARCH = "wasm"; go build -o main.wasm
```

### Test the engine on:
```bash
start index.html
```


