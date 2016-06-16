package controller

import (
	"net/http"
	"github.com/mageddo/go-examples/gowiki-vestigo-3rd-sql-dao/webapp/config"
	"github.com/mageddo/go-examples/gowiki-vestigo-3rd-sql-dao/webapp/dao/wiki"
	"github.com/mageddo/vestigo"
	"log"
)

func init(){
	log.Println("loading WikiEditController")
	App.Get("/edit/:title", func (w http.ResponseWriter, r *http.Request){
		var title string = vestigo.Param(r, "title")
		log.Println("M=wiki-edit, msg=starting")
		p, err := wiki.LoadPage(title)
		log.Println("M=wiki-edit, msg=loaded")
		if err != nil || p.Title == "" {
			p = &wiki.Page{Title: title}
		}
		config.RenderTemplate(w, "edit", p)
	})
}