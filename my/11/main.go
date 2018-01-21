package main

import (
	"net/http"
	"text/template"
)

type Context struct {
	FirstName string
	Message   string
	URL       string
}

const doc = `
<!DOCTYPE html>
<html>
<head lang = "en">
<meta charset= "UTF-8">
<title> MY Template </title>
</head>
<body>
<h1>Hi to all</h1>
{{if eq .URL "/nobeer"}}
	<h1>We're out beer. Sorry</h1>
{{else}}
	<h1> Yes, Grab another beer,{{.FirstName}}</h1>
{{end}}
<hr>
<h2>Here's all the data :</h2>
<p>{{.}}</p>
</body>
</html>
`

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		context := Context{"Todd", "Can i get  beer, please ", req.URL.Path}
		tmpl.Execute(w, context)
	}
}

func main() {

	http.HandleFunc("/", myHandlerFunc)
	http.ListenAndServe(":8080", nil)
}
