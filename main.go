package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func runWasm(w http.ResponseWriter, r *http.Request) {
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

	// Define stdout and stderr buffers
	var stdout, stderr bytes.Buffer

	//cmd := exec.Command("wasmtime", "main.wasm")
	cmd := exec.Command("wasmtime", "wasm/main.wasm") // Run the Go module instead of the WebAssembly module for testing

	// Set the input from HTTP request body as stdin for the command
	cmd.Stdin = bytes.NewReader([]byte(input))

	// Capture stdout and stderr from the command
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute module: %v\nStderr: %s", err, stderr.String()), http.StatusInternalServerError)
		return
	}

	// Write the output of the command execution to the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, stdout.String())
}

func main() {
	// Start the HTTP server
	http.HandleFunc("/", runWasm)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
