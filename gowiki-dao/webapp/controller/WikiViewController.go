package controller

import (
	"net/http"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/config"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/dao/wiki"
	"log"
)

func init(){
	log.Println("loading WikiViewController")
	http.HandleFunc("/view/", config.MakeHandler(func (w http.ResponseWriter, r *http.Request, title string){
		p, err := wiki.LoadPage(title)
		if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
		}
		config.RenderTemplate(w, "view", p)
	}))
}