package main

import (
	"fmt"
	"github.com/doga10/hands-on-golang-dev-bage/src/endpoints"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", endpoints.Status)
	fmt.Println("Init Server GoLang: http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
