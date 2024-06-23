# Keystroke Rainbow visualizer

A simple colorful line art generator that reacts to keystrokes, written in Go.

## WASM
This game can be played in the browser with webassembly. I have a copy of the Go wasm file in `/vendor/`, to download a fresh copy run 

```bash
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./vendor/
```

Compile the WASM binary with 
```bash
env GOOS=js GOARCH=wasm go build -o lineart.wasm github.com/simongle/go-art
```

serve with this (or your favorite local server option)

```bash
python3 -m http.server
```