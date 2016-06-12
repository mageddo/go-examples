package main

import (
	"github.com/mageddo/go-examples/gowiki-dao/webapp/config"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/dao/wiki"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/edit/", config.MakeHandler(editHandler))
	http.HandleFunc("/save/", config.MakeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := wiki.LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	config.RenderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}