Made by [Instruction](https://itnext.io/webassemply-with-golang-by-scratch-e05ec5230558)


## Running

#### Dependencies:

* [Go compiler](https://golang.org/)

#### Setup:

* Compiling [`wasm-code/main.go`](https://github.com/Dominux/go-projects/simple-wasm/wasm-code/main.go) file into `wasm` code:
```bash
cd wasm-code && GOOS=js GOARCH=wasm gobuild -o main.wasm && cd -
```

* Running example server
```bash
go run server.go
```

* Finally go to http://localhost.8080/wasm.html and open the DevTools console
