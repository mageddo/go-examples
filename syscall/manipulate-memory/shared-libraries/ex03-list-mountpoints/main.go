// https://github.com/mageddo/go-static-linking
package main

// #include <mntent.h>
// #include <stdio.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	f := C.fopen(C.CString("/etc/mtab"), C.CString("r"))
	mpoint := C.getmntent(f)

	fmt.Printf("fsname=%s, mntdir=%s, mnt_type=%s\n", C.GoString(mpoint.mnt_fsname), C.GoString(mpoint.mnt_dir), C.GoString(mpoint.mnt_type))

	fsnameBytes := (*[unsafe.Sizeof(mpoint.mnt_fsname)]byte)(unsafe.Pointer(mpoint.mnt_fsname))
	fmt.Printf("fromByteArray=%s, fromGoString=%s\n", string(fsnameBytes[:]), C.GoString(mpoint.mnt_fsname))

}
