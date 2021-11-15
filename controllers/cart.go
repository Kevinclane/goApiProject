package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Kevinclane/firstgoproject/models"
)

type cartsController struct {
	cartIDPattern *regexp.Regexp
}

func (cc cartsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matches := cc.cartIDPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		w.WriteHeader(http.StatusNotImplemented)
	}
	id, err := strconv.Atoi((matches[1]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	switch r.Method {
	case http.MethodGet:
		cc.getCartById(id, w)
	case http.MethodPut:
		cc.updateCart(id, w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

}

func (cc *cartsController) getCartById(cartId int, w http.ResponseWriter) {
	c, err := models.GetCartById(cartId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(c, w)
}

func (cc *cartsController) updateCart(userID int, w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var cu models.CartUpdate
	err := dec.Decode(&cu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	if cu.Path == "Add" {
		c, err := models.AddItemToCart(userID, cu.ProductID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		encodeResponseAsJSON(c, w)
	} else if cu.Path == "Remove" {
		c, err := models.RemoveItemFromCart(userID, cu.ProductID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		encodeResponseAsJSON(c, w)
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}

}

func newCartsController() *cartsController {
	return &cartsController{
		cartIDPattern: regexp.MustCompile(`^/cart/(\d+)/?`),
	}
}
