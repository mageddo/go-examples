package wiki

import "github.com/mageddo/go-examples/gowiki-dao/webapp/model"

func LoadPage(title string) (*model.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &model.Page{Title: title, Body: body}, nil
}
