package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	fmt.Fprintf(w, str)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/public/images/", http.StripPrefix("/public/images/", http.FileServer(http.Dir("public/images"))))
	http.ListenAndServe(":8000", nil)
}
