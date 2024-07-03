Wasm module is designed to print the received input. When the server receive data, it pass the data to wasm module. 

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

