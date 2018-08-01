package main

import (
	"net/http"
	"html/template"
	"log"
)

func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		// First method

		//writer.Write([]byte("Hello World\n"))
		//fmt.Fprintf(writer, "Hello, you've requested: %s ", request.URL.Path)
		//f, err := os.Open("public" + request.URL.Path)
		//if err != nil {
		//	writer.WriteHeader(http.StatusInternalServerError)
		//	log.Print(err)
		//}
		//defer f.Close()
		//
		//var contentType string
		//switch  {
		//case strings.HasSuffix(request.URL.Path, "css"):
		//	contentType = "text/css"
		//case strings.HasSuffix(request.URL.Path, "html"):
		//	contentType = "text/html"
		//case strings.HasSuffix(request.URL.Path, "png"):
		//	contentType = "image/png"
		//default:
		//	contentType = "text/plain"
		//}
		//writer.Header().Add("Content-Type", contentType)
		//io.Copy(writer, f)

		// End of first method

		// Second method

		//http.ServeFile(writer, request, "public"+request.URL.Path)

		// End of second method
	//})

	template := poplulateTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := template.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	})
	http.Handle("/image/", http.FileServer(http.Dir("public")))
	http.Handle("/scss/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", nil)
}

func poplulateTemplate() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}