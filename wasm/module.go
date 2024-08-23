package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Read input line by line
		input := scanner.Text()
		fmt.Println("========== WASM MODULE=============")
		fmt.Println("WASM says: ", input)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
