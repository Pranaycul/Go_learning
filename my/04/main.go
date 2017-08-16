package main

import "net/http"

type person struct {
	fName string
}

func (p *person)ServeHTTP(w http.ResponseWriter, r *http.Request)  {

	w.Write([]byte("First name : "+p.fName))
}
func main()  {

	personOne :=&person{fName:"Pranay"}
	http.ListenAndServe(":8080/", personOne)

}
