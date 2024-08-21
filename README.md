
Once HTTP server receives data from client, Wasm Module get the data using scanner and print.  


### Running w/ knative service 
```
kubectl apply -f service.yaml
```

To check on your domain, 
```
kubectl get kservice
```

```shell
# TESTING: POST  request
curl -X POST -d <your domain>

# TESTING: GET request
curl "<your-domain>/?input=Hello%20from%20GET"
```



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

### Test the wasm module locally 
```
wasmtime run main.wasm
``` 
 
### Build and Push an image 
```
docker build -t . <registry/image-name>
docker push <registry/image-name>
```


