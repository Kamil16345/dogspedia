package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Dog struct {
	Species string `json:"species"`
	Description string `json:"description"`
}
var dogs []Dog
func getDogHandler(w http.ResponseWriter, r *http.Request){
	//converting  "dog" variable to json
	dogListBytes, err := json.Marshal(dogs)
	if err != nil{
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//if all is okey, write JSON list of dogs to the response
	w.Write(dogListBytes)
}
func createDogHandler(w http.ResponseWriter, r *http.Request){
	//creating a new instance of dog
	dog := Dog{}
	//I send all data as html form data
	//ParseForm() method of the request parses the form
	//values
	err := r.ParseForm()
	//if there's other error, it responds
	if err != nil{
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Get information about dog from form info
	dog.Species=r.Form.Get("species")
	dog.Description=r.Form.Get("description")
	//append existing list of dogs with a new entry
	dogs=append(dogs, dog)
	//Redirecting user to original html page
	//(located at `/assets/`) using http libraries by
	//using Redirect method
	http.Redirect(w,r,"/assets/", http.StatusFound)
}

