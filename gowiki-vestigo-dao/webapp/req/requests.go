package req

import (
	"container/list"
	"log"
)

var Paths = list.New()

func Load(str string) string {
	log.Println("mapping to: ", str)
	Paths.PushBack(str)
	return str
}

func GetPaths() *list.List {
	return Paths
}