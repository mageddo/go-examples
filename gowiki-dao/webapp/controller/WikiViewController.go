package controller

import "net/http"
import "github.com/mageddo/go-examples/gowiki-dao/webapp/config"

func main(){
	http.HandleFunc("/view/", config.MakeHandler(func (w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage(title)
	if err != nil {
	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	return
	}
	config.RenderTemplate(w, "view", p)
	}))

}