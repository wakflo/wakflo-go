// Sample demo go plugin
package main

import (
	"github.com/wakflo/wakflo-go/sdk"
)

// PluginFn write your plugin here
func PluginFn(_ sdk.TaskContext) sdk.TaskResult {
	msg := "Hello, World!"

	// update output
	return sdk.TaskResult{
		Output: map[string]interface{}{
			"message": msg,
		},
		Errors: make([]sdk.SystemActivityLog, 0),
	}
}

//--------------------------------------DO NOT MODIFY BELOW ------------------------------------//

// DO NOT MODIFY
//
//go:wasm-module execute
//go:export main
func execute(ptr *int32) *int32 {
	return sdk.ExecutePlugin(PluginFn)(ptr)
}

// DO NOT MODIFY
// main .
func main() {}

//--------------------------------------DO NOT MODIFY ABOVE ------------------------------------//
