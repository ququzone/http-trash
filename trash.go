package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/trash", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Printf("GET: %s\n", html.EscapeString(r.URL.Path))
		default:
			length, err := strconv.Atoi(r.Header.Get("Content-Length"))
			if err != nil || length == 0 {
				log.Printf("%s: No content\n", r.Method)
			} else {
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Printf("%s: Error content\n", r.Method)
				} else {
					log.Printf("%s: %s\n", r.Method, string(body))
				}
			}
		}
		fmt.Fprint(w, "Hello, this is trash.")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
