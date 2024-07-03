
Once HTTP server receives data from client, Wasm Module get the data using scanner and print.  

Before running this command, make sure wasmtime is installed 
- Document: https://docs.wasmtime.dev/cli-install.html

### Compile from Go into WASM 
```shell 
GOOS=wasip1 GOARCH=wasm go build -o main.wasm module.go
```

### Test the wasm module 
```shell 
wasmtime run main.wasm
``` 
 


###  Run wasm module with server 
```
go run main.go
curl -X POST http://localhost:8080 -d "hi" 
``` 
