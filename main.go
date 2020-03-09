package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8000, "port")
	flag.Parse()

	onlineShopHandler := OnlineShopHandler{Products: make([]Product, 0)}

	log.Printf("Started on port: %d\n", *port)
	http.HandleFunc("/", onlineShopHandler.handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(*port), nil))
}