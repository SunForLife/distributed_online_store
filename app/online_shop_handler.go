package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Describes Product.
type Product struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Category string `json:"category"`
}

// Handler for HTTP requests, that stores all data.
type OnlineShopHandler struct {
	Db *gorm.DB
}

// ProductList handler method of OnlineShopHandler.
func (osh *OnlineShopHandler) handlerProductList(w http.ResponseWriter, r *http.Request) {
	log.Println("Got get-product-list request")

	products := []Product{}
	osh.Db.Find(&products)

	json, err := json.Marshal(products)
	if err != nil {
		errorRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(json))
}

// ProductInfo handler method of OnlineShopHandler.
func (osh *OnlineShopHandler) handlerProductInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("Got get-product-info request")

	if len(r.URL.Query()["name"]) == 0 {
		badRequest(w, "name param not found")
		return
	}
	name := r.URL.Query()["name"][0]

	var product Product
	if err := osh.Db.Where("Name = ?", name).First(&product).Error; err != nil {
		notFoundRequest(w, fmt.Sprint("Not found product with name:", name))
		return
	}
	json, err := json.Marshal(product)
	if err != nil {
		errorRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(json))
}

// NewProduct handler method of OnlineShopHandler.
func (osh *OnlineShopHandler) handlerNewProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Got create-new-product request")

	if len(r.URL.Query()["name"]) == 0 {
		badRequest(w, "name param not found")
		return
	}
	name := r.URL.Query()["name"][0]
	if len(r.URL.Query()["code"]) == 0 {
		badRequest(w, "code param not found")
		return
	}
	code := r.URL.Query()["code"][0]
	if len(r.URL.Query()["category"]) == 0 {
		badRequest(w, "category param not found")
		return
	}
	category := r.URL.Query()["category"][0]

	osh.Db.Create(&Product{Name: name, Code: code, Category: category})
}

// ChangeProductByName handler method of OnlineShopHandler.
func (osh *OnlineShopHandler) handlerChangeProductByName(w http.ResponseWriter, r *http.Request) {
	log.Println("Got change-product-by-name request")

	if len(r.URL.Query()["old-name"]) == 0 {
		badRequest(w, "old-name param not found")
		return
	}
	oldName := r.URL.Query()["old-name"][0]
	if len(r.URL.Query()["name"]) == 0 {
		badRequest(w, "name param not found")
		return
	}
	name := r.URL.Query()["name"][0]
	if len(r.URL.Query()["code"]) == 0 {
		badRequest(w, "code param not found")
		return
	}
	code := r.URL.Query()["code"][0]
	if len(r.URL.Query()["category"]) == 0 {
		badRequest(w, "category param not found")
		return
	}
	category := r.URL.Query()["category"][0]

	var product Product
	if err := osh.Db.Where("Name = ?", oldName).First(&product).Error; err != nil {
		notFoundRequest(w, fmt.Sprint("Not found product with name:", oldName))
		return
	}

	osh.Db.Delete(&product, "Name = ?", oldName)
	osh.Db.Create(&Product{Name: name, Code: code, Category: category})
}

// DeleteProduct handler method of OnlineShopHandler.
func (osh *OnlineShopHandler) handlerDeleteProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Got delete-product request")

	if len(r.URL.Query()["name"]) == 0 {
		badRequest(w, "name param not found")
		return
	}
	name := r.URL.Query()["name"][0]

	var product Product
	if err := osh.Db.Where("Name = ?", name).First(&product).Error; err != nil {
		notFoundRequest(w, fmt.Sprint("Not found product with name:", name))
		return
	}

	osh.Db.Delete(&product, "Name = ?", name)
}