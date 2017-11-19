package main

import "syscall"

func main() {
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
}
