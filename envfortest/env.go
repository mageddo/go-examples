package envfortest

import (
	"os"
	"strings"
	"fmt"
	"flag"
)

func init(){

	if flag.Lookup("test.v") == nil {
		fmt.Println("normal run")
	} else {
		fmt.Println("run under go test")
	}

	switch strings.ToUpper(os.Getenv("MG_ENV")) {
	case "TEST":
		fmt.Println("test mode")
		break
	case "DEV":
		fmt.Println("dev mode")
		break
	}
}


func GetMode() string {
	return strings.ToUpper(os.Getenv("MG_ENV"))
}