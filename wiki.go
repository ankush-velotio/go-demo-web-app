package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	title string
	body []byte
}

func (p *Page) save() error {
	filename := p.title + ".txt"
	return os.WriteFile(filename, p.body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title: title, body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		msg := "Page Not Found 404"
		fmt.Fprintf(w, "<h1>%s</h1>", msg)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.title, p.body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
