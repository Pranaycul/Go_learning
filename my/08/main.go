package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", myHandlerFunc)
	http.ListenAndServe(":8080", nil)

}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "Text/html")
	temp, err := template.New("anyNumberofTemplates").Parse(doc)
	if err == nil {
		temp.Execute(w, nil)
	}
}

const doc = `
<!DOCTYPE html>
<html>
<head lang = "en">
<meta charset= "UTF-8">
<title> MY Template </title>
</head>
<body>
<h1> Welcome to My world</h1>
</body>
</html>
`
