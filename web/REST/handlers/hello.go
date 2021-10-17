package handlers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Could not read request body", http.StatusBadRequest)
		return
	}

	log.Printf("Data: %s", d)
	io.WriteString(rw, "Hello world")
}
