
Once HTTP server receives data from client, it pass it into Wasm Module. 
Then, it scans the passed data and print back. 

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

### Test go module + http server 
Uncomment line 35 in `main.go` and line 26-28 in `module.go` 


### ⚠️  Run wasm module with server 
```
go run main.go
curl -X POST http://localhost:8080 -d "hi" 
``` 
Currently, wasm module is not working ; it does not receive the data. 