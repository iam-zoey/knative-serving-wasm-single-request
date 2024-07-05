# V1
Wasm module is designed to print the received input. When the server receive data, it pass the data to wasm module. 


### HTTP server with wasm module in container
```
docker run --rm -it -p 8080:8080 hong0331/server-w-wasm-v1
```

----
## HTTP server with wasm module in local environment 
Before running this command, make sure wasmtime is installed 
- Document: https://docs.wasmtime.dev/cli-install.html


### Compile from Go to Wasm 
```
GOOS=wasip1 GOARCH=wasm go build -o main.wasm module.go
```

### Run wasm module with server 
```
go run main.go 
```

```shell 
# TESTING: GET request 
curl -X POST -d "Hello from curl" http://localhost:8080/

# TESTING: POST request 
curl "http://localhost:8080/?input=Hello%20from%20GET"

```

