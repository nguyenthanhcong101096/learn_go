package data

import (
	"encoding/json"
	"io"
	"log"
	"os"
	loader "restful_api/loader"
)

type Books struct {
	Store []*loader.Book `json:"store"`
}

type Store []*loader.Book

func (b *Books) Initialize() {
	filename := "./assets/books.csv"
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b.Store = loader.LoadData(file)
}

func (b *Books) All() Store {
	return b.Store
}

func (b *Store) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}
