package internal

import (
	"strings"
	"unsafe"
)

// PointerToBytes returns a string from WebAssembly compatible numeric types
// representing its pointer and length.
func PointerToBytes(subject *int32) []byte {
	nth := 4
	var subjectStr strings.Builder
	var buf []byte
	pointer := uintptr(unsafe.Pointer(subject))
	for {
		s := *(*int32)(unsafe.Pointer(pointer + uintptr(nth)))
		if s == 0 {
			break
		}

		buf = append(buf, byte(s))
		subjectStr.WriteByte(byte(s))
		nth++
	}

	return buf
}

// BytesToPointer returns a string from WebAssembly compatible numeric types
// representing its pointer and length.
func BytesToPointer(message []byte) *int32 {
	r := make([]int32, 2)
	r[0] = int32(uintptr(unsafe.Pointer(&(message[0]))))
	r[1] = int32(len(message))

	return &r[0]
}
