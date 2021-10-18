package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/samirgadkari/learn/web/REST/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	data, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Error marshalling JSON value", http.StatusInternalServerError)
	}

	rw.Write(data)
}
