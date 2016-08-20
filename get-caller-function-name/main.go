package main

import (
	"runtime"
	"fmt"
	"log"
)

// http://stackoverflow.com/a/25927915/2979435
// http://stackoverflow.com/a/7053457/2979435
// http://stackoverflow.com/a/7053871/2979435

func main() {

	log.Println("file trace")
	Debug("")
	log.Println("caller function")
	tracex()
}


// Debug prints a debug information to the log with file and line.
func Debug(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)

	log.Printf("[cgl] debug %s:%d %v", file, line, info)

}

func tracex() {
	pc := make([]uintptr, 10)  // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	log.Printf("%s:%d %s\n", file, line, f.Name())
}