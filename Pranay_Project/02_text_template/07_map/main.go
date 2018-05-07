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

	stages := map[string]string{
		"India":    "Ghandhi",
		"America":  "MLK",
		"Meditate": "Bhudha",
		"Love":     "Jesus",
	}

	err := tpl.Execute(os.Stdout, stages)

	if err != nil {
		log.Fatalln(err)
	}

}
