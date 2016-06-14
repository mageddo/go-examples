package wiki

import (
	"io/ioutil"
	"fmt"
)

type Page struct {
	Title string
	Body []byte
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
	filename := getFilename(p.Title)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func getFilename(name string) string {
	return fmt.Sprintf("base/%s.txt", name)
}