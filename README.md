# Per-Request for Wasmtime
This project explores the integration of WebAssembly (WASM) within Knative Serving. In this approach, the HTTP server handles incoming requests by launching a new Wasmtime instance for each request. The server forwards the received data to the WASM module, which processes it using Scanner to read the input and Print to output the results. While this approach is straightforward, it may cause performance overhead due to the repeated initialization of Wasmtime for every request. 


## Project Structure 
```
├── Dockerfile           # Container build instructions
├── README.md            
├── go.mod              
├── go.sum               
├── main.go              # HTTP server implementation in Go
├── service.yaml         # Knative service configuration
└── wasm
    ├── main.wasm        # Compiled WASM module
    └── module.go        # Go source code for the WASM module


```
## Prerequisites
- [Wasmtime installed](https://docs.wasmtime.dev/cli-install.html) 
- [Docker](https://docs.docker.com/engine/install/) or [Knative Cluster](https://knative.dev/docs/getting-started/quickstart-install/) set up
- Go (v 1.21+)


## Compile Go into WebAssembly (WASM) 
To compile the Go code into a WASM binary (`.wasm`):
```
GOOS=wasip1 GOARCH=wasm go build -o wasmmain.wasm wasm/module.go
```


 
## Build and Push an image 
Build the Docker image and push it to your container registry:
```
docker build . -t <REGISTRY/IMAGE_NAME>
docker push <REGISTRY/IMAGE_NAME>

```

### Testing
Run the container: 
```
 docker run --rm -p 8080:8080 <REGISTRY/IMAGE_NAME>
```
Test the WASM container with GET and POST request:
```shell 
# TESTING: GET request 
curl -X POST -d "Hello from curl" http://localhost:8080/

# TESTING: POST request 
curl "http://localhost:8080/?input=Hello%20from%20GET"

```
To terminate the container: 
```
docker rm -f $(docker ps -l -q) 
```

---


## Run with Knative service 
### Edit`service.yaml`
Follow the `CONFIGUREME`tag and provide the name of your Docker image. 

### Apply the Knative service configuration
After updating the service.yaml with your Docker image name, apply the configuration to your Knative cluster:
```
kubectl apply -f service.yaml
```
Once applied, you can check the service URL by running:
```
kubectl get kservice
```

## Testing 
```shell
# TESTING: POST  request
curl -X POST -d "Hi WebAssembly" <SERVICE-URL>

# TESTING: GET request
curl "<SERVICE-URL>/?input=Hello%20from%20GET"
```

--- 
For more details on the work, refer my blog, [GSoC Journey: Knative Meets WASM](https://iam-zoey.notion.site/GSoC-Journey-Knative-Meets-WASM-94e5db450f944059bf648474e8f69b5d?pvs=4)

