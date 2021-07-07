package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

)
func newRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	//declaring the static file directory and
	//point it to directory I made
	staticFileDirectory := http.Dir("./assets/")
	//declaring handler that routes requests to its respective filename
	//The fileserver is wrapped in "stripPrefix" method,
	//because I want to remove "/assets/index.html"
	//when looking for files.
	//If I wouldn't do that, the file would look for
	//"./assets/assets/index.html" and throw out an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	//"PathPrefix"method acts as a matcher and matches all
	//routes starting with "/assets/", instead absolute
	//route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/dog", getDogHandler).Methods("GET")
	r.HandleFunc("/dog", createDogHandler).Methods("POST")
	return r
}
func main(){
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}
