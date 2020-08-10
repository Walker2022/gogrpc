package main

import (
	"fmt"
	"syscall"
	"unsafe"
)



func IntPtr(n int32) uintptr {
	return uintptr(n)
}

func WStrPtr(str string) uintptr {

	pchar, _ := syscall.UTF16PtrFromString(str)
	return uintptr(unsafe.Pointer(pchar))
}

func AStrPtr(str string) uintptr {

	var buf []byte = make([]byte, len(str)+1)
	copy(buf, str)
	return uintptr(unsafe.Pointer(&buf[0]))

}

func pChar2string(rt uintptr) string {

	p := (*byte)(unsafe.Pointer(rt))
	data := make([]byte, 0)
	for *p != 0 {
		data = append(data, *p)
		rt += unsafe.Sizeof(byte(0))
		p = (*byte)(unsafe.Pointer(rt))
	}
	rString := string(data)
	fmt.Println(rString)

	return rString
}

func newBuffer(buf []byte) uintptr {

	return uintptr(unsafe.Pointer(&buf[0]))
}

func newBufferSize( bufsize *int  )uintptr{

	return uintptr(unsafe.Pointer(bufsize))
}