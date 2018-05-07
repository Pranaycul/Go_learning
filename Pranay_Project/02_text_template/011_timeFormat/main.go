package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())

	if err != nil {
		log.Fatalln(err)
	}
}
