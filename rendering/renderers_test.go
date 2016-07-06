package rendering

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/corvuscrypto/birdnest/requests"
)

type MockResponse struct {
	header http.Header
}

func (m *MockResponse) Write(b []byte) (int, error) {
	m.WriteHeader(200)
	return 0, nil
}

func (m *MockResponse) WriteHeader(code int) {
	m.header.Set("Status", strconv.Itoa(code))
}

func (m *MockResponse) Header() http.Header {
	return m.header
}

func TestJsonRenderer(T *testing.T) {
	var mockRes = new(MockResponse)
	mockRes.header = make(http.Header)
	req := new(requests.Request)
	req.Response = mockRes
	req.Ctx = make(requests.Context)
	//just test for successful write
	err := JSONRenderer.Render(req)
	if err != nil {
		T.Errorf("unexpected error")
	}

	//test for failure
	//add a channel type into Ctx
	req.Ctx.Set("test", make(chan bool))
	err = JSONRenderer.Render(req)
	if err == nil {
		T.Errorf("unexpected success")
	}

	//test the new view
	mockRes = new(MockResponse)
	mockRes.header = make(http.Header)
	req2 := new(requests.Request)
	req2.Response = mockRes
	req2.Ctx = make(requests.Context)
	var handle = func(r *requests.Request) {}
	var a = JSONRenderer.NewView(handle)
	a(req2)
	if req2.Response.Header().Get("Status") != "200" {
		T.Errorf("Received incorrect status code")
	}

	//test the failure of the wrapped func
	req2.Ctx.Set("test", make(chan bool))
	a(req2)
	if req2.Response.Header().Get("Status") == "200" {
		T.Errorf("Received incorrect status code")
	}
}
