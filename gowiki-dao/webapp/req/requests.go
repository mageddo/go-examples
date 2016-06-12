package req

import (
	"container/list"
	"log"
)

type Path struct {
	Paths *list.List
}

var Paths *Path

func init(){
	Paths = new(Path)
	Paths.Paths = list.New()
}

func (p *Path) Load(str string) string {
	log.Println("mapping to: ", str)
	p.Paths.PushBack(str)
	return str
}

//func (p *Path) GetPaths() list.List {
//	return
//}