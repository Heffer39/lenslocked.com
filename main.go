package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template
var contactTemplate *template.Template

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err)
	}
	var notFound http.Handler = http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.NotFoundHandler = notFound
	http.ListenAndServe(":3000", r)
}

func notFound(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(writer, "<h1>We could not find the page you "+
		"were looking for :(</h1>"+
		"<p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
}

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(writer, nil); err != nil {
		panic(err)
	}
	//fmt.Fprint(writer, "<h1>Welcome to my awesome site!</h1>")
}

func contact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(writer, nil); err != nil {
		panic(err)
	}
	//fmt.Fprint(writer, "To get in touch, please send an email " +
	//	"to <a href=\"mailto:support@lenslinked.com\">"+
	//	"support@lenslinked.com</a>.")
}
