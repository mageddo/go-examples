package config

import (
	"github.com/mageddo/go-examples/gowiki-dao/webapp/dao/wiki"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/req"
	"regexp"
	"net/http"
	"fmt"
	"html/template"
	"bytes"
	"log"
)


func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	var buffer bytes.Buffer
	for e := req.Paths.Front(); e != nil; e = e.Next() {
		buffer.WriteString(e.Value.(string))
		buffer.WriteString("|")
	}
	buffer.Truncate(buffer.Len() - 1)
	regex := fmt.Sprintf("^(%s)([a-zA-Z0-9]+)$", buffer.String())
	log.Println("===> fixing: ", regex)
	var validPath = regexp.MustCompile(regex)
	//var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
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
