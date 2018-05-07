package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
	Door         int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": FirstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func FirstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	bhudha := sage{
		Name:  "Bhudha",
		Motto: "Love everyone",
	}

	mlk := sage{
		Name:  "MLK",
		Motto: "I have a dream",
	}
	ghandhi := sage{
		Name:  "Ghandhi",
		Motto: "Love",
	}
	jesus := sage{
		Name:  "Jesus",
		Motto: "Love everyone",
	}

	f := car{
		Manufacturer: "ford",
		Model:        "F150",
		Door:         2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corola",
		Door:         4,
	}

	stages := []sage{bhudha, mlk, ghandhi, jesus}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		stages,
		cars,
	}
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}
}
