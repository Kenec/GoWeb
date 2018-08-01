package main

import (
	"html/template"
	"fmt"
	"os"
)

func main() {
	templateString := `This is a sample template string`
	t,err := template.New("title").Parse(templateString)

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}


