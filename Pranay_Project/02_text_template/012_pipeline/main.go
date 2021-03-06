package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   sQuare,
	"fsqrt": squareRoot,
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(x int) int {
	return x + x
}

func sQuare(x int) float64 {
	return math.Pow(float64(x), 2)
}

func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)

	if err != nil {
		log.Fatalln(err)
	}
}
