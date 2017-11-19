// Chamando funcoes do sistema no windows - https://github.com/golang/go/wiki/WindowsDLLs
// messageBox documentation - https://msdn.microsoft.com/en-us/library/windows/desktop/ms645505(v=vs.85).aspx
// string to wchar_t - https://stackoverflow.com/a/8032108
// int to size_t - https://groups.google.com/forum/#!topic/golang-nuts/IreLjw1b0L0
// https://github.com/mozilla/masche/blob/master/listlibs/listlibs_windows.go
// lendo wchar para string - https://groups.google.com/forum/#!topic/golang-nuts/8c18AJFM6Kc

package main
// #include <windows.h>
// #include <Winnt.h>
// #include <wchar.h>
// #include<stdio.h>
import "C"

import (
"fmt"
"unsafe"
)

type MSG_BOX_RET int

const (
	YES = 0x06
	NO = 0x07
)

func main(){

  // buf := new(C.WCHAR)
  x := make([]byte, 500)
  //x := []byte(string("joao hi!"))
  buf := (*C.WCHAR)(unsafe.Pointer(&x))

	msg := "Dois corações apaixonados"

	var msize uint64 = uint64(len(msg))
	fmt.Printf("size=%d\n", msize)
	C.mbstowcs ((*C.wchar_t)(unsafe.Pointer(buf)), C.CString(msg), *(*C.size_t)(&msize));

//	pFile := C.fopen (C.CString("myfile.txt"),C.CString("r"));
//	C.fgetws (((*C.wchar_t)(unsafe.Pointer(buf))), 100, pFile);

  fmt.Printf("%#v, %+v\n", buf, buf)
	result := C.MessageBoxW(nil, buf, buf, 0x00000004)
	fmt.Printf("status=end, result=%#v, yes=%t\n", result, result == YES)
}

/*
	https://msdn.microsoft.com/en-us/library/windows/desktop/ff381407(v=vs.85).aspx
	https://msdn.microsoft.com/en-us/library/gg269344(v=exchg.10).aspx


	wchar_t e WCHAR sao o mesmo, sao alias por isso o codigo abaixo funciona
	buf := new(C.WCHAR
	var *C.wchar_t = (*C.wchar_t)(unsafe.Pointer(buf))

*/
