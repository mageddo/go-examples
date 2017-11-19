// https://github.com/mageddo/go-static-linking
package main

// #cgo CFLAGS: -I${SRCDIR}/lib/
// #cgo LDFLAGS: ${SRCDIR}/lib/libcounter.so
// #include <counter.h>
import "C"

func main() {
	println(C.incrementAndGet())
	println(C.incrementAndGet())
}
