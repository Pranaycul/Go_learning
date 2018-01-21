package main

import (
	"html/template"
	"net/http"
)

type Context struct {
	FirstName string
	Message   string
}

const doc = `
<!DOCTYPE html>
<html>
<head lang = "en">
<meta charset= "UTF-8">
<title> MY Template </title>
</head>
<body>
<h1> My Name is {{.FirstName}}</h1>
<p>  {{.Message}}</p>

</body>
</html>
`

func toddFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		context := Context{"Todd", "More go, please"}
		tmpl.Execute(w, context)
	}
}
func mingFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		context := Context{"Ming", "I am problem solver"}
		tmpl.Execute(w, context)
	}
}
func rioFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		context := Context{"Rio", "Lets party!!!"}
		tmpl.Execute(w, context)
	}
}
func jamesFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		context := Context{"James", "Go get me Beer"}
		tmpl.Execute(w, context)
	}
}

func main() {
	http.HandleFunc("/todd", toddFunc)
	http.HandleFunc("/ming", mingFunc)
	http.HandleFunc("/rio", rioFunc)
	http.HandleFunc("/", jamesFunc)
	http.ListenAndServe(":8080", nil)
}
