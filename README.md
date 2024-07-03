# server-w-wasm
The WebAssembly modules in versions v1 and v2 are designed to print the input they receive. However, the methods for receiving the data differ between the two versions.

## v1 
In this version, the server executes the WebAssembly module and passes the data as an argument.

## v2 
Instead of passing the data directly into the WebAssembly module (i.e., injecting the data), this version allows the WebAssembly module to grab the input and automatically print it.