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



WASM
```bash
cd WASM
$env:GOOS = "js"; $env:GOARCH = "wasm"; go build -o main.wasm
```

Test the engine on:
```bash
start index.html
```