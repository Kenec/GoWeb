package main

import (
	"html/template"
	"net/http"
	"log"
	"os"
	"io/ioutil"
)

func main() {
	templates := populateTemplate()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		requestedFile := request.URL.Path[1:]
		t := templates[requestedFile + ".html"]
		if t != nil {
			err := t.Execute(writer, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/image/", http.FileServer(http.Dir("public")))
	http.Handle("/scss/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", nil)
}


func populateTemplate() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath + "/_head.html", basePath + "/_header.html", basePath + "/_script.html", basePath + "/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template block directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory" + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}

	return result
}