package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/corvuscrypto/birdnest/requests"
)

var middleWareValues = []string{}

func middleWareGenerator(value string) RequestHandler {
	return func(r *requests.Request) {
		middleWareValues = append(middleWareValues, value)
	}
}

func TestMiddleware(T *testing.T) {
	router := NewRouter(nil)

	server := httptest.NewServer(router)
	defer server.Close()

	//add test route
	router.Handle("GET", "/", testHandler)

	//Add a few handlers (3 to be exact)
	testValues := []string{"one", "two", "three"}

	for _, v := range testValues {
		RegisterMiddleware(middleWareGenerator(v))
	}

	//Now add a nil value to ensure we don't get weird behavior
	RegisterMiddleware()

	//do one test request
	http.Get(server.URL + "/")

	//check that we indeed got 3 executions
	if len(middleWareValues) != 3 {
		T.Errorf("Incorrect number of Middleware executions")
	}

	//check that the values are in order
	for i, v := range middleWareValues {
		if testValues[i] != v {
			T.Errorf("Encountered an out of order execution")
		}
	}
}
