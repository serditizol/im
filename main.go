package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type catalog struct {
	id                       int16
	name, image, text, price string
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "shablon/header.html", "shablon/footer.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, "index")
}

func furl() {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	http.Handle("/", r)

	http.ListenAndServe(":80", nil)
	fmt.Println("123")
}

func main() {
	furl()

}
