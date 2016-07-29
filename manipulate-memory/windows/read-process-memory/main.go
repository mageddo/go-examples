package main

import (
	"syscall"
	"fmt"
	"unsafe"
)
/**
 * This program reads the integer value of specified program address
 */
func main() {
	programPID := uint32(1344) // use Task Manager to get wanted PID
	address := uintptr(0xC08200A2E0) // program variable address to get the value as int (use cheat engine to get a address)

	handle, error := syscall.OpenProcess(uint32(0x0010), true, programPID)
	fmt.Println("handler=", handle, "error=", error)
	var jx int = 1;
	var value *int = &jx;
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	readProcessMemoryAddress,_ := syscall.GetProcAddress(kernel32, "ReadProcessMemory")
	r1, r2, err := syscall.Syscall6(
		(readProcessMemoryAddress), 5, uintptr(handle),
		address, uintptr(unsafe.Pointer(value)), unsafe.Sizeof(8), 0, 0,
	)

	fmt.Println("r1=", r1, "r2=", r2, "err=", err, "value=", *value)

}
