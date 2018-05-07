package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

func init()  {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main()  {

	bhudha:= sage{
		Name:"Bhudha",
		Motto:"Love everyone",
		}


	mlk:= sage{
		Name:"MLK",
		Motto:"I have a dream",
	}
	ghandhi:= sage{
		Name:"Ghandhi",
		Motto:"Love",
	}
	jesus:= sage{
		Name:"Jesus",
		Motto:"Love everyone",
	}

	stages:=[]sage{bhudha,mlk,ghandhi,jesus}
	err := tpl.Execute(os.Stdout,stages)

	if err!= nil {
		log.Fatalln(err)
	}
}

