package controller

import (
	"net/http"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/config"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/dao/wiki"
	"log"
)

func init(){
	log.Println("loading WikiEditController")
	http.HandleFunc("/edit/", config.MakeHandler(func (w http.ResponseWriter, r *http.Request, title string){
		p, err := wiki.LoadPage(title)
		if err != nil {
			p = &wiki.Page{Title: title}
		}
		config.RenderTemplate(w, "edit", p)
	}))
}