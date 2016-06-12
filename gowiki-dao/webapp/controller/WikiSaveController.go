package controller

import (
	"net/http"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/config"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/dao/wiki"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/req"
	"log"
)

func init(){

	log.Println("loading WikiSaveController")

	http.HandleFunc(req.Paths.Load("/save/"), config.MakeHandler(func (w http.ResponseWriter, r *http.Request, title string) {
		body := r.FormValue("body")
		p := &wiki.Page{Title: title, Body: []byte(body)}
		err := p.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}))
}