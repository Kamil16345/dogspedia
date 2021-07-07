package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T){
	//here I create new requests. They'll be passing to handler
	//First argument - method, second - route(empty for now),
	//third - request
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	//recorder is a target of request
	recorder := httptest.NewRecorder()
	//creating http handler from function handler in main
	hf := http.HandlerFunc(handler)
	//Serving http request to recorder
	//this lines executes handler that I want to test
	hf.ServeHTTP(recorder, req)
	//Check the status code is what I expect
	if status := recorder.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code, got: %v," +
			"expected %v", status, http.StatusOK)
	}
	//check the response body is what I expect
	expected := "Hello World"
	actual := recorder.Body.String()
	if actual != expected{
		t.Errorf("handler unexpected body: got %v, want %v", actual, expected)
	}
}
func TestRouter(t *testing.T){
	//initiating router constructor which I've defined
	//previously
	r := newRouter()
	//craeting a new server using "httptest" libraries
	//using `newServer` method https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)
	//mock server which I started runs and shows its
	//location at URL
	//I make GET request to "hello" route I've defined
	//in the router
	resp, err := http.Get(mockServer.URL+"/hello")

	if err != nil{
		t.Fatal(err)
	}
	//I want my status to be 200(ok)
	if resp.StatusCode != http.StatusOK{
		t.Errorf("Status should be ok, %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		t.Fatal(err)
	}
	//converting the bytes to a string
	respString := string(b)
	expected := "Hello world"
	//Verificating if response match to the defined in handler
	//If it's happend, it confirms
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}
//testing not existing requirement
func TestRouterForNonExistentRoute(t *testing.T){
	r := newRouter()
	mockServer := httptest.NewServer(r)
	//Only difference than in previous function is
	//that in this I make request to route I didn't define.
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil{
		t.Fatal(err)
	}
	//I want status to be 405(method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed{
		t.Errorf("Status should be 405, got %v", resp.StatusCode)
	}
	//Testing the body id almost the same, except
	//this time we want empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""
	if respString != expected{
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}
func TestStaticFileServer(t *testing.T){
	r := newRouter()
	mockServer := httptest.NewServer(r)
	//I want to hit `GET/assets/`route to get index.html
	// file response
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil{
		t.Fatal(err)
	}
	//I want to get status 200(ok)
	if resp.StatusCode != http.StatusOK{
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}
	//Instead of testing whole html file I test
	//if content-type-header is text/html; charset=utf-8"
	//to know that an html file has been delivered
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if expectedContentType != contentType{
		t.Errorf("Wrong content type, expected $s, got %s", expectedContentType, contentType)
	}
}