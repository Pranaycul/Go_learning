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
	err := tpl.Execute(os.Stdout, `is to embress yourself before embrssing the world`)

	if err != nil {
		log.Fatalln(err)
	}

}
