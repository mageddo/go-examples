package config

import (
	"github.com/mageddo/go-examples/gowiki-vestigo-dao/webapp/dao/wiki"
	"net/http"
	"fmt"
	"html/template"
)

var templates = loadTemplates()

func RenderTemplate(w http.ResponseWriter, tmpl string, p *wiki.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadTemplates() *template.Template {
	defer func(){
		if str := recover(); str != nil {
			str = fmt.Sprintf("Not found templates %s", str)
			fmt.Println("error: ", str)
			panic(str)
		}
	}()

	var templates = template.Must(template.ParseFiles("view/edit.html", "view/view.html"))
	return templates;
}
