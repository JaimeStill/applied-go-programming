package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== Go Program Structure Demo ===")
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("CPU Count: %d\n", runtime.NumCPU())
}
