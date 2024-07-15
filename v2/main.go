package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	// Start the HTTP server
	http.HandleFunc("/", runWasm)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// runWasm handles HTTP requests (POST, GET) by invoking the Wasmtime module
func runWasm(w http.ResponseWriter, r *http.Request) {
	var input string

	switch r.Method {
	case http.MethodGet:
		// Read input from query parameter for GET requests
		input = r.URL.Query().Get("input")
	case http.MethodPost:
		// Copy input from request body for POST requests
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Convert buffer to string
		input = buf.String()
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

	// Command to execute
	cmd := exec.Command("wasmtime", "main.wasm")

	// Set the input from HTTP request as stdin for the command
	cmd.Stdin = bytes.NewBufferString(input)

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
	w.Write(stdout.Bytes())
}
