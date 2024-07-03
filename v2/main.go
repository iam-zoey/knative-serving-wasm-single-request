package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {

	// Handle POST requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Read the entire request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Prepare to capture stdout and stderr from wasmtime
		var stdout, stderr bytes.Buffer

		// Execute the WebAssembly module that wait for user's input, read and print it from stdin
		cmd := exec.Command("wasmtime", "main.wasm")
		// cmd := exec.Command("go", "run", "module.go") // Run the Go module instead of the WebAssembly module for testing

		// Set the input from HTTP request body as stdin for the wasm command
		cmd.Stdin = bytes.NewReader(body)

		// Capture stdout and stderr from wasmtime command
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// Run the command
		err = cmd.Run()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to execute WebAssembly module: %v\nStderr: %s", err, stderr.String()), http.StatusInternalServerError)
			return
		}

		// Write the output of the WebAssembly module execution to the response
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Println(stdout.String())

	})

	// Start the HTTP server
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
