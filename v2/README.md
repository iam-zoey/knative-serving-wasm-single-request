# v2
Once HTTP server receives data from client, Wasm Module get the data using scanner and print.  


### HTTP server with WASM module in container 
```
 docker run --rm -it -p 8080:8080 hong0331/server-w-wasm-v2
```



```shell 
# TESTING: GET request 
curl -X POST -d "Hello from curl" http://localhost:8080/

# TESTING: POST request 
curl "http://localhost:8080/?input=Hello%20from%20GET"

```



----
## HTTP server with wasm module in local environment 
Before running this command, make sure wasmtime is installed 
- Document: https://docs.wasmtime.dev/cli-install.html

### Compile from Go into WASM 
```
GOOS=wasip1 GOARCH=wasm go build -o main.wasm module.go
```

### Test the wasm module 
```
wasmtime run main.wasm
``` 
 


