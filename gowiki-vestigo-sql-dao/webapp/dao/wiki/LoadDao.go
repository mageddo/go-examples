package wiki

import (
	"io/ioutil"
	"fmt"
	"github.com/mageddo/go-examples/gowiki-vestigo-sql-dao/webapp/crud"
	"database/sql"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {
	filename := getFilename(title)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func (p *Page) Save() error {
	log.Print("m=Save,msg=starting wiki insert")
	qtd, err := crud.Run(func(db *sql.DB) (interface{}, error) {

		row := db.QueryRow(`
			INSERT INTO wiki (name,description)
			VALUES
				($1, $2);
		`, p.Title, string(p.Body))

		var qtd int64
		err := row.Scan(qtd)
		return qtd, err
	})
	log.Println(qtd, "inserted wiki(s)")
	return err
}

func getFilename(name string) string {
	return fmt.Sprintf("base/%s.txt", name)
}