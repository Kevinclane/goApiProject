package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	pc := newProductsController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	http.Handle("/products", *pc)
	http.Handle("/products/", *pc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
