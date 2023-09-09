package sdk

import (
	"github.com/wakflo/wakflo-go/internal"
)

// ExecutePlugin wrapper function for plugin closure
func ExecutePlugin(fn PluginFn) func(ptr *int32) *int32 {
	// closure abstracts memory access convert pointers to structs
	return func(ptr *int32) *int32 {
		inputBytes := internal.PointerToBytes(ptr)
		var ctx TaskContext
		_, err := ctx.UnmarshalMsg(inputBytes)
		if err != nil {
			panic(err)
		}

		// call function
		rsp := fn(ctx)

		// process output
		outputBytes, err := rsp.MarshalMsg(nil)
		if err != nil {
			panic(err)
		}
		return internal.BytesToPointer(outputBytes)
	}
}
