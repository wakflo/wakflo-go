package main

import (
	"fmt"
	"github.com/wakflo/wakflo-go/sdk/v1"
)

// MainExecute .
func MainExecute(ptr uint32) {
	// Yay!!! you can start hacking
	sdk.ExecutePlugin(func(ctx sdk.TaskContext) sdk.TaskResult {
		fmt.Println("Hello, World!")
		return sdk.TaskResult{}
	})(ptr)
}

// main .
func main() {}
