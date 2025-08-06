package main

import (
	"fmt"
	"os"
)

func main() {
	if err := runTUI(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
