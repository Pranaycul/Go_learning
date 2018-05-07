package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init()  {

	tpl = template.Must(template.ParseGlob("*.gmao"))

}

func main()  {



	err := tpl.Execute(os.Stdout,nil)

	err = tpl.ExecuteTemplate(os.Stdout,"two.gmao",nil)

	if err!=nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"vespa.gmao",nil)

	if err!=nil {
		log.Fatalln(err)
	}


	err = tpl.ExecuteTemplate(os.Stdout,"one.gmao",nil)

	if err!=nil {
		log.Fatalln(err)
	}


	err = tpl.Execute(os.Stdout,nil)

	if err!=nil {
		log.Fatalln(err)
	}
}
