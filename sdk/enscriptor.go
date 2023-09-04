package sdk

import (
	"fmt"

	"unsafe"
)

func ExecutePlugin(fn PluginFn) func(ptr int32, size *int32) int32 {
	return func(ptr int32, size *int32) int32 {
		arr := unsafe.Slice(size, ptr)
		fmt.Println("Helloo")
		fmt.Println(arr)
		var ctx TaskContext
		//err := msgpack.Unmarshal(b, &ctx)
		//if err != nil {
		//	panic(err)
		//}

		fn(ctx)
		return 0
	}
}
