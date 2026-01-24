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
- [Issues](#issues)


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


## Issues
### Issue Template
**Filename:** `.github/ISSUE_TEMPLATE/issue_template.md` (You may need to create the folders `.github` and `ISSUE_TEMPLATE` first).

```markdown
---
name: Bug Report or Feature Request
about: Create a report to help us improve 3T
title: ''
labels: ''
assignees: ''
---

## Description

> A clear and concise description of what the bug is or what feature you want implemented.

## Steps to Reproduce

1. Go to '...'
2. Run command '...'
3. See error
4. Or wrong output for an input by a function

## Expected Behavior

> A clear and concise description of what you expected to happen.

## Environment (if applicable/specific)

* OS: [e.g. Windows/Linux]
* Go Version: [e.g. 1.21]
* Browser (if WASM related): [e.g. Chrome, Firefox]

## Logs / Screenshots

If applicable, add logs or screenshots to help explain your problem.


