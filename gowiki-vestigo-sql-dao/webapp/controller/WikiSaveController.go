package controller

import (
	"net/http"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/dao/wiki"
	"github.com/mageddo/vestigo"
	"log"
)

func init(){

	log.Println("loading WikiSaveController")

	App.Post("/save/:title", func (w http.ResponseWriter, r *http.Request) {
		title := vestigo.Param(r, "title")
		body := r.FormValue("body")
		p := &wiki.Page{Title: title, Body: []byte(body)}
		err := p.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	})
}