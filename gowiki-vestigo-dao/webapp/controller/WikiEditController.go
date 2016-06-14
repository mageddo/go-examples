package controller

import (
	"net/http"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/config"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/dao/wiki"
	"github.com/mageddo/vestigo"
	"log"
)

func init(){
	log.Println("loading WikiEditController")
	App.Get("/edit/:title", func (w http.ResponseWriter, r *http.Request){
		var title string = vestigo.Param(r, "title")
		p, err := wiki.LoadPage(title)
		if err != nil {
			p = &wiki.Page{Title: title}
		}
		config.RenderTemplate(w, "edit", p)
	})
}