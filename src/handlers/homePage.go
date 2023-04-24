package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	l *log.Logger
}

func NewHomePage(l *log.Logger) *HomePage {
	return &HomePage{l}
}

func (h *HomePage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Home page accessed!")

	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseGlob("homePage.html"))
	myTemplate.ExecuteTemplate(rw, "homePage.html", nil)
}
