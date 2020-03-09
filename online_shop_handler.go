package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Describes Product.
type Product struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Category string `json:"category"`
}

// Handler for HTTP requests, that stores all data.
type OnlineShopHandler struct {
	Products []Product
}

// handler method of OnlineShopHandler.
func (osh *OnlineShopHandler) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if strings.Contains(r.RequestURI, "get-product-list") {
			log.Println("Got get-product-list request")

			json, err := json.Marshal(osh.Products)
			if err != nil {
				ErrorRequest(w, err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, string(json))

		} else if strings.Contains(r.RequestURI, "get-product-info") {
			log.Println("Got get-product-info request")

			if len(r.URL.Query()["name"]) == 0 {
				BadRequest(w, "name param not found")
				return
			}
			name := r.URL.Query()["name"][0]

			for _, product := range osh.Products {
				if product.Name == name {
					json, err := json.Marshal(product)
					if err != nil {
						ErrorRequest(w, err)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprintln(w, string(json))
					return
				}
			}

			NotFoundRequest(w, fmt.Sprint("Not found product with name:", name))

		} else {
			BadRequest(w, "Got unexpected GET request")
		}
		return

	case "POST":
		if strings.Contains(r.RequestURI, "create-new-product") {
			log.Println("Got create-new-product request")

			if len(r.URL.Query()["name"]) == 0 {
				BadRequest(w, "name param not found")
				return
			}
			name := r.URL.Query()["name"][0]
			if len(r.URL.Query()["code"]) == 0 {
				BadRequest(w, "code param not found")
				return
			}
			code := r.URL.Query()["code"][0]
			if len(r.URL.Query()["category"]) == 0 {
				BadRequest(w, "category param not found")
				return
			}
			category := r.URL.Query()["category"][0]

			osh.Products = append(osh.Products, Product{Name: name, Code: code, Category: category})

		} else if strings.Contains(r.RequestURI, "change-product-by-name") {
			log.Println("Got change-product-by-name request")

			if len(r.URL.Query()["old-name"]) == 0 {
				BadRequest(w, "old-name param not found")
				return
			}
			oldName := r.URL.Query()["old-name"][0]
			if len(r.URL.Query()["name"]) == 0 {
				BadRequest(w, "name param not found")
				return
			}
			name := r.URL.Query()["name"][0]
			if len(r.URL.Query()["code"]) == 0 {
				BadRequest(w, "code param not found")
				return
			}
			code := r.URL.Query()["code"][0]
			if len(r.URL.Query()["category"]) == 0 {
				BadRequest(w, "category param not found")
				return
			}
			category := r.URL.Query()["category"][0]

			for i := range osh.Products {
				if osh.Products[i].Name == oldName {
					osh.Products[i] = Product{Name: name, Code: code, Category: category}
					return
				}
			}

			NotFoundRequest(w, fmt.Sprint("Not found product with name:", name))

		} else {
			BadRequest(w, "Got unexpected POST request")
		}
		return

	case "DELETE":
		if strings.Contains(r.RequestURI, "delete-product") {
			log.Println("Got delete-product request")
			if len(r.URL.Query()["name"]) == 0 {
				BadRequest(w, "name param not found")
				return
			}

			name := r.URL.Query()["name"][0]
			for i := range osh.Products {
				if osh.Products[i].Name == name {
					osh.Products[i] = osh.Products[len(osh.Products)-1]
					osh.Products = osh.Products[:len(osh.Products)-1]
					return
				}
			}

			NotFoundRequest(w, fmt.Sprint("Not found product with name:", name))

		} else {
			BadRequest(w, "Got unexpected DELETE request")
		}
		return

	}
}

// func ParseParamByName(w http.ResponseWriter, urlQuery url.Values, paramName string) string {
// 	return ""
// }

func BadRequest(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusBadRequest)
	if err != "" {
		log.Println("error: ", err)
	} else {
		log.Println("Bad request")
	}
}

func NotFoundRequest(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusNotFound)
	if err != "" {
		log.Println("error: ", err)
	} else {
		log.Println("Bad request")
	}
}

func ErrorRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		log.Println("error: ", err)
	} else {
		log.Println("Error while handling request")
	}
}
