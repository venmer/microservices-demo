package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("New Request:", r)
	data, err := ioutil.ReadFile("/data/currency_conversion.json")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	http.HandleFunc("/data", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
