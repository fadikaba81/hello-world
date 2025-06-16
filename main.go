package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloworld)

	http.ListenAndServe(":8000", nil)

}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World", html.EscapeString(r.URL.Path))
}

