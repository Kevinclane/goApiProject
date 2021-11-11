package controllers

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

var api = "https://fakestoreapi.com/products"

type productsController struct {
	userIDPattern *regexp.Regexp
}

func (pc productsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/products/all" {
		if r.Method == http.MethodGet {
			pc.getAllProducts(w, r)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := pc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			pc.getProductById(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (pc *productsController) getAllProducts(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(api)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	sb := string(body)
	encodeResponseAsJSON(sb, w)
}

func (pc *productsController) getProductById(id int, w http.ResponseWriter) {
	apiStr := api + "/" + strconv.Itoa(id)
	res, err := http.Get(apiStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	sb := string(body)
	encodeResponseAsJSON(sb, w)
}

func newProductsController() *productsController {
	return &productsController{
		userIDPattern: regexp.MustCompile(`^/products/(\d+)/?`),
	}
}
