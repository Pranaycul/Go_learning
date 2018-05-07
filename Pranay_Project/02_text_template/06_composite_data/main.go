package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))

}

func main() {
	satges := []string{"Ghandi", "MLK", "Buddha", "Jesus"}

	err := tpl.Execute(os.Stdout, satges)

	if err != nil {
		log.Fatalln(err)
	}
}
