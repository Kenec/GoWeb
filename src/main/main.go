package main

import (
	"html/template"
	"net/http"
	"log"
)

func main() {
	templates := populateTemplate()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		requestedFile := request.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(writer, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
		http.Handle("/image/", http.FileServer(http.Dir("public")))
		http.Handle("/scss/", http.FileServer(http.Dir("public")))

	})
	http.ListenAndServe(":3000", nil)
}


func populateTemplate() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}