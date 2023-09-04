package main

import (
	"fmt"
	"github.com/wakflo/wakflo-go/sdk"
)

//go:wasm-module execute
//go:export main
func execute(ptr int32, size *int32) int32 {
	// Yay!!! you can start hacking
	return sdk.ExecutePlugin(func(ctx sdk.TaskContext) sdk.TaskResult {
		// write here
		fmt.Println("Hello, World!")
		return sdk.TaskResult{}
	})(ptr, size)
}

// main .
func main() {}
