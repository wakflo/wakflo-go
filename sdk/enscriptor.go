package sdk

import (
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
	"unsafe"
)

func ExecutePlugin(fn PluginFn) func(ptr uint32) {
	size := uintptr(4)

	return func(ptr uint32) {
		var b []byte
		s := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		s.Len = int(size)
		s.Data = uintptr(ptr)

		var ctx TaskContext
		err := msgpack.Unmarshal(b, &ctx)
		if err != nil {
			panic(err)
		}

		fn(ctx)
	}
}
