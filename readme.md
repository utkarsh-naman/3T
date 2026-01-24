<div align="center">
  <img width="50%" alt="roman men playing tic-tac-toe" src="assets/TTTLogo.png" />
</div>

# 3T : Highly resource optimized Tic-Tac-Toe Game Engine

>3T is a game engine for TIC-TAC-TOE for the most resource optimized tic-tac-toe game bot.</br>


>### For understanding the project and model structure, please refer [3T Project Documentation](./Docs/docs.md) 

## Table of Contents
- [Build Instructions](#build-instructions)
- - [Building GAME](#building-the-game-graph)
- - [Building WASM GameEngine](#wasm)
- [Testing the WASM engine](#test-the-engine-on)
- [Community Standards](#community-standards--community-health-files)


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


## Community Standards / Community Health Files
* [LICENSE](LICENSE)
* [CONTRIBUTING](CONTRIBUTING.md)
* [CODE OF CONDUCT](CODE_OF_CONDUCT.md)
* [ISSUE TEMPLATE](.github/ISSUE_TEMPLATE/issue_template.md)
* [PULL REQUEST TEMPLATE](.github/PULL_REQUEST_TEMPLATE.md)