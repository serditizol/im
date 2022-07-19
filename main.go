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

// Папка, которая будет доступна на web сервере
func includeStaticFolder(folder string) {
	http.Handle(folder, http.StripPrefix(folder, http.FileServer(http.Dir("."+folder))))
}
func includeStatic() {
	includeStaticFolder("/css/")
	includeStaticFolder("/images/")
	includeStaticFolder("/fonts/")
	includeStaticFolder("/include/")
}

func include_str(url_str string, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(url_str, "shablon/header.html", "shablon/footer.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, include_str)
}

// Подключение страниц
func index(w http.ResponseWriter, r *http.Request) {
	include_str("templates/index.html", w, r)
}

func furl() {
	includeStatic()
	r := mux.NewRouter()
	// r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/", index).Methods("GET")
	http.Handle("/", r)

	http.ListenAndServe(":80", nil)
	fmt.Println("123")
}

func main() {
	furl()

}
