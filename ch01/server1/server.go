// server1 : minimum echo server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // all request call "handler"
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler return request URL path
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
