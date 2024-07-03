package main

import (
	"bufio"
	"fmt"
	"os"
)

func runwasm() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Read input line by line
		input := scanner.Text()
		fmt.Println("WASM says", input)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

/*
============================================
FOR TESTING:
Uncomment this part if you are running a go module, instead of WebAssembly
*/
// func main() {
// 	runwasm()
// }
