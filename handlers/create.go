package handlers

import (
	"net/http"

	"github.com/alwarmra/golang-webserver/data"
)

func (p *Products) CreateProduct(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProducts(prod)
	p.l.Printf("Prod: %#v", prod)
}
