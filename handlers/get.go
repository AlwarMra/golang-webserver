package handlers

import (
	"net/http"

	"github.com/alwarmra/golang-webserver/data"
)

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal products", http.StatusInternalServerError)
	}
}
