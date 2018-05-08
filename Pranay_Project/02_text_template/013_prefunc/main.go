package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	bhudha := sage{
		Name:  "Bhudha",
		Motto: "Love everyone",
		Admin: true,
	}

	mlk := sage{
		Name:  "",
		Motto: "I have a dream",
		Admin: true,
	}
	ghandhi := sage{
		Name:  "Ghandhi",
		Motto: "Love",
		Admin: false,
	}

	stages := []sage{bhudha, mlk, ghandhi}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", stages)

	if err != nil {
		log.Fatalln(err)
	}
}
