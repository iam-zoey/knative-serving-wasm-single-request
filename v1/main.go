package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func runWasm(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("RunWasm funciton is called")
	var input string

	switch r.Method {
	case http.MethodGet:
		// Read input from query parameter for GET requests
		input = r.URL.Query().Get("input")
	case http.MethodPost:
		// Read input from request body for POST requests
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		input = string(body)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if input == "" {
		http.Error(w, "Error: 'input' parameter is required", http.StatusBadRequest)
		return
	}
	// Run Wasmtime command with the input string as an argument.
	cmd := exec.Command("wasmtime", "main.wasm", input)

	// Execute the Wasmtime command and read the output.
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(w, "Error running Wasmtime: %v", err)
		return
	}

	// Return the output to the client.
	fmt.Fprintf(w, "==== V1: Output from wasm module:\n%s", output)
}

func main() {
	http.HandleFunc("/", runWasm)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
