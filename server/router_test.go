package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/corvuscrypto/birdnest/requests"
	"github.com/julienschmidt/httprouter"
)

var asserter string
var params httprouter.Params

func testHandler(r *requests.Request) {
	asserter = r.Request.Method
	r.Response.WriteHeader(200)
}

func testParamHandler(r *requests.Request) {
	params = r.Params
	r.Response.WriteHeader(200)
}

func TestRouter(T *testing.T) {
	router := NewRouter(nil)

	server := httptest.NewServer(router)
	defer server.Close()

	//test a method-specific route...
	router.Handle("GET", "/", testHandler)
	http.Get(server.URL + "/")
	if asserter != "GET" {
		T.Errorf("Incorrect route utilized!")
	}

	//...and its failure
	res, _ := http.Post(server.URL+"/", "", nil)
	if res.StatusCode != http.StatusMethodNotAllowed {
		T.Errorf("Incorrect Response")
		fmt.Println(res.StatusCode)
	}

	//test for a notFound error
	res, _ = http.Get(server.URL + "/error")
	if res.StatusCode != http.StatusNotFound {
		T.Errorf("Incorrect Response")
		fmt.Println(res.StatusCode)
	}

	//test a single param GET route
	router.GET("/param/:p1", testParamHandler)
	http.Get(server.URL + "/param/test")
	if len(params) != 1 || params[0].Value != "test" {
		T.Errorf("Incorrect Parameters detected")
	}

	//test a wildcard GET route
	router.GET("/wild/*wild", testParamHandler)
	http.Get(server.URL + "/wild/test")
	if len(params) != 1 || params[0].Value != "/test" {
		T.Errorf("Incorrect Parameters detected")
	}

	//test a 2-param GET route
	router.GET("/param/:p1/:p2", testParamHandler)
	http.Get(server.URL + "/param/test/test2")
	if len(params) != 2 || params[0].Value != "test" || params[1].Value != "test2" {
		T.Errorf("Incorrect Parameters detected")
	}

	//Now we just test the other methods for good measure. Just for one route
	router.GET("/test/:method", testParamHandler)
	router.DELETE("/test/:method", testParamHandler)
	router.HEAD("/test/:method", testParamHandler)
	router.OPTIONS("/test/:method", testParamHandler)
	router.PUT("/test/:method", testParamHandler)
	router.PATCH("/test/:method", testParamHandler)
	router.POST("/test/:method", testParamHandler)

	methods := []string{
		"GET",
		"DELETE",
		"HEAD",
		"OPTIONS",
		"PUT",
		"PATCH",
		"POST",
	}

	for _, m := range methods {
		req, _ := http.NewRequest(m, server.URL+"/test/"+m, nil)
		http.DefaultClient.Do(req)
		if params[0].Value != m {
			T.Errorf("Incorrect route utilized")
		}
	}

}
