package main

import (
	"syscall"
	"fmt"
	"unsafe"
)

func main() {
	handle, error := syscall.OpenProcess(uint32(0x0010), true, 1344)
	fmt.Println("handler=", handle, "error=", error)
	var jx int = 1;
	var value *int = &jx;
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	readProcessMemoryAddress,_ := syscall.GetProcAddress(kernel32, "ReadProcessMemory")
	r1, r2, err := syscall.Syscall6(
		(readProcessMemoryAddress), 5, uintptr(handle),
		uintptr(0xC08200A2E0), uintptr(unsafe.Pointer(value)), unsafe.Sizeof(8), 0, 0,
	)

	fmt.Println("r1=", r1, "r2=", r2, "err=", err, "value=", *value)

}
