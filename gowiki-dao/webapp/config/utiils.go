package config

import (
	"regexp"
	"net/http"
	"fmt"
	"html/template"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

var templates = loadTemplates()

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
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
