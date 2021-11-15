package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	pc := newProductsController()
	cc := newCartsController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	http.Handle("/products", *pc)
	http.Handle("/products/", *pc)
	http.Handle("/cart/", *cc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
