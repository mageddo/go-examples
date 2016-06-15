package controller

import (
	"net/http"
	"github.com/mageddo/go-examples/gowiki-vestigo-sql-dao/webapp/config"
	"github.com/mageddo/go-examples/gowiki-vestigo-sql-dao/webapp/dao/wiki"
	"log"
	"github.com/mageddo/vestigo"
)

func init(){
	log.Println("loading WikiViewController")
	App.Get("/view/:title", func (w http.ResponseWriter, r *http.Request){
		title := vestigo.Param(r, "title")
		log.Println("M=wiki-view, msg=loading")
		p, err := wiki.LoadPage(title)
		log.Println("M=wiki-view, msg=loaded")
		if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
		}
		config.RenderTemplate(w, "view", p)
	})
}