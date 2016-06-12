package main

import (
	"github.com/mageddo/go-examples/gowiki-dao/webapp/config"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/dao/wiki"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/controller"
	"net/http"
)

func main() {
	//controller.LoadController()
	http.HandleFunc("/edit/", config.MakeHandler(editHandler))
	http.HandleFunc("/save/", config.MakeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &wiki.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}