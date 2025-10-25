# GasmEffects
GasmEffects is a wasm project for image manipulation.
It uses a Go backend, compiled to a WebAssembly module,
to handle all the image processing logic.
This will be used in my [website](https://noah-ruben.de)!
## Description

G(o-W)asmEffects is a Go library that it is used to manipulate images and apply different effects to them.
All with the goal of using wasm to do the image manipulation fast in the browser.

The project is split into 2 parts:
- A Go library that can be used to manipulate images
- A "runtime" that uses the go library to manipulate Images in the browser

## Development Usage

0. Requirement:
    - Go 1.25+
    - Make
1. Build
    ```bash
    make 
    ```
2. Start a http server for this folder
   ```bash
   python -m http.server -d "."
   ```
3. Open the browser and go to http://localhost:8080/

## Effects
- Grayscale
- ASCII ART
  - "Normal"
  - With Edge Detection
  - With Colors
  - Into pixel art (like a shader, see deadcells)
- As an hilbert curve
  - With colors
  - Greyscale with thickness as greyness

## JS Runtime
__TBD__


## HOTLIST
- [ ] Add Tests
- [ ] Add CI
- [ ] Create package
  - This means a bundle of files that can be dropped into any project and be used with a simple import.
- [ ] Finish Effects
- [ ] Finish JS Runtime
- [ ] Restructure project