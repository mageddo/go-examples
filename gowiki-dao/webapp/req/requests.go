package req

import "container/list"

type Path struct {
	Paths *list.List
}

var Paths *Path

func init(){
	Paths = new(Path)
	Paths.Paths = list.New()
}

func (p *Path) Load(str string) string {
	p.Paths.PushBack(str)
	return str
}

//func (p *Path) GetPaths() list.List {
//	return
//}