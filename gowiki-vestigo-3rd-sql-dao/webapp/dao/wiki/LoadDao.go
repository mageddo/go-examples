package wiki

import (
	"github.com/mageddo/go-examples/gowiki-vestigo-3rd-sql-dao/webapp/crud"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {

	defer func(){
		msg := recover()
		if msg != nil{
			log.Println("error at select wiki", msg)
			return error("select wiki fails")
		}

	}()

	log.Println("m=LoadPage,msg=starting")
	db := crud.GetConnection()
	var u *Page
	err := db.Get(&u, "SELECT description FROM wiki WHERE name=$1", title)
	if err != nil {
		log.Println("m=LoadPage,msg=error at select query", err)
		return u, err
	}
	u.Title = title
	return u, nil
}

func (p *Page) Save() error {

	defer func(){
		msg := recover()
		if msg != nil{
			log.Println("error at insert wiki", msg)
			return error("insert wiki fails")
		}

	}()

	log.Print("m=Save,msg=starting wiki insert")
	db := crud.GetConnection()
	db.MustExec(`
			INSERT INTO wiki (name,description)
			VALUES
				($1, $2) RETURNING name;
		`, p.Title, string(p.Body))

	return nil
}
