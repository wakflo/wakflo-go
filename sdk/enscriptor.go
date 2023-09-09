package sdk

import (
	"fmt"
	"github.com/wakflo/wakflo-go/internal"
)

func ExecutePlugin(fn PluginFn) func(ptr *int32) *int32 {
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
			fmt.Println("*****************", err)
			panic(err)
		}

		fmt.Println("*****************")
		fmt.Println(string(outputBytes[:]))
		fmt.Println("*****************")

		return internal.BytesToPointer(outputBytes)
	}
}
