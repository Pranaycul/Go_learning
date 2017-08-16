package main

import "net/http"

func main()  {
	Mymux := http.NewServeMux()
	Mymux.HandleFunc("/",someFunc)
	http.ListenAndServe(":8080", Mymux)
}
func someFunc(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("Hello Mr. Wankhede. Welcome to new world"))
}
