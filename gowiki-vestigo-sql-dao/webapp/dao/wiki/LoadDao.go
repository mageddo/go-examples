package wiki

import (
	"github.com/mageddo/go-examples/gowiki-vestigo-sql-dao/webapp/crud"
	"database/sql"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {

	log.Println("m=LoadPage,msg=starting")
	o, err := crud.Run(func(db *sql.DB) (interface{}, error){
		rows, err := db.Query("SELECT description FROM wiki WHERE name=$1", title);
		var u Page
		if err != nil {
			log.Println("m=LoadPage,msg=error at select query", err)
			return u, err
		}
		defer func(){
			log.Println("m=LoadPage,msg=closing rows")
			rows.Close()
		}()
		rows.Next()
		rows.Scan(&u.Body)
		u.Title = title
		return u, nil
	})

	log.Println("m=LoadPage,msg=do casting")
	var v = o.(Page)
	log.Println("m=LoadPage,msg=casting done")
	return &v, err
}

func (p *Page) Save() error {
	log.Print("m=Save,msg=starting wiki insert")
	qtd, err := crud.Run(func(db *sql.DB) (interface{}, error) {

		row := db.QueryRow(`
			INSERT INTO wiki (name,description)
			VALUES
				($1, $2) RETURNING name;
		`, p.Title, string(p.Body))

		var name string
		err := row.Scan(&name)
		return name, err
	})
	if err == nil{
		log.Println(qtd, " wiki inserted")
	}
	return err
}
